package router

import (
	"github.com/emadghaffari/api-teacher/app/http/middleware"
	"github.com/emadghaffari/api-teacher/controller/teachercontroller"
)

func teacher() {
	authorized := router.Group("/teacher")
	authorized.Use(middleware.AccessToken.CheckMiddleware)

	// store new course by teacher
	authorized.POST("/store", teachercontroller.Router.Store)
}
