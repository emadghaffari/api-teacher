package router

import (
	"github.com/emadghaffari/api-teacher/app/http/middleware"
	"github.com/emadghaffari/api-teacher/controller/coursecontroller"
)

func course() {
	// show all courses
	router.GET("/course", coursecontroller.Router.Index)

	teacher := router.Group("/course")
	teacher.Use(middleware.AccessToken.CheckMiddleware)
	teacher.Use(middleware.Role.Check("teacher"))

	// Store new course
	teacher.POST("/store", coursecontroller.Router.Store)
	// update new course
	teacher.POST("/update", coursecontroller.Router.Update)

	student := router.Group("/course")
	student.Use(middleware.AccessToken.CheckMiddleware)
	student.Use(middleware.Role.Check("student"))

	// Take new course
	student.POST("/take", coursecontroller.Router.Take)
}
