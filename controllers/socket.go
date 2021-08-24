package controllers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/its-dastan/go-blog/helper"
	"github.com/its-dastan/go-blog/models"
	"github.com/its-dastan/go-blog/service"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Handler(w http.ResponseWriter, r *http.Request){
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {

		log.Println(err)
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})

		return
	}

	for{
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second*5)
		var results []*models.Blog
		err= service.GetBlogs(&results)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err:= conn.WriteJSON(&results); err!= nil {
			fmt.Println(err)
			return
		}
	}
}