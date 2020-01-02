package models

import (
	"database/sql"
	"go_framework/app/models/entities"
)

type UserModel struct {
	Db *sql.DB
}

func (userModel UserModel) FindUser(username string) entities.UserEntity {

	db := userModel.Db

	var user = entities.UserEntity{}

	db.QueryRow(`
	SELECT id, 
	username, 
	email, 
	password 
	FROM users WHERE username=?
	`, username).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
		)

	defer db.Close()

	return user
}
