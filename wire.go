//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Kimmoonsu/directconnect-api/internal/protocols"
	httprouter "github.com/Kimmoonsu/directconnect-api/internal/protocols/router"
	httphandler "github.com/Kimmoonsu/directconnect-api/src/handlers"

	"github.com/google/wire"
)

// Wiring for http protocol
var httpHandler = wire.NewSet(
	httphandler.NewHttpHandler,
)

// Wiring protocol routing
var httpRouter = wire.NewSet(
	httprouter.NewHttpRoute,
)

func InitHttpProtocol() *protocols.HttpImpl {
	wire.Build(
		httpHandler,
		httpRouter,
		protocols.NewHttpProtocol,
	)
	return &protocols.HttpImpl{}
}
