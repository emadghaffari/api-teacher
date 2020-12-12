package course

import "github.com/emadghaffari/res_errors/errors"

var (
	// Model for Course
	Model course = &Course{}
)

type course interface {
	Index() ([]*Course, errors.ResError)
}

// Course struct
type Course struct {
	ID        int64  `json:"id"`
	Teacher   string `json:"teacher"`
	Name      string `json:"name"`
	Identitiy string `json:"identitiy"`
	Valence   uint32 `json:"valence"`
	Time      string `json:"time"`
}

// Courses list of Course
type Courses []Course
