package httpServerModel

type CORS struct {
	allowedHeaders []string
	allowedOrigins []string
}

func NewCORS() *CORS {
	return &CORS{
		allowedHeaders: make([]string, 0),
		allowedOrigins: make([]string, 0),
	}
}

func (c *CORS) AddHeader(header ...string) {
	c.allowedHeaders = append(c.allowedHeaders, header...)
}

func (c *CORS) AddOrigin(origin ...string) {
	c.allowedOrigins = append(c.allowedOrigins, origin...)
}

func (c *CORS) AllowedHeaders() []string {
	return c.allowedHeaders
}

func (c *CORS) AllowedOrigins() []string {
	return c.allowedOrigins
}
