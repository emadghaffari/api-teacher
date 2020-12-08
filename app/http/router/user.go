package router

import "github.com/emadghaffari/api-teacher/controller/usercontroller"

func user() {
	// login
	router.POST("/login", usercontroller.Router.Login)

	// register
	router.POST("/register", usercontroller.Router.Register)
}
