package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/its-dastan/go-blog/models"
	"github.com/its-dastan/go-blog/service"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(w http.ResponseWriter, r *http.Request){

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for{
		_, p, _:= conn.ReadMessage()
		log.Println(string(p))

		//time.Sleep(time.Second*5)
		var results []*models.Blog
		err := service.GetBlogs(&results)
		if err!= nil {
			log.Println(err)
		}
		res, _ := json.Marshal(results)
		log.Println(string(res))
		if err:= conn.WriteMessage(1, res); err!= nil {
			fmt.Println(err)
			return
		}
	}
}