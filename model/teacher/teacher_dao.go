package teacher

import (
	"fmt"

	"github.com/emadghaffari/api-teacher/database/postgres"
	"github.com/emadghaffari/api-teacher/model/course"
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
	log "github.com/sirupsen/logrus"
)

var (
	indexQuery = `SELECT name,identitiy,valence,time from courses where user_id = $1;`
)

// Index meth, get all courses for student
func (ts *Teacher) Index() (course.Courses, errors.ResError) {
	db := postgres.DB.GetDB()

	rows, err := db.Query(indexQuery, user.Model.Get().ID)
	if err != nil {
		log.Error(fmt.Sprintf("Error in Query for get a list of courses: %s", err))
		return nil, errors.HandlerInternalServerError(err.Error(), err)
	}

	items := make(course.Courses, 0)
	for rows.Next() {
		var cs course.Course
		err := rows.Scan(
			&cs.Name,
			&cs.Identitiy,
			&cs.Valence,
			&cs.Time,
		)
		if err != nil {
			log.Error(fmt.Sprintf("Error in Scan rows for get a list of courses: \ncourse:%v \nerror: %s", cs, err))
			return nil, errors.HandlerInternalServerError(err.Error(), err)
		}
		items = append(items, &cs)
	}

	return items, nil
}
