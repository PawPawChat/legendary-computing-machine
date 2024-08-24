package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type MuxRouter struct {
	mux *mux.Router
}

type Route struct {
	Path    string
	Methods []string
	Handler http.Handler
}

type Config struct {
	Routes      []Route
	Middlewares []func(http.Handler) http.Handler
}

func NewMuxRouter(mux *mux.Router, config *Config) *MuxRouter {
	return (&MuxRouter{mux: mux}).configure(config)
}

func (r *MuxRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *MuxRouter) configure(config *Config) *MuxRouter {
	for _, route := range config.Routes {
		r.mux.NewRoute().Path(route.Path).Methods(route.Methods...).Handler(route.Handler)
	}

	for _, middleware := range config.Middlewares {
		r.mux.Use(middleware)
	}

	return r
}
