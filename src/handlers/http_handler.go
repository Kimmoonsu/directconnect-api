package handlers

import (
	"net/http"
)

type HttpHandlerImpl struct {
}

func NewHttpHandler() *HttpHandlerImpl {
	return &HttpHandlerImpl{}
}

func (h *HttpHandlerImpl) Router(r *http.ServeMux) {
	r.HandleFunc("/CreateConnection", h.CreateConnection)
	r.HandleFunc("/", h.DescribeConnections)
}
