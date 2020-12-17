package student

import (
	"strings"
	"testing"
	"time"

	rd "github.com/emadghaffari/api-teacher/database/redis"
	"github.com/emadghaffari/api-teacher/model/course"
	std "github.com/emadghaffari/api-teacher/model/student"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

type rdMock struct {
	MockFunc func() error
}

func (r *rdMock) New() {}
func (r *rdMock) GetDB() *redis.Client {
	return nil
}

func (r *rdMock) Get(key string, dest interface{}) error {
	return r.MockFunc()
}

func (r *rdMock) Set(key string, value interface{}, duration time.Duration) error {
	return r.MockFunc()
}

func (r *rdMock) Del(key ...string) error {
	return r.MockFunc()
}

type stdMock struct {
	MockFunc func() (course.Courses, errors.ResError)
}

func (te *stdMock) Index() (course.Courses, errors.ResError) { return te.MockFunc() }

func TestIndex(t *testing.T) {
	redi := rdMock{}
	stdh := stdMock{}

	tests := []struct {
		step      string
		error     string
		redisMock func() error
		stdMock   func() (course.Courses, errors.ResError)
	}{
		{
			step:      "a",
			error:     "",
			redisMock: func() error { return nil },
			stdMock:   func() (course.Courses, errors.ResError) { return []*course.Course{}, nil },
		},
		{
			step:      "a",
			error:     "TEST",
			redisMock: func() error { return nil },
			stdMock:   func() (course.Courses, errors.ResError) { return []*course.Course{}, errors.HandlerBadRequest("TEST") },
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			redi.MockFunc = tc.redisMock
			stdh.MockFunc = tc.stdMock

			rd.DB = &redi
			std.Model = &stdh

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
