package course

import (
	"regexp"
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
	Store() errors.ResError
	Update() errors.ResError
	Take() errors.ResError
	StoreValidate() errors.ResError
	UpdateValidate() errors.ResError
	TakeValidate() errors.ResError
	Set(u *Course)
	Get() *Course
}

// Course struct
type Course struct {
	ID          int64      `json:"id,omitempty"`
	Teacher     *user.User `json:"teacher,omitempty"`
	Name        string     `json:"name,omitempty"`
	Identitiy   string     `json:"identitiy,omitempty"`
	Valence     uint32     `json:"valence,omitempty"`
	Value       uint32     `json:"value,omitempty"`
	Start       string     `json:"start_at,omitempty"`
	End         string     `json:"end_at,omitempty"`
	Description string     `json:"Description,omitempty"`
}

// Courses var list of Course
type Courses []*Course

// StoreValidate meth, validate items before store
func (cs *Course) StoreValidate() errors.ResError {
	if cs.Name == "" {
		return errors.HandlerBadRequest("Invalid Course Name")
	}
	re := regexp.MustCompile("([0-2][0-9]):([0-5][0-9]):([0-5][0-9])")
	if !re.MatchString(cs.Start) {
		return errors.HandlerBadRequest("Invalid Course Start time")
	}
	if !re.MatchString(cs.End) {
		return errors.HandlerBadRequest("Invalid Course End time")
	}
	if cs.Description == "" {
		return errors.HandlerBadRequest("Invalid Course Description")
	}
	if cs.Valence <= 1 {
		return errors.HandlerBadRequest("invalid Course Valence")
	}
	if cs.Value < 1 || cs.Value >= 4 {
		return errors.HandlerBadRequest("invalid Course Value")
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
	re := regexp.MustCompile("([0-2][0-9]):([0-5][0-9]):([0-5][0-9])")
	if !re.MatchString(cs.Start) {
		return errors.HandlerBadRequest("Invalid Course Start time")
	}
	if !re.MatchString(cs.End) {
		return errors.HandlerBadRequest("Invalid Course End time")
	}
	if cs.Description == "" {
		return errors.HandlerBadRequest("Invalid Course Description")
	}
	if cs.Valence <= 1 {
		return errors.HandlerBadRequest("invalid Course Valence")
	}
	if cs.Value <= 1 {
		return errors.HandlerBadRequest("invalid Course Value")
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

// Set user
func (cs *Course) Set(u *Course) {
	*cs = *u
}

// Get Course
func (cs *Course) Get() *Course {
	return cs
}
