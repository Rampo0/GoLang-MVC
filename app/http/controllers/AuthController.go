package controllers

import (
	"go_framework/app/models"
	"go_framework/database"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	db := database.Connect()
	userModel := models.UserModel{
		Db: db,
	}

	user := userModel.FindUser(username)

	isAuth := CheckPasswordHash(password, user.Password)

	if isAuth == true {

		err := AuthLogin(c, user)
		if err != nil {
			panic(err.Error())
		}

		return c.Redirect(http.StatusSeeOther, "/")
	}

	defer db.Close()

	return c.Render(http.StatusOK, "login.html", map[string]interface{}{"error": "Login Failed"})
}

func Logout(c echo.Context) error {
	err := AuthLogout(c)
	if err != nil {
		panic(err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/user/login")
}

func Register(c echo.Context) error {

	db := database.Connect()

	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	sqlRes, err := db.Prepare("INSERT INTO users(username , email, password) VALUES(? ,?, ?)")

	if err != nil {
		panic(err.Error())
	}

	hash, _ := HashPassword(password)

	sqlRes.Exec(username, email, hash)

	defer db.Close()

	return c.Redirect(http.StatusSeeOther, "/user/login")

}
