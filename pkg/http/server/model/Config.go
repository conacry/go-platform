package httpServerModel

import "net/http"

type Config struct {
	ReadTimeout     int
	WriteTimeout    int
	IdleTimeOut     int
	ShutdownTimeout int
	port            int

	CORS *CORS

	publicHandlers  []*RequestHandler
	privateHandlers []*RequestHandler
	middlewares     []func(http.Handler) http.Handler
}

func NewDefaultConfig(port int) *Config {
	return &Config{
		ReadTimeout:     30,
		WriteTimeout:    30,
		IdleTimeOut:     2 * 30,
		ShutdownTimeout: 10,
		port:            port,
		CORS:            NewCORS(),
		publicHandlers:  make([]*RequestHandler, 0),
		privateHandlers: make([]*RequestHandler, 0),
		middlewares:     make([]func(http.Handler) http.Handler, 0),
	}
}

func (c *Config) RegisterPublicHandler(method, route string, handlerFn http.HandlerFunc) {
	c.publicHandlers = append(c.publicHandlers, &RequestHandler{
		method:     method,
		route:      route,
		handleFunc: handlerFn,
	})
}

func (c *Config) RegisterPrivateHandler(method, route string, handlerFn http.HandlerFunc) {
	c.privateHandlers = append(c.privateHandlers, &RequestHandler{
		method:     method,
		route:      route,
		handleFunc: handlerFn,
	})
}

func (c *Config) Port() int {
	return c.port
}

func (c *Config) PublicHandlers() []*RequestHandler {
	return c.publicHandlers
}

func (c *Config) PrivateHandlers() []*RequestHandler {
	return c.privateHandlers
}

func (c *Config) RegisterMiddleware(middleware func(http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, middleware)
}

func (c *Config) IsContainMiddlewares() bool {
	return len(c.middlewares) > 0
}

func (c *Config) Middlewares() []func(http.Handler) http.Handler {
	return c.middlewares
}
