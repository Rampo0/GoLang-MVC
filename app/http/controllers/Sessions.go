package controllers

import (
	"encoding/gob"
	"go_framework/app/models/entities"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

var store *sessions.CookieStore

func init() {

	store = sessions.NewCookieStore([]byte("super-secret-key"))
	// authKeyOne := securecookie.GenerateRandomKey(64)
	// encryptionKeyOne := securecookie.GenerateRandomKey(32)

	// store = sessions.NewCookieStore(
	// 	authKeyOne,
	// 	encryptionKeyOne,
	// )

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(entities.UserEntity{})

}

func AuthGetUser(c echo.Context) entities.UserEntity {

	session, err := store.Get(c.Request(), os.Getenv("LOGIN_SESSION_ID"))
	if err != nil {
		return entities.UserEntity{}
	}

	val := session.Values["user"]
	var user = entities.UserEntity{}
	user, ok := val.(entities.UserEntity)
	if !ok {
		return entities.UserEntity{}
	}

	return user
}

func AuthLogin(c echo.Context, user entities.UserEntity) error {

	session, err := store.Get(c.Request(), os.Getenv("LOGIN_SESSION_ID"))
	if err != nil {
		return err
	}

	session.Values["user"] = user
	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}

	return nil
}

func AuthLogout(c echo.Context) error {
	session, err := store.Get(c.Request(), os.Getenv("LOGIN_SESSION_ID"))
	if err != nil {
		return err
	}

	session.Values["user"] = entities.UserEntity{}
	session.Options.MaxAge = -1

	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}

	return nil
}
