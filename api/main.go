package main

import (
	"context"
	"net/http"

	"github.com/elizabeth-dev/f1catalog/api/ports"
	"github.com/elizabeth-dev/f1catalog/api/server"
	"github.com/elizabeth-dev/f1catalog/api/service"
	"github.com/go-chi/chi"
)

func main() {
	ctx := context.Background()

	app := service.NewApplication(ctx)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})
}
