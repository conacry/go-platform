package httpServerMiddleware

import (
	"fmt"
	log "github.com/conacry/go-platform/pkg/logger"
	"net/http"
	"time"
)

type LogRequestMiddleware struct {
	logPublisher log.Logger
	excludedURL  map[string]bool
}

func NewLogRequestMiddleware(logger log.Logger, excludedURL []string) *LogRequestMiddleware {
	requestLogHandler := &LogRequestMiddleware{logPublisher: logger}
	requestLogHandler.setExcludedUrl(excludedURL)
	return requestLogHandler
}

func (m *LogRequestMiddleware) Process(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.isExcluded(r) {
			next.ServeHTTP(w, r)
			return
		}

		startTime := time.Now()
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		defer func() {
			message := fmt.Sprintf("%s %s://%s%s %s\", ", r.Method, scheme, r.Host, r.URL.Path, r.Proto) +
				fmt.Sprintf("from '%s %s', ", r.RemoteAddr, r.UserAgent()) +
				fmt.Sprintf("duration %s, ", time.Since(startTime)) +
				fmt.Sprintf("response code %d", http.StatusOK)

			m.logPublisher.LogInfo(r.Context(), message)
		}()

		next.ServeHTTP(w, r)
	})
}

func (m *LogRequestMiddleware) isExcluded(r *http.Request) bool {
	_, ok := m.excludedURL[r.URL.Path]
	return ok
}

func (m *LogRequestMiddleware) setExcludedUrl(urls []string) {
	excluded := make(map[string]bool)
	for _, url := range urls {
		excluded[url] = true
	}
	m.excludedURL = excluded
}
