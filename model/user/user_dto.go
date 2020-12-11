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
	Set(*User)
	Get() *User
}

// User struct
type User struct {
	ID        int64     `json:"-"`
	FirstName string    `json:"name"`
	LastName  string    `json:"lname"`
	Identitiy string    `json:"identitiy"`
	CreatedAt string    `json:"created_at"`
	Password  string    `json:"password"`
	Role      role.Role `json:"role"`
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
	if us.Role.Name == "" {
		return errors.HandlerBadRequest("Role is Empty.")
	}
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
