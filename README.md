# Go-Blog

The project is to make an API to do operation like Login, Register, posting blog and update it and like and dislike and comment on it.


This project uses Golang as backend and mongodb as database.
<br />
As it's in development phase, you need run it on locally.
<br />

### Add packages 
- set/export GO111MODULE=on if not enable.
```
$ go mod tidy
```

### Run Server
```
$ go run main.go
```

### About Packages:
- gorilla/mux : Package [gorilla/mux](https://github.com/gorilla/mux) implements a request router and dispatcher for matching incoming requests to their respective handler.
- gorilla/websocket : The Gorilla [WebSocket](https://github.com/gorilla/websocket) package provides a complete and tested implementation of the WebSocket protocol. The package API is stable.

## Differ Operation : It's Routes
* **Login** : localhost:3000/auth/login (POST)
* **Register** :localhost:3000/auth/register (POST)
* **Get All Blogs** : localhost:3000 (GET)
* **Create A Blog** : localhost:3000/blog/add-blog/{userId} (POST)
* **Update A Blog** : localhost:3000/blog/update-blog/{blogId} (PUT)
* **Like Or Dislike a Blog** : localhost:3000/blog/like-dislike/{userId}/{blogId} (GET)
* **Add A Comment to the Blog** : localhost:3000/blog/add-comment/{userId}/{blogId} (POST)



