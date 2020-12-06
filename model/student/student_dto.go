package student

import (
	"github.com/emadghaffari/api-teacher/model/course"
	"github.com/emadghaffari/api-teacher/model/user"
)

var (
	// Model for Student
	Model student = &Student{}
)

type student interface{}

// Student struct
type Student struct {
	user.User
	Courses course.Courses `json:"courses"`
}
