package httpServerModel

import "net/http"

type RequestHandler struct {
	method     string
	route      string
	handleFunc http.HandlerFunc
}

func (h *RequestHandler) Method() string {
	return h.method
}

func (h *RequestHandler) Route() string {
	return h.route
}

func (h *RequestHandler) HandlerFunc() http.HandlerFunc {
	return h.handleFunc
}
