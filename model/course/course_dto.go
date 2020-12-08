package course

var (
	// Model for courses
	Model course = &Courses{}
)

type course interface{}

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
type Courses []*Course
