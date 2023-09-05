package httpServerMiddleware

import (
	"context"
	httpServerModel "github.com/conacry/go-platform/pkg/http/server/model"
	"github.com/google/uuid"
	"net/http"
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
		ctx := context.WithValue(r.Context(), "RequestID", requestID)
		w.Header().Add(httpServerModel.RequestIDHeaderName, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
