//mongodb://localhost:27017/minor-project2

package main

import (
	"fmt"
	"github.com/its-dastan/go-blog/routes"
	"log"
	"net/http"
	"time"
)





func main() {
	fmt.Println("Hello")
	r:= routes.NewRouter()
	s := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
