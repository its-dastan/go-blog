package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/its-dastan/go-blog/controllers"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("POST", "/auth/register", controllers.Register, nil)
	register("POST", "/auth/login", controllers.Login, nil)
	//register("GET", "/", controllers.GetBlogs, nil)
	register("POST", "/blog/add-blog/{userId}", controllers.AddBlog, nil)
	register("GET", "/blog/like-dislike/{userId}/{blogId}", controllers.LikeOrDislike, nil)
	register("PUT", "/blog/update-blog/{blogId}", controllers.UpdateBlog, nil)
	register("POST", "/blog/add-comment/{userId}/{blogId}", controllers.AddComment, nil)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).
			Path(route.Pattern)
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
