package token

import (
	"net/http"
	"testing"

	jst "github.com/emadghaffari/api-teacher/config/token"
	model "github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/stretchr/testify/assert"
)

type tMock struct {
	GenMock func() (*jst.TokenDetails, error)
	VerMock func() (*jst.AccessDetails, error)
}

func (t *tMock) Generate(user model.User) (*jst.TokenDetails, error) { return t.GenMock() }
func (t *tMock) VerifyToken(string) (*jst.AccessDetails, error)      { return t.VerMock() }

func TestGenerateSuccess(t *testing.T) {
	mock := tMock{}
	mock.GenMock = func() (*jst.TokenDetails, error) { return &jst.TokenDetails{AccessToken: "TEST"}, nil }
	jst.Conf = &mock
	dt, err := Service.Generate(model.User{})
	if err != nil {
		assert.Equal(t, nil, err.Message())
	}
	assert.Equal(t, nil, err)
	assert.Equal(t, "TEST", dt.AccessToken)
}
func TestGenerateFail(t *testing.T) {
	mock := tMock{}
	mock.GenMock = func() (*jst.TokenDetails, error) {
		return &jst.TokenDetails{AccessToken: ""}, errors.HandlerInternalServerError("ERROR", nil)
	}
	jst.Conf = &mock
	_, err := Service.Generate(model.User{})
	if err == nil {
		assert.Equal(t, "Error in Generate Token", err.Message())
	}
	assert.Equal(t, "Error in Generate Token", err.Message())
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}
