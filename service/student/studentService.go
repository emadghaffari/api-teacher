package student

import (
	"fmt"
	"time"

	"github.com/emadghaffari/api-teacher/database/redis"
	"github.com/emadghaffari/api-teacher/model/course"
	model "github.com/emadghaffari/api-teacher/model/student"
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// Service var
	Service students = &student{}
)

type students interface {
	Index() (items course.Courses, err errors.ResError)
}

type student struct{}

func (st *student) Index() (items course.Courses, err errors.ResError) {
	redis.DB.Get(fmt.Sprintf("user-courses-%d", user.Model.Get().ID), &items)
	if len(items) == 0 {
		items, err = model.Model.Index()
		if err != nil {
			return nil, err
		}
		redis.DB.Set(fmt.Sprintf("user-courses-%d", user.Model.Get().ID), items, time.Duration(time.Hour*72))
	}
	return items, nil
}
