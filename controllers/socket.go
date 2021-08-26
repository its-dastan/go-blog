package controllers

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/its-dastan/go-blog/models"
	"github.com/its-dastan/go-blog/service"
	"log"
	"time"
)

func Handler() *socketio.Server{

	server := socketio.NewServer(nil)
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		s.SetContext("")
		for{
			var results []*models.Blog
			err:= service.GetBlogs(&results)
			if err != nil {
				log.Println(err)
				return
			}
			s.Emit("reply", results)
			time.Sleep(time.Second*10)
		}
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	return server
}