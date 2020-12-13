package teacher

import (
	"fmt"
	"time"

	"github.com/emadghaffari/api-teacher/database/redis"
	"github.com/emadghaffari/api-teacher/model/course"
	tec "github.com/emadghaffari/api-teacher/model/teacher"
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// Service var
	Service teachers = &teacher{}
)

type teachers interface {
	Index() (items course.Courses, err errors.ResError)
}

type teacher struct{}

func (te *teacher) Index() (items course.Courses, err errors.ResError) {
	redis.DB.Get(fmt.Sprintf("teacher-courses-%d", user.Model.Get().ID), &items)
	if len(items) == 0 {
		items, err = tec.Model.Index()
		if err != nil {
			return nil, err
		}
		redis.DB.Set(fmt.Sprintf("teacher-courses-%d", user.Model.Get().ID), items, time.Duration(time.Hour*72))
	}
	return items, nil
}
