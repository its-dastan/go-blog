package controllers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/its-dastan/go-blog/models"
	"log"
	"net/http"
	"time"
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

		time.Sleep(time.Second*5)
		var results []*models.Blog

		if err:= conn.WriteJSON(&results); err!= nil {
			fmt.Println(err)
			return
		}
	}
}