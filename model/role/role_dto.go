package role

var (
	// Model for role
	Model role = &Role{}
)

type role interface{}

// Role struct teacher, student
type Role struct {
	ID   int64  `json:"-,omitempty"`
	Name string `json:"name,omitempty"`
}
