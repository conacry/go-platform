package httpServer

import (
	"context"
	"errors"
	"fmt"
	httpServerMiddleware "github.com/conacry/go-platform/pkg/http/server/middleware"
	httpServerModel "github.com/conacry/go-platform/pkg/http/server/model"
	httpResponse "github.com/conacry/go-platform/pkg/http/server/response"
	log "github.com/conacry/go-platform/pkg/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	healthy int32
	ready   int32
)

type HttpServer struct {
	router              *chi.Mux
	server              *http.Server
	logger              log.Logger
	config              *httpServerModel.Config
	responseWriter      *httpResponse.Writer
	errorResponseWriter *httpResponse.ErrorWriter
}

func (s *HttpServer) Start(ctx context.Context) error {
	err := s.startWorker(ctx)
	if err != nil {
		s.logger.LogError(ctx, err)
		return err
	}

	msg := fmt.Sprintf("Http server is successfully started on port: %d", s.config.Port())
	s.logger.LogInfo(ctx, msg)
	return nil
}

func (s *HttpServer) Stop(ctx context.Context) error {
	err := s.stopServer(ctx)
	if err != nil {
		s.logger.LogError(ctx, err)
		return err
	}

	s.logger.LogInfo(ctx, "HTTP server is successfully stopped")
	return nil
}

func (s *HttpServer) startWorker(ctx context.Context) error {
	s.initHandlersAndMiddlewares()
	s.initServer()
	s.startHttpListener(ctx)
	s.setActiveState()
	return nil
}

func (s *HttpServer) initHandlersAndMiddlewares() {
	s.router.Group(func(r chi.Router) {
		s.initPublicMethods(r)
	})
}

func (s *HttpServer) initPublicMethods(r chi.Router) {
	requestIDDetectorMiddleware := httpServerMiddleware.NewRequestIDDetector()
	logRequestMiddleware := httpServerMiddleware.NewLogRequestMiddleware(s.logger, []string{})

	s.initCustomMiddlewares(r)

	r.Use(requestIDDetectorMiddleware.Process)
	r.Use(logRequestMiddleware.Process)
	r.Get(httpServerModel.HealthURL, s.healthzHandler)
	r.Get(httpServerModel.ReadyURL, s.readyzHandler)
	r.Options("/*", func(w http.ResponseWriter, r *http.Request) {})

	for _, requestHandler := range s.config.PublicHandlers() {
		r.MethodFunc(requestHandler.Method(), requestHandler.Route(), requestHandler.HandlerFunc())
	}
}

func (s *HttpServer) initCustomMiddlewares(r chi.Router) {
	if !s.config.IsContainMiddlewares() {
		return
	}

	for _, middleware := range s.config.Middlewares() {
		r.Use(middleware)
	}
}

func (s *HttpServer) initServer() {
	s.server = &http.Server{
		Addr:         fmt.Sprintf(":%v", s.config.Port()),
		WriteTimeout: time.Duration(s.config.WriteTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(s.config.ReadTimeout) * time.Millisecond,
		IdleTimeout:  time.Duration(s.config.IdleTimeOut) * time.Millisecond,
		Handler:      s.router,
	}
}

func (s *HttpServer) startHttpListener(ctx context.Context) {
	go func() {
		err := s.server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			s.logError(ctx, "ListenAndServe", err)
		}
	}()
}

func (s *HttpServer) stopServer(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(s.config.ShutdownTimeout)*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logError(ctx, "HTTP server graceful shutdown failed", err)
		return err
	}

	s.setInactiveState()
	return nil
}

func (s *HttpServer) setActiveState() {
	atomic.StoreInt32(&healthy, 1)
	atomic.StoreInt32(&ready, 1)
}

func (s *HttpServer) setInactiveState() {
	atomic.StoreInt32(&healthy, 0)
	atomic.StoreInt32(&ready, 0)
}

func (s *HttpServer) logError(ctx context.Context, msg string, err error) {
	s.logger.LogError(ctx, fmt.Errorf("%s: %v", msg, err))
}

func (s *HttpServer) healthzHandler(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		err := s.responseWriter.JSONResponse(w, r, map[string]string{"status": "OK"}, http.StatusOK)
		if err != nil {
			s.errorResponseWriter.ErrorResponse(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusServiceUnavailable)
}

func (s *HttpServer) readyzHandler(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&ready) == 1 {
		err := s.responseWriter.JSONResponse(w, r, map[string]string{"status": "OK"}, http.StatusOK)
		if err != nil {
			s.errorResponseWriter.ErrorResponse(w, r, err)
		}
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}
