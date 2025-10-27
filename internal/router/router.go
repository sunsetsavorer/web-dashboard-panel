package router

import "net/http"

type Router struct {
	mux http.ServeMux
}

func New() *Router {

	return &Router{}
}

func (router *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {

	router.mux.HandleFunc(pattern, handler)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	router.mux.ServeHTTP(w, r)
}
