package router

import (
	"github.com/siobh9/first-htmx/db"
	"github.com/siobh9/first-htmx/dist"
	"github.com/siobh9/first-htmx/server/handler"
	"github.com/siobh9/first-htmx/server/middleware"
	"log/slog"
	"net/http"
)

func New(logger *slog.Logger, database db.Database) http.Handler {
	h := &handler.Handler{
		Logger:   logger,
		Database: database,
	}

	mux := http.NewServeMux()

	mux.Handle(newPath(http.MethodGet, "/assets/"), middleware.CacheMiddleware(http.FileServer(http.FS(dist.AssetsDir))))
	mux.HandleFunc(newPath(http.MethodGet, "/"), h.Home)

	return middleware.NewLoggingMiddleware(logger, mux)
}

func newPath(method string, path string) string {
	return method + " " + path
}
