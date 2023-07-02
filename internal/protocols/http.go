package protocols

import (
	"fmt"
	"net/http"

	"github.com/Kimmoonsu/directconnect-api/internal/protocols/router"
)

type HttpImpl struct {
	HttpRouter *router.HttpRouterImpl
	httpServer *http.Server
}

func NewHttpProtocol(HttpRouter *router.HttpRouterImpl) *HttpImpl {
	return &HttpImpl{
		HttpRouter: HttpRouter,
	}
}

func (p *HttpImpl) setupRouter(app *http.ServeMux) {
	p.HttpRouter.Router(app)
}

func (p *HttpImpl) Listen() {
	app := http.NewServeMux()

	p.setupRouter(app)

	serverPort := fmt.Sprintf(":%d", 5000)
	p.httpServer = &http.Server{
		Addr:    serverPort,
		Handler: app,
	}

	fmt.Println("Server started on Port %s", serverPort)
	p.httpServer.ListenAndServe()
}
