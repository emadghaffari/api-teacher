package user

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/emadghaffari/api-teacher/config/token"
	model "github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/stretchr/testify/assert"
)

type mockModel struct {
	MockFuc func() errors.ResError
}

func (mc *mockModel) Register() errors.ResError {
	return mc.MockFuc()
}
func (mc *mockModel) Login() errors.ResError {
	return mc.MockFuc()
}
func (mc *mockModel) Set(m *model.User) {}
func (mc *mockModel) Get() *model.User  { return nil }

type mockToken struct {
	GenMockFuc func() (*token.TokenDetails, error)
	VerMockFuc func() (*token.AccessDetails, error)
}

func (mc *mockToken) Generate(user model.User) (*token.TokenDetails, error) {
	return mc.GenMockFuc()
}
func (mc *mockToken) VerifyToken(string) (*token.AccessDetails, error) {
	return mc.VerMockFuc()
}

func TestLogin(t *testing.T) {
	modelMocked := mockModel{}
	tokenMocked := mockToken{}
	tests := []struct {
		step       string
		status     int
		error      string
		param      model.User
		modelMock  func() errors.ResError
		genMockFuc func() (*token.TokenDetails, error)
		verMockFuc func() (*token.AccessDetails, error)
	}{
		{
			step:       "a",
			status:     http.StatusOK,
			error:      "",
			param:      model.User{},
			modelMock:  func() errors.ResError { return nil },
			genMockFuc: func() (*token.TokenDetails, error) { return nil, nil },
			verMockFuc: func() (*token.AccessDetails, error) { return nil, nil },
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			// model mock
			modelMocked.MockFuc = tc.modelMock
			model.Model = &modelMocked

			// token mock
			tokenMocked.GenMockFuc = tc.genMockFuc
			tokenMocked.VerMockFuc = tc.verMockFuc

			tdt, err := Service.Login(tc.param)
			if err != nil {
				assert.Equal(t, "", "")
			}
			fmt.Println(tdt)
		})
	}
}
