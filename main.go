//mongodb://localhost:27017/minor-project2

package main

import (
	"fmt"
	"github.com/its-dastan/go-blog/routes"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s! ", r.URL.Path[1:])
}

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server is running on", 3000)
}

func main() {


	fmt.Print("Listening on 3000")

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
