package httpResponse

import (
	"context"
	"encoding/json"
	"fmt"
	httpServerModel "github.com/conacry/go-platform/pkg/http/server/model"
	log "github.com/conacry/go-platform/pkg/logger"
	"io"
	"net/http"
)

type Writer struct {
	logger log.Logger
}

func NewWriter(logger log.Logger) (*Writer, error) {
	if logger == nil {
		return nil, ErrLoggerIsRequired
	}

	return &Writer{logger: logger}, nil
}

func (s *Writer) JSONResponse(w http.ResponseWriter, r *http.Request, result interface{}, responseCode int) error {
	body, err := s.marshalBody(result)
	if err != nil {
		s.logError(r.Context(), "Marshal json error", err)
		return ErrWriteResponse
	}

	s.Response(w, r, body, responseCode)
	return nil
}

func (s *Writer) Response(w http.ResponseWriter, r *http.Request, result []byte, responseCode int) {
	w.Header().Set(httpServerModel.HeaderContentType, "application/json; charset=utf-8")
	w.Header().Set(httpServerModel.HeaderXContentTypeOptions, "nosniff")
	w.WriteHeader(responseCode)

	_, err := w.Write(result)
	if err != nil {
		s.logError(r.Context(), "Response Writer Error", err)
		return
	}
}

func (s *Writer) StreamResponse(w http.ResponseWriter, r *http.Request, resp *http.Response) {
	defer resp.Body.Close()
	w.Header().Set(httpServerModel.HeaderContentType, r.Header.Get(httpServerModel.HeaderContentType))
	w.Header().Set(httpServerModel.HeaderContentLength, r.Header.Get(httpServerModel.HeaderContentLength))

	_, err := io.Copy(w, resp.Body)
	if err != nil {
		s.logError(r.Context(), "io.Copy for body is failed", err)
	}
}

func (s *Writer) marshalBody(result interface{}) ([]byte, error) {
	if result == nil || result == "" {
		return []byte{}, nil
	}

	body, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *Writer) logError(ctx context.Context, msg string, err error) {
	s.logger.LogError(ctx, fmt.Errorf("%s: %v", msg, err))
}
