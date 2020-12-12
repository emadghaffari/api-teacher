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
	Index() ([]*model.Course, errors.ResError)
}

type course struct{}

func (cs *course) Index() (items []*model.Course, err errors.ResError) {
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
