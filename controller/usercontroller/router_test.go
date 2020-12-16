package usercontroller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/emadghaffari/api-teacher/config/token"
	model "github.com/emadghaffari/api-teacher/model/user"
	service "github.com/emadghaffari/api-teacher/service/user"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
	viper.Set("user.MinIdentitiy", 100000000000000)
	viper.Set("user.MaxIdentitiy", 999999999999999)
}

type usServiceMock struct {
	MockFuc func() (*token.TokenDetails, errors.ResError)
}

func (usr *usServiceMock) Login() (*token.TokenDetails, errors.ResError) {
	return usr.MockFuc()
}

func (usr *usServiceMock) Register() (*token.TokenDetails, errors.ResError) {
	return usr.MockFuc()
}

type usModelMock struct {
	MockFuc func() errors.ResError
}

func (usr *usModelMock) Register() errors.ResError         { return nil }
func (usr *usModelMock) Login() errors.ResError            { return nil }
func (usr *usModelMock) Set(*model.User)                   {}
func (usr *usModelMock) Get() *model.User                  { return nil }
func (usr *usModelMock) RegisterValidate() errors.ResError { return usr.MockFuc() }
func (usr *usModelMock) LoginValidate() errors.ResError    { return usr.MockFuc() }

func TestLogin(t *testing.T) {
	mocked := usServiceMock{}
	modelMocked := usModelMock{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	tests := []struct {
		step   string
		status int
		error  string
		params string
		mock   func() (*token.TokenDetails, errors.ResError)
	}{
		{
			step:   "a",
			params: `{"identitiy":"identitiy","password":"identitiy"}`,
			error:  "",
			status: http.StatusOK,
			mock: func() (*token.TokenDetails, errors.ResError) {
				return &token.TokenDetails{
					AccessToken:  "access",
					RefreshToken: "refresh",
				}, nil
			},
		},
		{
			step:   "b",
			params: ``,
			error:  "Invalid JSON Body.",
			status: http.StatusBadRequest,
			mock: func() (*token.TokenDetails, errors.ResError) {
				return &token.TokenDetails{
					AccessToken:  "access",
					RefreshToken: "refresh",
				}, nil
			},
		},
		{
			step:   "c",
			params: `{"identitiy":"identitiy"}`,
			error:  "Invalid Password.",
			status: http.StatusBadRequest,
			mock: func() (*token.TokenDetails, errors.ResError) {
				return &token.TokenDetails{
					AccessToken:  "access",
					RefreshToken: "refresh",
				}, nil
			},
		},
		{
			step:   "d",
			params: `{"identitiy":"identitiy","password":"identitiy"}`,
			error:  "Internal Test Error",
			status: http.StatusInternalServerError,
			mock: func() (*token.TokenDetails, errors.ResError) {
				return nil, errors.HandlerInternalServerError("Internal Test Error", nil)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			mocked.MockFuc = tc.mock
			service.Service = &mocked

			model.Model = &modelMocked

			// Mock HTTP Request and it's return
			c.Request, _ = http.NewRequest("POST", "/login", strings.NewReader(string(tc.params)))
			c.Request.Header.Add("Content-Type", gin.MIMEJSON)

			// call login func
			Router.Login(c)

			// if param is null and status is bad request (for check request)
			if tc.params == `` && tc.status == http.StatusBadRequest {
				if !strings.Contains(w.Body.String(), tc.error) {
					assert.Equal(t, tc.error, w.Body.String())
				}
			}

			// if param is not null and status is bad request (for check request)
			if tc.params != `` && tc.status == http.StatusBadRequest {
				if !strings.Contains(w.Body.String(), tc.error) {
					assert.Equal(t, tc.error, w.Body.String())
				}
			}

			// if status is internal server error (for check service)
			if tc.status == http.StatusInternalServerError {
				if !strings.Contains(w.Body.String(), tc.error) {
					assert.Equal(t, fmt.Sprintf("Expact: %s", tc.error), w.Body.String())
				}
			}

			// if status is OK
			if tc.status == http.StatusOK {
				assert.Equal(t, tc.status, w.Code)
				response, _ := tc.mock()
				if !strings.Contains(w.Body.String(), response.AccessToken) {
					assert.Equal(t, fmt.Sprintf("Expact %s in Body", response.AccessToken), w.Body.String())
				}
			}
		})
	}
}

func TestRegister(t *testing.T) {
	mocked := usServiceMock{}
	modelMocked := usModelMock{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	tests := []struct {
		step       string
		status     int
		error      string
		params     string
		mock       func() (*token.TokenDetails, errors.ResError)
		Validation func() errors.ResError
	}{
		{
			step:       "a",
			status:     http.StatusOK,
			error:      "",
			params:     `{"password": "123456798","name":"emad","lname":"ghaffari","role": {"name": "student"}}`,
			Validation: func() errors.ResError { return nil },
			mock: func() (*token.TokenDetails, errors.ResError) {
				return &token.TokenDetails{
					AccessToken:  "access",
					RefreshToken: "refresh",
				}, nil
			},
		},
		{
			step:       "b",
			status:     http.StatusBadRequest,
			error:      "Invalid JSON Body.",
			params:     ``,
			Validation: func() errors.ResError { return nil },
			mock: func() (*token.TokenDetails, errors.ResError) {
				return &token.TokenDetails{
					AccessToken:  "access",
					RefreshToken: "refresh",
				}, nil
			},
		},
		{
			step:       "c",
			status:     http.StatusBadRequest,
			error:      "Role is Empty",
			Validation: func() errors.ResError { return errors.HandlerBadRequest("Role is Empty.") },
			params:     `{"password": "123456798","name":"emad","lname":"ghaffari"}`,
			mock: func() (*token.TokenDetails, errors.ResError) {
				return &token.TokenDetails{
					AccessToken:  "access",
					RefreshToken: "refresh",
				}, nil
			},
		},
		{
			step:       "d",
			status:     http.StatusInternalServerError,
			error:      "Internal Test Error",
			params:     `{"password": "123456798","name":"emad","lname":"ghaffari","role": {"name": "student"}}`,
			Validation: func() errors.ResError { return nil },
			mock: func() (*token.TokenDetails, errors.ResError) {
				return nil, errors.HandlerInternalServerError("Internal Test Error", nil)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			mocked.MockFuc = tc.mock
			service.Service = &mocked

			modelMocked.MockFuc = tc.Validation
			model.Model = &modelMocked

			// Mock HTTP Request and it's return
			c.Request, _ = http.NewRequest("POST", "/register", strings.NewReader(string(tc.params)))
			c.Request.Header.Add("Content-Type", gin.MIMEJSON)
			Router.Register(c)

			// if param is null and status is bad request (for check request)
			if tc.params == `` && tc.status == http.StatusBadRequest {
				if !strings.Contains(w.Body.String(), tc.error) {
					assert.Equal(t, tc.error, w.Body.String())
				}
			}

			// if param is not null and status is bad request (for check request)
			if tc.params != `` && tc.status == http.StatusBadRequest {
				if !strings.Contains(w.Body.String(), tc.error) {
					assert.Equal(t, tc.error, w.Body.String())
				}
			}

			// if status is internal server error (for check service)
			if tc.status == http.StatusInternalServerError {
				if !strings.Contains(w.Body.String(), tc.error) {
					assert.Equal(t, fmt.Sprintf("Expact: %s", tc.error), w.Body.String())
				}
			}

			// if status is OK
			if tc.status == http.StatusOK {
				assert.Equal(t, tc.status, w.Code)
				response, _ := tc.mock()
				if !strings.Contains(w.Body.String(), response.AccessToken) {
					assert.Equal(t, fmt.Sprintf("Expact %s in Body", response.AccessToken), w.Body.String())
				}
			}

		})
	}

}

func ExampleRouter_Login() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mocked := usServiceMock{}
	mocked.MockFuc = func() (*token.TokenDetails, errors.ResError) {
		return &token.TokenDetails{
			AccessToken:  "access",
			RefreshToken: "refresh",
		}, nil
	}

	service.Service = &mocked

	// Mock HTTP Request and it's return
	c.Request, _ = http.NewRequest("POST", "/login", strings.NewReader(string(`{"identitiy":"identitiy","password":"identitiy"}`)))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)

	// call login func
	Router.Login(c)

	fmt.Println(w.Body.String())
	// Output: {"access_token":"access","refresh_token":"refresh"}
}

func ExampleRouter_Register() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mocked := usServiceMock{}
	mocked.MockFuc = func() (*token.TokenDetails, errors.ResError) {
		return &token.TokenDetails{
			AccessToken:  "access",
			RefreshToken: "refresh",
		}, nil
	}

	modelMocked := usModelMock{}
	modelMocked.MockFuc = func() errors.ResError { return nil }
	model.Model = &modelMocked

	service.Service = &mocked

	// Mock HTTP Request and it's return
	c.Request, _ = http.NewRequest("POST", "/register", strings.NewReader(string(`{"password": "123456798","name":"emad","lname":"ghaffari","role": {"name": "student"}}`)))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)

	// call login func
	Router.Register(c)

	fmt.Println(w.Body.String())
	// Output: {"access_token":"access","refresh_token":"refresh"}
}
