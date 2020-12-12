package course

import (
	"fmt"

	"github.com/emadghaffari/api-teacher/database/postgres"
	"github.com/emadghaffari/res_errors/errors"
	log "github.com/sirupsen/logrus"
)

var (
	indexQuery = "SELECT c.id,c.name,c.identitiy,c.valence,c.time,u.name FROM courses as c INNER JOIN users as u ON c.user_id=u.id WHERE c.valence >= $1"
)

// Index meth, get courses
func (u *Course) Index() ([]*Course, errors.ResError) {
	db := postgres.DB.GetDB()

	rows, err := db.Query(indexQuery, 0)
	if err != nil {
		log.Error(fmt.Sprintf("Error in Query for get a list of courses: %s", err))
		return nil, errors.HandlerInternalServerError(err.Error(), err)
	}

	items := make([]*Course, 0)
	for rows.Next() {
		var cs Course
		err := rows.Scan(&cs.ID, &cs.Name, &cs.Identitiy, &cs.Valence, &cs.Time, &cs.Teacher)
		if err != nil {
			log.Error(fmt.Sprintf("Error in Scan rows for get a list of courses: \ncourse:%v \nerror: %s", cs, err))
			return nil, errors.HandlerInternalServerError(err.Error(), err)
		}
		items = append(items, &cs)
	}

	return items, nil
}

// Store meth, store a new course
func (u *Course) Store() errors.ResError {

	return nil
}

// Update meth, update a course
func (u *Course) Update() errors.ResError {

	return nil
}

// Take meth, take a course by user
func (u *Course) Take() errors.ResError {

	return nil
}
