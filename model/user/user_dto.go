package user

import (
	"strings"

	"github.com/emadghaffari/api-teacher/model/role"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// Model for User
	Model user = &User{}
)

type user interface {
	Register() errors.ResError
	Login() errors.ResError
	Set(*User)
	Get() *User
	RegisterValidate() errors.ResError
	LoginValidate() errors.ResError
}

// User struct
type User struct {
	ID        int64      `json:"id,omitempty"`
	FirstName string     `json:"name,omitempty"`
	LastName  string     `json:"lname,omitempty"`
	Identitiy string     `json:"identitiy,omitempty"`
	CreatedAt string     `json:"created_at,omitempty"`
	Password  string     `json:"password,omitempty"`
	Role      *role.Role `json:"role,omitempty"`
}

// RegisterValidate user
func (us *User) RegisterValidate() errors.ResError {
	if us.FirstName == "" {
		return errors.HandlerBadRequest("FirstName is Empty.")
	}
	if us.LastName == "" {
		return errors.HandlerBadRequest("LastName is Empty.")
	}
	if us.Password == "" {
		return errors.HandlerBadRequest("Password is Empty.")
	}

	if us.Role == nil {
		return errors.HandlerBadRequest("Role is Empty.")
	}

	rl := role.Role{Name: us.Role.Name}
	rl.Name = strings.TrimSpace(rl.Name)
	if err := rl.GetByName(); err != nil {
		return err
	}
	role.Model.Set(&rl)

	return nil
}

// LoginValidate user
func (us *User) LoginValidate() errors.ResError {
	us.Password = strings.TrimSpace(us.Password)
	if us.Password == "" {
		return errors.HandlerBadRequest("Invalid Password.")
	}

	if us.Identitiy == "" {
		return errors.HandlerBadRequest("Invalid Identitiy.")
	}

	return nil
}

// Set user
func (us *User) Set(u *User) {
	*us = *u
}

// Get user
func (us *User) Get() *User {
	return us
}
