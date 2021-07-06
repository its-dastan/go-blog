package main

import (
	"fmt"
	"github.com/its-dastan/go-blog/route"
	"net/http"
)

func main() {
	r := route.NewRouter()
	fmt.Println("Listening on 3000")
	_ = http.ListenAndServe(":3000", r)

}
