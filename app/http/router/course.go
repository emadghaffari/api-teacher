package router

import (
	"github.com/emadghaffari/api-teacher/app/http/middleware"
	"github.com/emadghaffari/api-teacher/controller/coursecontroller"
)

func course() {

	authorized := router.Group("/course")
	authorized.Use(middleware.AccessToken.CheckMiddleware)

	// Store new course
	middleware.Role.SetRole("teacher")
	authorized.POST("/store", coursecontroller.Router.Store, middleware.Role.Check)

	// update new course
	authorized.POST("/update", coursecontroller.Router.Update, middleware.Role.Check)

	// Take new course
	middleware.Role.SetRole("student")
	authorized.POST("/store", coursecontroller.Router.Take, middleware.Role.Check)

	// show all courses
	authorized.GET("/", coursecontroller.Router.Index)
}
