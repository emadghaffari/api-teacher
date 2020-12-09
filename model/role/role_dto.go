package role

var (
	// Model for role
	Model role = &Role{}
)

type role interface{}

// Role struct teacher, student
type Role struct {
	ID   int64  `json:"-"`
	Name string `json:"name"`
}
