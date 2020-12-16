package role

import "github.com/emadghaffari/res_errors/errors"

var (
	// Model for role
	Model role = &Role{}
)

type role interface {
	Set(*Role)
	Get() *Role
	Insert() errors.ResError
	GetByName() errors.ResError
	Assign(id int64) errors.ResError
}

// Role struct teacher, student
type Role struct {
	ID   int64  `json:"-,omitempty"`
	Name string `json:"name,omitempty"`
}

// Set meth, for set a role for client
func (r *Role) Set(rl *Role) {
	*r = *rl
}

// Get meth, for get client role
func (r *Role) Get() *Role {
	return r
}
