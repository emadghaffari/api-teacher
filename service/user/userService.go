package user

import (
	jst "github.com/emadghaffari/api-teacher/config/token"
	"github.com/emadghaffari/api-teacher/model/role"
	model "github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/api-teacher/service/token"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// Service var
	Service users = &user{}
)

type users interface {
	Login() (*jst.TokenDetails, errors.ResError)
	Register() (*jst.TokenDetails, errors.ResError)
}

type user struct{}

func (usr *user) Login() (*jst.TokenDetails, errors.ResError) {
	if err := model.Model.Login(); err != nil {
		return nil, err
	}

	return token.Service.Generate(*model.Model.Get())
}

func (usr *user) Register() (*jst.TokenDetails, errors.ResError) {
	if err := model.Model.Register(); err != nil {
		return nil, err
	}

	us := model.Model.Get()
	if err := role.Model.Assign(us.ID); err != nil {
		return nil, err
	}

	us.Role = &role.Role{ID: role.Model.Get().ID, Name: role.Model.Get().Name}

	return token.Service.Generate(*us)
}
