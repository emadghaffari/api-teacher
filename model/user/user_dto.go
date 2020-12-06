package user

var (
	// Model for User
	Model user = &User{}
)

type user interface{}

// User struct
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"lname"`
	Identitiy string `json:"identitiy"`
	CreatedAt string `json:"created_at"`
	Password  string `json:"_"`
	Role      role
}

// Roles: teacher, student
type role struct {
	RoleID int64  `json:"id"`
	Name   string `json:"name"`
}
