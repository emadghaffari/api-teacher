package course

import (
	"time"

	"github.com/emadghaffari/api-teacher/database/redis"
	model "github.com/emadghaffari/api-teacher/model/course"
	"github.com/emadghaffari/res_errors/errors"
)

var (
	// Service var
	Service courses = &course{}
)

type courses interface {
	Index() (model.Courses, errors.ResError)
	Store(item *model.Course) errors.ResError
	Update(item *model.Course) errors.ResError
	Take(item *model.Course) errors.ResError
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

func (cs *course) Store(item *model.Course) errors.ResError {

	if err := item.Store(); err != nil {
		return err
	}

	return nil
}

func (cs *course) Update(item *model.Course) errors.ResError {

	if err := item.Update(); err != nil {
		return err
	}

	return nil
}

func (cs *course) Take(item *model.Course) errors.ResError {

	if err := item.Take(); err != nil {
		return err
	}

	return nil
}
