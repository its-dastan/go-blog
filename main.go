package main

import (
	"fmt"
	"github.com/its-dastan/go-blog/controllers"
	"github.com/its-dastan/go-blog/db"
	"github.com/its-dastan/go-blog/route"
	"net/http"
)

func main() {
	fmt.Println(db.ListBucket())
	r := route.NewRouter()
	server := controllers.Handler()
	go server.Serve()
	defer server.Close()
	r.Handle("/socket.io/", server)
	r.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Listening on 3000")

	_ = http.ListenAndServe(":3000", r)

}
