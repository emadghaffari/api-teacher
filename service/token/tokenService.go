package token

import (
	jst "github.com/emadghaffari/api-teacher/config/token"
	model "github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// Service var
	Service tokens = &token{}
)

type tokens interface {
	Generate(model.User) (*jst.TokenDetails, errors.ResError)
}

type token struct{}

func (t *token) Generate(u model.User) (*jst.TokenDetails, errors.ResError) {
	ts, err := jst.Conf.Generate(u)
	if err != nil {
		return nil, errors.HandlerInternalServerError("Error in Generate Token: %s", err)
	}

	return ts, nil
}
