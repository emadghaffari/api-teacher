package user

import (
	"github.com/emadghaffari/api-teacher/config/token"
	model "github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// Service var
	Service users = &user{}
)

type users interface {
	Login(model.User) (*token.TokenDetails, errors.ResError)
	Register(model.User) (*token.TokenDetails, errors.ResError)
}

type user struct{}

func (usr *user) Login(us model.User) (*token.TokenDetails, errors.ResError) {
	err := us.Login()
	if err != nil {
		return nil, err
	}
	ts, _ := token.Conf.Generate(us)
	return ts, nil
}

func (usr *user) Register(us model.User) (*token.TokenDetails, errors.ResError) {
	err := us.Register()
	if err != nil {
		return nil, err
	}
	ts, _ := token.Conf.Generate(us)
	return ts, nil
}
