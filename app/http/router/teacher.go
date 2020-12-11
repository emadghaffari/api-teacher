package router

import (
	"github.com/emadghaffari/api-teacher/app/http/middleware"
	"github.com/emadghaffari/api-teacher/controller/teachercontroller"
)

func teacher() {
	authorized := router.Group("/teacher")
	authorized.Use(middleware.AccessToken.CheckMiddleware)

	middleware.Role.SetRole("teacher")
	authorized.Use(middleware.Role.Check)

	// show teacher courses
	authorized.GET("/", teachercontroller.Router.Index)
}
