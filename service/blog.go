package service

import (
	"errors"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/its-dastan/go-blog/db"
	"github.com/its-dastan/go-blog/models"
)

const (
	blogsCollection    = "blogs"
	likesCollection    = "likes"
	commentsCollection = "comments"
)

func GetBlogs(results interface{}) error{
	s, c := db.Connect(blogsCollection)
	defer s.Close()
	return c.Find(nil).All(results)
}

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

func UpdateBlog(blog *models.Blog, blogId string) (string, error) {
	s, c := db.Connect(blogsCollection)
	defer s.Close()
	err := c.UpdateId(bson.ObjectIdHex(blogId), bson.M{"$set": blog})
	if err != nil {
		return "", err
	}
	return "Updated the blog", nil
}

func LikeOrDislike(like models.Likes) (string, error) {
	// Linking to the likes, blogs, user Collection
	s, c := db.Connect(likesCollection)
	s1, c1 := db.Connect(blogsCollection)
	s2, c2 := db.Connect(usersCollection)

	// Closing two sessions
	defer s.Close()
	defer s1.Close()
	defer s2.Close()

	// Creating variables of likeD
	var likeD *models.Likes

	err := c.Find(bson.M{"likedBy": like.LikedBy, "blogId": like.BlogId}).One(&likeD)
	if err == nil {
		// If the user has already liked the blog

		// Removing like from the blog.Likes slice
		change := mgo.Change{
			Update: bson.M{
				"$pull": bson.M{
					"likes": likeD.ID,
				},
			},
			ReturnNew: true,
		}
		_, err = c1.FindId(likeD.BlogId).Apply(change, nil)
		if err != nil {
			return "", err
		}

		// removing the blog to the user's liked blogs
		match := bson.M{"_id": likeD.LikedBy}
		change1 := bson.M{"$pull": bson.M{"likedBlogs": likeD.BlogId}}
		err = c2.Update(match, change1)
		if err != nil {
			return "", nil
		}

		// Removing like document from likes collection
		err = c.RemoveId(likeD.ID)
		if err != nil {
			return "", nil
		}
		return "Disliked The Blog", nil
	} else {
		// If the user has not liked the blog

		// Create a like document
		err = c.Insert(like)
		if err != nil {
			return "", nil
		}

		// Getting the liked data
		count, err := c.Count()
		if err != nil {
			return "", nil
		}
		err = c.Find(nil).Skip(count - 1).One(&likeD)
		if err != nil {
			return "", nil
		}

		// Adding the like to the blog
		match := bson.M{"_id": likeD.BlogId}
		change := bson.M{"$push": bson.M{"likes": likeD.ID}}
		err = c1.Update(match, change)
		if err != nil {
			return "", nil
		}

		// Adding the blog to the user's liked blogs
		match = bson.M{"_id": likeD.LikedBy}
		change = bson.M{"$push": bson.M{"likedBlogs": likeD.BlogId}}
		err = c2.Update(match, change)
		if err != nil {
			return "", nil
		}
		return "Liked The Blog", nil
	}
}

func AddComment(comment *models.Comments) (string, error) {
	s, c := db.Connect(blogsCollection)
	defer s.Close()
	var result *models.Blog
	err := c.FindId(comment.BlogId).One(&result)
	if err != nil {
		return "", err
	}
	s1, c1 := db.Connect(commentsCollection)
	defer s1.Close()
	err = c1.Insert(comment)
	if err != nil {
		return "", err
	}
	var commentResult *models.Comments
	count, err := c1.Count()
	err = c1.Find(nil).Skip(count - 1).One(&commentResult)
	if err != nil {
		return "", err
	}
	change := bson.M{"$push": bson.M{"comments": commentResult.ID}}
	err = c.UpdateId(comment.BlogId, change)
	if err != nil {
		return "", err
	}
	err = c.FindId(comment.BlogId).One(&result)
	if err != nil {
		return "", err
	}
	return "Commented", nil
}
