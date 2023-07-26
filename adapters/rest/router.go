package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type router struct {
	route chi.Router
	port  int
}

func GetRouterInstance(port int) *router {
	return &router{
		route: chi.NewRouter(),
		port:  port,
	}
}

func (r *router) Get(path string, f http.HandlerFunc) {
	r.route.Get(path, f)
}

func (r *router) Post(path string, f http.HandlerFunc) {
	r.route.Post(path, f)
}

func (r *router) ListenAndServe() {
	http.ListenAndServe(fmt.Sprintf(":%d", r.port), r.route)
}
