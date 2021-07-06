package routes

import (
	"github.com/gorilla/mux"
	"github.com/its-dastan/go-blog/controllers"
	"net/http"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("GET", "/get-users", controllers.GetUsers, nil)
	register("POST", "/register", controllers.Register, nil)
	register("POST", "/login", controllers.Login, nil)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).Path(route.Path)
		if route.Middleware != nil {
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler)
		}
	}
	return router
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
