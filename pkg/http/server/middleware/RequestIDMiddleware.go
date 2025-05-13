package httpServerMiddleware

import (
	"context"
	"net/http"

	httpServerModel "github.com/conacry/go-platform/pkg/http/server/model"
	"github.com/google/uuid"
)

type ValueKey string

const (
	RequestIDKey ValueKey = "RequestID"
)

type RequestIDDetectorMiddleware struct{}

func NewRequestIDDetector() *RequestIDDetectorMiddleware {
	return &RequestIDDetectorMiddleware{}
}

func (m *RequestIDDetectorMiddleware) Process(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(httpServerModel.RequestIDHeaderName)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		//TODO: Вынести ключи для Value контекста в отдельное место
		ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
		w.Header().Add(httpServerModel.RequestIDHeaderName, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
