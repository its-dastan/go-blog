package service

import (
	"errors"
	"github.com/its-dastan/go-blog/db"
	"github.com/its-dastan/go-blog/models"
)

const (
	blogsCollection = "blogs"
)

func AddBlog(blog *models.Blog, result interface{}) error {
	s, c := db.Connect(blogsCollection)
	defer s.Close()
	err := c.Insert(blog)
	if err != nil {
		return errors.New("internal error! please try again later")
	}
	count, _ := c.Count()
	return c.Find(nil).Skip(count - 1).One(result)
}
