package router

import (
	"github.com/emadghaffari/api-teacher/app/http/middleware"
	"github.com/emadghaffari/api-teacher/controller/studentcontroller"
)

func student() {
	authorized := router.Group("/student")
	authorized.Use(middleware.AccessToken.CheckMiddleware)

	// Take new course
	authorized.POST("/store", studentcontroller.Router.Store)

	// show all courses taked by user
	authorized.GET("/", studentcontroller.Router.Index)
}
