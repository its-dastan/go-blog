package controllers

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/its-dastan/go-blog/helper"
	"github.com/its-dastan/go-blog/models"
	"github.com/its-dastan/go-blog/service"
	"net/http"
	"time"
)

func AddBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var result *models.Blog
	var blogData *models.Blog
	err := json.NewDecoder(r.Body).Decode(&blogData)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	if blogData.Caption == "" && blogData.Image == "" && blogData.Video == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please give an input"})
		return
	}
	blogData.PostedBy = bson.ObjectIdHex(vars["userId"])
	blogData.CreatedAt = time.Now()
	err = service.AddBlog(blogData, &result)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: "Successfully registered", Data: result})
}

func LikeOrDislike(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var likeData models.Likes
	likeData.LikedAt = time.Now()
	likeData.LikedBy = bson.ObjectIdHex(vars["userId"])
	likeData.BlogId = bson.ObjectIdHex(vars["blogId"])
	//fmt.Println(likeData)
	str,err := service.LikeOrDislike(likeData)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: str})
}
