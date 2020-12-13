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
	authorized.POST("/store", coursecontroller.Router.Store, middleware.Role.Check("teacher"))

	// update new course
	authorized.POST("/update", coursecontroller.Router.Update, middleware.Role.Check("teacher"))

	// Take new course
	authorized.POST("/take", coursecontroller.Router.Take, middleware.Role.Check("student"))
}
