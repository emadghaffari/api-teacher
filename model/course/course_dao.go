package course

import (
	"fmt"
	"strings"

	"github.com/emadghaffari/api-teacher/database/postgres"
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/api-teacher/utils/date"
	"github.com/emadghaffari/res_errors/errors"
	log "github.com/sirupsen/logrus"
)

var (
	indexQuery  = "SELECT c.id,c.name,c.identitiy,c.valence, c.value, c.start_at, c.end_at, c.description,u.name,u.lname FROM courses as c INNER JOIN users as u ON c.user_id=u.id WHERE c.valence >= $1"
	storeQuery  = "INSERT INTO courses (user_id, name, identitiy, valence, value, start_at, end_at, description , created_at) VALUES ($1 , $2 , $3 , $4 , $5 , $6, $7 , $8 , $9 ) RETURNING id;"
	updateQuery = "UPDATE courses SET name = $1 , valence = $2 , value= $3 , start_at= $4 , end_at= $5, description= $6 WHERE identitiy = $7 AND user_id = $8 RETURNING id;"

	userCourseQuery = "INSERT into user_course (course_id,user_id,created_at) VALUES($1 , $2 , $3 ) ON CONFLICT DO NOTHING;"
	takeQuery       = "UPDATE courses SET valence = valence - 1 WHERE valence > 0 AND identitiy = $1 AND id NOT IN (select id from user_course WHERE user_id = $2 AND user_course.course_id = courses.id) RETURNING id,valence,user_id;"
)

// Index meth, get courses
func (u *Course) Index() (Courses, errors.ResError) {
	db := postgres.DB.GetDB()

	rows, err := db.Query(indexQuery, 0)
	if err != nil {
		log.Error(fmt.Sprintf("Error in Query for get a list of courses: %s", err))
		return nil, errors.HandlerInternalServerError(err.Error(), err)
	}

	items := make(Courses, 0)
	for rows.Next() {
		var cs Course
		var us user.User
		err := rows.Scan(&cs.ID,
			&cs.Name,
			&cs.Identitiy,
			&cs.Valence,
			&cs.Value,
			&cs.Start,
			&cs.End,
			&cs.Description,
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

// Store meth, store a new course
func (u *Course) Store() errors.ResError {
	db := postgres.DB.GetDB()

	if err := db.QueryRow(storeQuery, user.Model.Get().ID, u.Name, u.Identitiy, u.Valence, u.Value, u.Start, u.End, u.Description, date.GetNowString()).Scan(&u.ID); err != nil {
		log.Error(fmt.Sprintf("Error in store new course: %s", err.Error()))
		return errors.HandlerInternalServerError(err.Error(), err)
	}

	return nil
}

// Update meth, update a course
func (u *Course) Update() errors.ResError {
	db := postgres.DB.GetDB()

	err := db.QueryRow(updateQuery, u.Name, u.Valence, u.Value, u.Start, u.End, u.Description, u.Identitiy, user.Model.Get().ID).Err()
	if err != nil {
		log.Error(fmt.Sprintf("Error in update course: %s", err))
		return errors.HandlerInternalServerError(err.Error(), err)
	}

	return nil
}

// Take meth, take a course by user
func (u *Course) Take() errors.ResError {

	db := postgres.DB.GetDB()

	tx, err := db.Begin()
	if err != nil {
		log.Error(fmt.Sprintf("Error in BeginDB: %s", err))
		return errors.HandlerInternalServerError(err.Error(), err)
	}
	defer tx.Commit()

	var us user.User
	if err := tx.QueryRow(takeQuery, u.Identitiy, user.Model.Get().ID).Scan(&u.ID, &u.Valence, &us.ID); err != nil {
		if u.Valence == 0 {
			return errors.HandlerBadRequest(fmt.Sprintf("The valence of the course(%s) is over", u.Identitiy))
		}
		if !strings.Contains(err.Error(), "no rows in result set") {
			log.WithFields(log.Fields{
				"user":   user.Model.Get().ID,
				"course": u.Identitiy,
			}).Error(fmt.Sprintf("Error in take course for student: %s", err))
			tx.Rollback()
			return errors.HandlerInternalServerError(err.Error(), err)
		}

		return errors.HandlerBadRequest(fmt.Sprintf("This course(%s) has already been taken", u.Identitiy))
	}

	if err := tx.QueryRow(userCourseQuery, u.ID, user.Model.Get().ID, date.GetNowString()).Err(); err != nil {
		log.WithFields(log.Fields{
			"user":   user.Model.Get().ID,
			"course": u.Identitiy,
		}).Error(fmt.Sprintf("Error in store user_course: %s", err))
		tx.Rollback()
		return errors.HandlerInternalServerError(err.Error(), err)
	}
	u.Teacher = &us

	return nil
}
