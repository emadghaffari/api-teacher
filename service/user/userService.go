package user

import (
	"fmt"

	"github.com/emadghaffari/api-teacher/config/token"
	"github.com/emadghaffari/api-teacher/model/role"
	model "github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
	log "github.com/sirupsen/logrus"
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
	if err := us.Login(); err != nil {
		return nil, err
	}

	ts, _ := token.Conf.Generate(us)
	return ts, nil
}

func (usr *user) Register(us model.User) (*token.TokenDetails, errors.ResError) {
	rl := role.Role{Name: us.Role.Name}
	if err := rl.GetByName(); err != nil {
		log.Error(fmt.Sprintf("Role Not Found: %s", err))

		return nil, err
	}

	if err := us.Register(); err != nil {
		return nil, err
	}

	if err := rl.Assign(us.ID); err != nil {
		log.WithFields(log.Fields{
			"user_id": us.ID,
			"role_id": rl.ID,
		}).Error(fmt.Sprintf("Faild Assign Role To User: %s", err))
		return nil, err
	}

	us.Role.Name = rl.Name
	ts, err := token.Conf.Generate(us)
	if err != nil {
		return nil, errors.HandlerInternalServerError("Error in Generate Token: %s", err)
	}

	return ts, nil
}
