package student

import (
	"fmt"

	"github.com/emadghaffari/api-teacher/database/postgres"
	"github.com/emadghaffari/api-teacher/model/course"
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/res_errors/errors"
	log "github.com/sirupsen/logrus"
)

var (
	indexQuery = `SELECT c.name,c.identitiy,c.time,us2.name,us2.lname from courses as c 
					LEFT JOIN user_course as us 
						ON c.id=us.course_id 			
					LEFT JOIN users as cp 
						ON us.user_id = cp.id
					LEFT JOIN users AS us2 
						ON c.user_id  = us2.id
					
					where cp.id= $1`
)

// Index meth, get all courses for student
func (u *Student) Index() (course.Courses, errors.ResError) {
	db := postgres.DB.GetDB()

	rows, err := db.Query(indexQuery, user.Model.Get().ID)
	if err != nil {
		log.Error(fmt.Sprintf("Error in Query for get a list of courses: %s", err))
		return nil, errors.HandlerInternalServerError(err.Error(), err)
	}

	items := make(course.Courses, 0)
	for rows.Next() {
		var cs course.Course
		var us user.User
		err := rows.Scan(
			&cs.Name,
			&cs.Identitiy,
			&cs.Time,
			&us.FirstName,
			&us.LastName,
		)
		cs.Teacher = &us
		if err != nil {
			log.Error(fmt.Sprintf("Error in Scan rows for get a list of courses: \ncourse:%v \nerror: %s", cs, err))
			return nil, errors.HandlerInternalServerError(err.Error(), err)
		}
		items = append(items, &cs)
	}

	return items, nil
}
