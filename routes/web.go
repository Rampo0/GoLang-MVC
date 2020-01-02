package routes

import (
	"go_framework/app/http/controllers"

	"github.com/labstack/echo"
)

func Index() *echo.Echo {

	e := config()

	// routes
	e.GET("/", controllers.Index)

	userGroup := e.Group("/user")
	userGroup.GET("/login", controllers.LoginPage)
	userGroup.GET("/register", controllers.RegisterPage)
	userGroup.GET("/logout", controllers.Logout)

	userGroup.POST("/register/submit", controllers.Register)
	userGroup.POST("/login/submit", controllers.Login)

	postGroup := e.Group("/post")
	postGroup.POST("/create", controllers.CreatePost)
	postGroup.GET("/delete/:id", controllers.DeletePost)
	postGroup.GET("/edit/:id", controllers.EditPost)
	postGroup.POST("/update", controllers.UpdatePost)

	apiGroup := e.Group("/api")
	apiGroup.GET("/posts", controllers.PostApiList)

	return e
}
