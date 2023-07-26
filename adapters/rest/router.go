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

func (r *router) Get(queryString string, f http.HandlerFunc) {
	r.route.Get(queryString, f)
}

func (r *router) ListenAndServe() {
	http.ListenAndServe(fmt.Sprintf(":%d", r.port), r.route)
}
