package main

import (
	"fmt"
	"github.com/its-dastan/go-blog/controllers"
	"github.com/its-dastan/go-blog/route"
	"net/http"
)

func main() {
	r := route.NewRouter()
	fmt.Println("Listening on 3000")
	http.HandleFunc("/ws", controllers.Handler)
	_ = http.ListenAndServe(":3000", r)

}
