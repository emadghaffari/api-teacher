package course

import (
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// Model for Course
	Model course = &Course{}
)

type course interface {
	Index() ([]*Course, errors.ResError)
}

// Course struct
type Course struct {
	ID        int64     `json:"id"`
	Teacher   user.User `json:"teacher,omitempty"`
	Name      string    `json:"name,omitempty"`
	Identitiy string    `json:"identitiy,omitempty"`
	Valence   uint32    `json:"valence,omitempty"`
	Time      string    `json:"time,omitempty"`
}

// Courses var list of Course
type Courses []Course

// StoreValidate meth, validate items before store
func (cs *Course) StoreValidate() {

}
