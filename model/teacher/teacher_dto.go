package teacher

import (
	"github.com/emadghaffari/api-teacher/model/course"
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// Model for Teacher
	Model teacher = &Teacher{}
)

type teacher interface {
	Index() (course.Courses, errors.ResError)
}

// Teacher struct
type Teacher struct {
	user.User
	Courses course.Courses `json:"courses"`
}
