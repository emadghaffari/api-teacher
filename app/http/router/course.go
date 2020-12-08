package router

import (
	"github.com/emadghaffari/api-teacher/app/http/middleware"
	"github.com/emadghaffari/api-teacher/controller/coursecontroller"
)

func course() {

	authorized := router.Group("/course")
	authorized.Use(middleware.AccessToken.CheckMiddleware)

	// Take new course
	authorized.POST("/store", coursecontroller.Router.Store)

	// show all courses
	authorized.GET("/", coursecontroller.Router.Index)
}
