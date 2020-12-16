package user

import (
	"testing"

	"github.com/emadghaffari/api-teacher/config/token"
	"github.com/emadghaffari/api-teacher/model/role"
	model "github.com/emadghaffari/api-teacher/model/user"
	srv "github.com/emadghaffari/api-teacher/service/token"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/stretchr/testify/assert"
)

// mock for user model
type mockModel struct {
	MockFuc func() errors.ResError
}

func (mc *mockModel) Register() errors.ResError {
	return mc.MockFuc()
}
func (mc *mockModel) Login() errors.ResError {
	return mc.MockFuc()
}
func (mc *mockModel) Set(m *model.User)                 {}
func (mc *mockModel) Get() *model.User                  { return &model.User{} }
func (mc *mockModel) RegisterValidate() errors.ResError { return nil }
func (mc *mockModel) LoginValidate() errors.ResError    { return nil }

// mock for token
type mockToken struct {
	GenMockFuc func() (*token.TokenDetails, errors.ResError)
}

func (mc *mockToken) Generate(model.User) (*token.TokenDetails, errors.ResError) {
	return mc.GenMockFuc()
}

// mock for role
type mockRole struct {
	MockFuc func() *role.Role
	MockErr func() errors.ResError
}

func (mc *mockRole) Set(rl *role.Role)               {}
func (mc *mockRole) Get() *role.Role                 { return mc.MockFuc() }
func (mc *mockRole) Insert() errors.ResError         { return mc.MockErr() }
func (mc *mockRole) GetByName() errors.ResError      { return mc.MockErr() }
func (mc *mockRole) Assign(id int64) errors.ResError { return mc.MockErr() }

func TestRegister(t *testing.T) {
	// mock user model AND tokenService
	modelMocked := mockModel{}
	tokenMocked := mockToken{}
	roleMocked := mockRole{}
	tests := []struct {
		step        string
		status      int
		error       string
		MockFuc     func() errors.ResError
		genMockFuc  func() (*token.TokenDetails, errors.ResError)
		roleMockFuc func() *role.Role
		MockErr     func() errors.ResError
	}{
		{
			step:        "a",
			error:       "",
			MockFuc:     func() errors.ResError { return nil },
			genMockFuc:  func() (*token.TokenDetails, errors.ResError) { return &token.TokenDetails{}, nil },
			roleMockFuc: func() *role.Role { return &role.Role{ID: 1, Name: ""} },
			MockErr:     func() errors.ResError { return nil },
		},
		{
			step:        "b",
			error:       "BAD REQUEST",
			MockFuc:     func() errors.ResError { return errors.HandlerBadRequest("BAD REQUEST") },
			genMockFuc:  func() (*token.TokenDetails, errors.ResError) { return &token.TokenDetails{}, nil },
			roleMockFuc: func() *role.Role { return &role.Role{ID: 1, Name: ""} },
			MockErr:     func() errors.ResError { return nil },
		},
		{
			step:        "c",
			error:       "BAD REQUEST",
			MockFuc:     func() errors.ResError { return nil },
			genMockFuc:  func() (*token.TokenDetails, errors.ResError) { return &token.TokenDetails{}, nil },
			roleMockFuc: func() *role.Role { return &role.Role{ID: 0, Name: ""} },
			MockErr:     func() errors.ResError { return errors.HandlerBadRequest("BAD REQUEST") },
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {

			modelMocked.MockFuc = tc.MockFuc
			model.Model = &modelMocked

			// mock token generate
			tokenMocked.GenMockFuc = tc.genMockFuc
			srv.Service = &tokenMocked

			// mock role
			roleMocked.MockFuc = tc.roleMockFuc
			roleMocked.MockErr = tc.MockErr
			role.Model = &roleMocked

			Service.Register()

		})
	}
}

func TestLogin(t *testing.T) {
	// mock user model AND tokenService
	modelMocked := mockModel{}
	tokenMocked := mockToken{}

	tests := []struct {
		step       string
		status     int
		error      string
		param      model.User
		MockFuc    func() errors.ResError
		genMockFuc func() (*token.TokenDetails, errors.ResError)
	}{
		{
			step:       "a",
			error:      "",
			param:      model.User{},
			MockFuc:    func() errors.ResError { return nil },
			genMockFuc: func() (*token.TokenDetails, errors.ResError) { return &token.TokenDetails{}, nil },
		},
		{
			step:       "a",
			error:      "BAD REQUEST",
			param:      model.User{},
			MockFuc:    func() errors.ResError { return errors.HandlerBadRequest("BAD REQUEST") },
			genMockFuc: func() (*token.TokenDetails, errors.ResError) { return &token.TokenDetails{}, nil },
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {

			modelMocked.MockFuc = tc.MockFuc
			model.Model = &modelMocked

			// mock token generate
			tokenMocked.GenMockFuc = tc.genMockFuc
			srv.Service = &tokenMocked

			jwt, err := Service.Login()
			if err == nil {
				if tc.error == "" {
					assert.Equal(t, jwt.AccessToken, "")
				} else {
					assert.Fail(t, err.Error())
				}
			}
			if err != nil {
				if tc.error != "" {
					assert.Equal(t, err.Message(), tc.error)
				} else {
					assert.Fail(t, err.Message())
				}
			}

		})
	}
}
