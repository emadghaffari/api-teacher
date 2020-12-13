package router

import (
	"github.com/emadghaffari/api-teacher/app/http/middleware"
	"github.com/emadghaffari/api-teacher/controller/coursecontroller"
)

func course() {

	authorized := router.Group("/course")
	authorized.Use(middleware.AccessToken.CheckMiddleware)

	// show all courses
	authorized.GET("/", coursecontroller.Router.Index)

	// Store new course
	authorized.Use(middleware.Role.Check("teacher"))
	{
		authorized.POST("/store", coursecontroller.Router.Store)

		// update new course
		authorized.POST("/update", coursecontroller.Router.Update)
	}

	// Take new course
	authorized.Use(middleware.Role.Check("student"))
	{
		authorized.POST("/take", coursecontroller.Router.Take)
	}
}
