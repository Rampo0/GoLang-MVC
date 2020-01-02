package controllers

import (
	"fmt"
	"go_framework/app/models"
	"go_framework/database"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func LoginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{})
}

func RegisterPage(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", map[string]interface{}{})
}

func Index(c echo.Context) error {

	// middleware
	user := AuthGetUser(c)
	if user.Username == "" {
		return c.Redirect(http.StatusSeeOther, "/user/login")
	}
	// end middleware

	db := database.Connect()
	postModel := models.PostModel{Db: db}
	posts := postModel.FindAll()

	defer db.Close()

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"posts": posts,
		"user":  user,
	})
}

func CreatePost(c echo.Context) error {

	title := c.FormValue("title")
	content := c.FormValue("content")

	// Source
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	image := file.Filename
	dst, err := os.Create("public/static/media/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// Exec Query

	db := database.Connect()

	sqlRes, err := db.Prepare("INSERT INTO posts(title , content, image) VALUES(? ,?, ?)")

	if err != nil {
		panic(err.Error())
	}

	sqlRes.Exec(title, content, image)

	defer db.Close()

	return c.Redirect(http.StatusSeeOther, "/")
}

func DeletePost(c echo.Context) error {

	db := database.Connect()
	postId := c.Param("id")

	// delete file

	postModel := models.PostModel{Db: db}
	post := postModel.Find(postId)

	err2 := os.Remove("public/static/media/" + post.Image)
	if err2 != nil {
		fmt.Printf("Delete file failed !!")
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	postModel.Delete(post)

	defer db.Close()

	return c.Redirect(http.StatusSeeOther, "/")
}

func EditPost(c echo.Context) error {
	db := database.Connect()
	postId := c.Param("id")

	postModel := models.PostModel{Db: db}
	post := postModel.Find(postId)

	defer db.Close()

	return c.Render(http.StatusOK, "edit.html", map[string]interface{}{
		"post": post,
	})

}

func UpdatePost(c echo.Context) error {

	title := c.FormValue("title")
	content := c.FormValue("content")
	postId := c.FormValue("post_id")

	image, imageErr := c.FormFile("image")

	db := database.Connect()
	postModel := models.PostModel{Db: db}

	post := postModel.Find(postId)

	if title != "" {
		post.Title = title
	}

	if content != "" {
		post.Content = content
	}

	switch imageErr {
	case nil:
		// delete old file
		err2 := os.Remove("public/static/media/" + post.Image)
		if err2 != nil {
			fmt.Printf("Delete file failed !!")
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}

		// create file

		src, err := image.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create("public/static/media/" + image.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		post.Image = image.Filename
	case http.ErrMissingFile:
		log.Println("No Image File")
	default:
		log.Println(imageErr)
	}

	postModel.Update(post)

	defer db.Close()

	return c.Redirect(http.StatusSeeOther, "/")
}

func PostApiList(c echo.Context) error {
	db := database.Connect()
	postModel := models.PostModel{Db: db}
	posts := postModel.FindAll()

	defer db.Close()

	return c.JSON(http.StatusOK, posts)
}
