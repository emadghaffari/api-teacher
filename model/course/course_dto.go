package course

import (
	"strings"

	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// Model for Course
	Model course = &Course{}
)

type course interface {
	Index() (Courses, errors.ResError)
}

// Course struct
type Course struct {
	ID        int64      `json:"id,omitempty"`
	Teacher   *user.User `json:"teacher,omitempty"`
	Name      string     `json:"name,omitempty"`
	Identitiy string     `json:"identitiy,omitempty"`
	Valence   uint32     `json:"valence,omitempty"`
	Time      string     `json:"time,omitempty"`
}

// Courses var list of Course
type Courses []*Course

// StoreValidate meth, validate items before store
func (cs *Course) StoreValidate() errors.ResError {
	if cs.Name == "" {
		return errors.HandlerBadRequest("Invalid Course Name")
	}
	if cs.Time == "" {
		return errors.HandlerBadRequest("Invalid Course Time")
	}
	if cs.Valence <= 1 {
		return errors.HandlerBadRequest("invalid Course Valence")
	}
	return nil
}

// UpdateValidate meth, validate items before Update
func (cs *Course) UpdateValidate() errors.ResError {
	if cs.Name == "" {
		return errors.HandlerBadRequest("Invalid Course Name")
	}
	cs.Identitiy = strings.TrimSpace(cs.Identitiy)
	if cs.Identitiy == "" {
		return errors.HandlerBadRequest("Invalid Course Identitiy")
	}
	if cs.Time == "" {
		return errors.HandlerBadRequest("Invalid Course Time")
	}
	if cs.Valence <= 1 {
		return errors.HandlerBadRequest("invalid Course Valence")
	}
	return nil
}

// TakeValidate meth, validate items before Take
func (cs *Course) TakeValidate() errors.ResError {

	if user.Model.Get().ID == 0 {
		return errors.HandlerBadRequest("Invalid User Details")
	}

	cs.Identitiy = strings.TrimSpace(cs.Identitiy)
	if cs.Identitiy == "" {
		return errors.HandlerBadRequest("Invalid Course Identitiy")
	}
	return nil
}
