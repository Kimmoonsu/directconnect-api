package router

import (
	"net/http"

	"github.com/Kimmoonsu/directconnect-api/src/handlers"
)

type HttpRouterImpl struct {
	handlers *handlers.HttpHandlerImpl
}

func NewHttpRoute(handlers *handlers.HttpHandlerImpl) *HttpRouterImpl {
	return &HttpRouterImpl{
		handlers: handlers,
	}
}

func (h *HttpRouterImpl) Router(r *http.ServeMux) {
	h.handlers.Router(r)
}
