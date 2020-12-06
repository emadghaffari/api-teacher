package teacher

import (
	"github.com/emadghaffari/api-teacher/model/course"
	"github.com/emadghaffari/api-teacher/model/user"
)

var (
	// Model for Teacher
	Model teacher = &Teacher{}
)

type teacher interface{}

// Teacher struct
type Teacher struct {
	user.User
	Courses course.Courses `json:"courses"`
}
