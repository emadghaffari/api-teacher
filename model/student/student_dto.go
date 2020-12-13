package student

import (
	"github.com/emadghaffari/api-teacher/model/course"
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// Model for Student
	Model student = &Student{}
)

type student interface {
	Index() (course.Courses, errors.ResError)
}

// Student struct
type Student struct {
	user.User
	Courses course.Courses `json:"courses"`
}
