package router

import (
	"github.com/emadghaffari/api-teacher/app/http/middleware"
	"github.com/emadghaffari/api-teacher/controller/teachercontroller"
)

func teacher() {
	authorized := router.Group("/teacher")
	authorized.Use(middleware.AccessToken.CheckMiddleware)

	authorized.Use(middleware.Role.Check("teacher"))
	{
		// show teacher courses
		authorized.GET("/", teachercontroller.Router.Index)
	}

}
