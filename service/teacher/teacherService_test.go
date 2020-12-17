package teacher

import (
	"strings"
	"testing"
	"time"

	rd "github.com/emadghaffari/api-teacher/database/redis"
	"github.com/emadghaffari/api-teacher/model/course"
	tec "github.com/emadghaffari/api-teacher/model/teacher"
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

type techMock struct {
	MockFunc func() (course.Courses, errors.ResError)
}

func (te *techMock) Index() (course.Courses, errors.ResError) { return te.MockFunc() }

func TestIndex(t *testing.T) {
	redi := rdMock{}
	tech := techMock{}

	tests := []struct {
		step      string
		error     string
		redisMock func() error
		techMock  func() (course.Courses, errors.ResError)
	}{
		{
			step:      "a",
			error:     "",
			redisMock: func() error { return nil },
			techMock:  func() (course.Courses, errors.ResError) { return []*course.Course{}, nil },
		},
		{
			step:      "a",
			error:     "TEST",
			redisMock: func() error { return nil },
			techMock:  func() (course.Courses, errors.ResError) { return []*course.Course{}, errors.HandlerBadRequest("TEST") },
		},
	}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			redi.MockFunc = tc.redisMock
			tech.MockFunc = tc.techMock

			rd.DB = &redi
			tec.Model = &tech

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
