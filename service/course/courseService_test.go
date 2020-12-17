package course

import (
	"strings"
	"testing"
	"time"

	rd "github.com/emadghaffari/api-teacher/database/redis"
	model "github.com/emadghaffari/api-teacher/model/course"
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

// redis mock
type rdMock struct {
	MockFunc func() error
}

func (r *rdMock) New()                                               {}
func (r *rdMock) GetDB() *redis.Client                               { return nil }
func (r *rdMock) Get(k string, d interface{}) error                  { return r.MockFunc() }
func (r *rdMock) Set(k string, v interface{}, d time.Duration) error { return r.MockFunc() }
func (r *rdMock) Del(k ...string) error                              { return r.MockFunc() }

// curse mock
type curMock struct {
	MockFunc   func() errors.ResError
	MockCourse func() (model.Courses, errors.ResError)
}

func (c *curMock) Index() (model.Courses, errors.ResError) { return c.MockCourse() }
func (c *curMock) Store() errors.ResError                  { return c.MockFunc() }
func (c *curMock) Update() errors.ResError                 { return c.MockFunc() }
func (c *curMock) Take() errors.ResError                   { return c.MockFunc() }
func (c *curMock) StoreValidate() errors.ResError          { return c.MockFunc() }
func (c *curMock) UpdateValidate() errors.ResError         { return c.MockFunc() }
func (c *curMock) TakeValidate() errors.ResError           { return c.MockFunc() }
func (c *curMock) Set(u *model.Course)                     {}
func (c *curMock) Get() *model.Course                      { return &model.Course{ID: 1, Teacher: &user.User{ID: 1}} }

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
func (mc *mockModel) Set(m *user.User)                  {}
func (mc *mockModel) Get() *user.User                   { return &user.User{ID: 1} }
func (mc *mockModel) RegisterValidate() errors.ResError { return nil }
func (mc *mockModel) LoginValidate() errors.ResError    { return nil }

func TestIndex(t *testing.T) {
	redi := rdMock{}
	stdh := curMock{}

	tests := []struct {
		step      string
		error     string
		redisMock func() error
		stdMock   func() (model.Courses, errors.ResError)
	}{
		{
			step:      "a",
			error:     "",
			redisMock: func() error { return nil },
			stdMock:   func() (model.Courses, errors.ResError) { return []*model.Course{}, nil },
		},
		{
			step:      "a",
			error:     "TEST",
			redisMock: func() error { return nil },
			stdMock:   func() (model.Courses, errors.ResError) { return []*model.Course{}, errors.HandlerBadRequest("TEST") },
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			stdh.MockCourse = tc.stdMock

			redi.MockFunc = tc.redisMock
			rd.DB = &redi
			model.Model = &stdh

			_, err := Service.Index()
			if err != nil {
				assert.Equal(t, "", "")
				if !strings.Contains(tc.error, err.Message()) {
					assert.Equal(t, tc.error, err.Message())
				}
			}
		})
	}
}

func TestStore(t *testing.T) {
	stdh := curMock{}
	redi := rdMock{}

	tests := []struct {
		step      string
		error     string
		MockFunc  func() errors.ResError
		redisMock func() error
		stdMock   func() (model.Courses, errors.ResError)
	}{
		{
			step:      "a",
			error:     "",
			MockFunc:  func() errors.ResError { return nil },
			redisMock: func() error { return nil },
			stdMock:   func() (model.Courses, errors.ResError) { return []*model.Course{}, nil },
		},
		{
			step:      "a",
			error:     "TEST",
			MockFunc:  func() errors.ResError { return errors.HandlerBadRequest("TEST") },
			redisMock: func() error { return nil },
			stdMock:   func() (model.Courses, errors.ResError) { return nil, errors.HandlerBadRequest("TEST") },
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			stdh.MockCourse = tc.stdMock
			stdh.MockFunc = tc.MockFunc
			model.Model = &stdh

			redi.MockFunc = tc.redisMock
			rd.DB = &redi

			err := Service.Store()
			if err != nil {
				if !strings.Contains(tc.error, err.Message()) {
					assert.Equal(t, tc.error, err.Message())
				}
			}
		})
	}

}

func TestUpdate(t *testing.T) {
	stdh := curMock{}
	redi := rdMock{}
	tests := []struct {
		step      string
		error     string
		MockFunc  func() errors.ResError
		redisMock func() error
		stdMock   func() (model.Courses, errors.ResError)
	}{
		{
			step:      "a",
			error:     "",
			MockFunc:  func() errors.ResError { return nil },
			redisMock: func() error { return nil },
			stdMock:   func() (model.Courses, errors.ResError) { return []*model.Course{}, nil },
		},
		{
			step:      "a",
			error:     "TEST",
			MockFunc:  func() errors.ResError { return errors.HandlerBadRequest("TEST") },
			redisMock: func() error { return nil },
			stdMock:   func() (model.Courses, errors.ResError) { return nil, errors.HandlerBadRequest("TEST") },
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			stdh.MockCourse = tc.stdMock
			stdh.MockFunc = tc.MockFunc
			model.Model = &stdh

			redi.MockFunc = tc.redisMock
			rd.DB = &redi

			err := Service.Update()
			if err != nil {
				if !strings.Contains(tc.error, err.Message()) {
					assert.Equal(t, tc.error, err.Message())
				}
			}
		})
	}

}
func TestTake(t *testing.T) {
	stdh := curMock{}
	redi := rdMock{}
	usr := mockModel{}

	tests := []struct {
		step      string
		error     string
		MockFunc  func() errors.ResError
		redisMock func() error
		stdMock   func() (model.Courses, errors.ResError)
	}{
		{
			step:      "a",
			error:     "",
			MockFunc:  func() errors.ResError { return nil },
			redisMock: func() error { return nil },
			stdMock:   func() (model.Courses, errors.ResError) { return []*model.Course{}, nil },
		},
		{
			step:      "a",
			error:     "TEST",
			MockFunc:  func() errors.ResError { return errors.HandlerBadRequest("TEST") },
			redisMock: func() error { return nil },
			stdMock:   func() (model.Courses, errors.ResError) { return nil, errors.HandlerBadRequest("TEST") },
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			stdh.MockCourse = tc.stdMock
			stdh.MockFunc = tc.MockFunc
			model.Model = &stdh

			usr.MockFuc = tc.MockFunc
			user.Model = &usr

			redi.MockFunc = tc.redisMock
			rd.DB = &redi

			err := Service.Take()
			if err != nil {
				if !strings.Contains(tc.error, err.Message()) {
					assert.Equal(t, tc.error, err.Message())
				}
			}
		})
	}

}
