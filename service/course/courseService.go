package course

import (
	"fmt"
	"time"

	"github.com/emadghaffari/api-teacher/database/redis"
	model "github.com/emadghaffari/api-teacher/model/course"
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// Service var
	Service courses = &course{}
)

type courses interface {
	Index() (model.Courses, errors.ResError)
	Store() errors.ResError
	Update() errors.ResError
	Take() errors.ResError
}

type course struct{}

func (cs *course) Index() (items model.Courses, err errors.ResError) {
	redis.DB.Get("courses", &items)
	if len(items) == 0 {
		items, err = model.Model.Index()
		if err != nil {
			return nil, err
		}
		redis.DB.Set("courses", items, time.Duration(time.Hour*72))
	}
	return items, nil
}

func (cs *course) Store() errors.ResError {

	if err := model.Model.Store(); err != nil {
		return err
	}

	redis.DB.Del("courses", fmt.Sprintf("teacher-courses-%d", user.Model.Get().ID))
	return nil
}

func (cs *course) Update() errors.ResError {

	if err := model.Model.Update(); err != nil {
		return err
	}

	redis.DB.Del("courses", fmt.Sprintf("teacher-courses-%d", user.Model.Get().ID))
	return nil
}

func (cs *course) Take() errors.ResError {

	if err := model.Model.Take(); err != nil {
		return err
	}

	redis.DB.Del(fmt.Sprintf("user-courses-%d", user.Model.Get().ID), fmt.Sprintf("teacher-courses-%d", model.Model.Get().Teacher.ID))
	return nil
}
