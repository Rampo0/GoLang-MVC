package models

import (
	"database/sql"
	"fmt"
	"go_framework/app/models/entities"
)

type PostModel struct {
	Db *sql.DB
}

func (postModel PostModel) FindAll() []entities.PostEntity {

	db := postModel.Db

	rows, err := db.Query("SELECT id, title, content, image, created_at FROM posts")
	if err != nil {
		fmt.Printf("Select Error \n")
		return nil
	}

	posts := []entities.PostEntity{}

	for rows.Next() {
		var id int64
		var title string
		var content string
		var image string
		var created_at string
		err2 := rows.Scan(&id, &title, &content, &image, &created_at)
		if err2 != nil {
			fmt.Printf("Scan Error \n")
			return nil
		}

		post := entities.PostEntity{ID: id, Title: title, Content: content, Image: image, Created_at: created_at}
		posts = append(posts, post)
	}

	return posts
}

func (postModel PostModel) Find(post_id string) entities.PostEntity {

	db := postModel.Db
	rows, err := db.Query("SELECT id, title, content, image, created_at FROM posts WHERE id = ?", post_id)
	if err != nil {
		fmt.Printf("Select Error \n")
		return entities.PostEntity{}
	}

	post := entities.PostEntity{}

	for rows.Next() {
		var id int64
		var title string
		var content string
		var image string
		var created_at string
		err2 := rows.Scan(&id, &title, &content, &image, &created_at)
		if err2 != nil {
			fmt.Printf("Scan Error \n")
			return entities.PostEntity{}
		}

		post = entities.PostEntity{id, title, content, image, created_at}
	}

	return post
}

func (postModel PostModel) Update(post entities.PostEntity) {

	db := postModel.Db
	sqlRes, err := db.Prepare("UPDATE posts SET title=?, content=?, image=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	sqlRes.Exec(post.Title, post.Content, post.Image, post.ID)

}

func (postModel PostModel) Delete(post entities.PostEntity) {

	db := postModel.Db
	sqlRes, err := db.Prepare("DELETE FROM posts WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	sqlRes.Exec(post.ID)
}
