package controllers

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/its-dastan/go-blog/models"
	"github.com/its-dastan/go-blog/service"
	"log"
	"time"
)

func Handler() *socketio.Server {
	options := &engineio.Options{
		PingTimeout: time.Hour,
	}
	server := socketio.NewServer(options)
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		ticker := time.NewTicker(time.Second * 10)
		i := 0
		for {
			var results []*models.Blog
			err := service.GetBlogs(&results)
			if err != nil {
				log.Println(err)
				return
			}
			s.Emit("reply", results)
			i++
			fmt.Println(i)
			<-ticker.C
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
