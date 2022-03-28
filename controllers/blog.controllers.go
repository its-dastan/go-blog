package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/its-dastan/go-blog/db"
	"log"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/its-dastan/go-blog/helper"
	"github.com/its-dastan/go-blog/models"
	"github.com/its-dastan/go-blog/service"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	fmt.Println(key)
	buffer := db.GetImage(key)
	w.Header().Set("Content-Type", "image/png")
	w.Write(buffer)
	return
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	var results []*models.Blog
	err := service.GetBlogs(&results)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: "Successfully registered", Data: results})
}

func AddBlog(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseMultipartForm(10 << 20)
	vars := mux.Vars(r)
	var result *models.Blog
	//var blogData *models.Blog

	blogData := &models.Blog{
		Caption: r.Form.Get("caption"),
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileName, err := db.UploadFileToS3(file, handler)
	if err != nil {
		//fmt.Fprintf(w, "Could not upload file")
		fmt.Println(err)
	}
	if blogData.Caption == "" && blogData.Image == "" && blogData.Video == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please give an input"})
		return
	}
	blogData.Image = fileName
	blogData.PostedBy = bson.ObjectIdHex(vars["userId"])
	blogData.CreatedAt = time.Now()
	err = service.AddBlog(blogData, &result)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}

	helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: "Successfully registered", Data: result})
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	var blogData *models.Blog
	vars := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&blogData)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	if blogData.Caption == "" && blogData.Image == "" && blogData.Video == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please give an input"})
		return
	}
	str, err := service.UpdateBlog(blogData, vars["blogId"])
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: str})
}

func LikeOrDislikeBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var likeData models.Likes
	likeData.LikedAt = time.Now()
	likeData.LikedBy = bson.ObjectIdHex(vars["userId"])
	likeData.BlogId = bson.ObjectIdHex(vars["blogId"])

	str, err := service.LikeOrDislike(likeData)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	log.Println("str")
	helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: str})
}

func AddComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var commentData *models.Comments

	err := json.NewDecoder(r.Body).Decode(&commentData)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}

	if commentData.Comment == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: "Please give an input"})
		return
	}

	commentData.CommentedAt = time.Now()
	commentData.BlogId = bson.ObjectIdHex(vars["blogId"])
	commentData.CommentedBy = bson.ObjectIdHex(vars["userId"])

	str, err := service.AddComment(commentData)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, helper.Response{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}

	helper.ResponseWithJson(w, http.StatusOK, helper.Response{Code: http.StatusOK, Msg: str})
}