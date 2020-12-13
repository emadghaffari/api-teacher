package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/emadghaffari/api-teacher/config/vip"
	"github.com/emadghaffari/api-teacher/database/postgres"
	"github.com/emadghaffari/api-teacher/model/role"
	"github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/api-teacher/utils/date"
	"github.com/emadghaffari/api-teacher/utils/hash"
)

type seed struct {
	users []user.User
}

var (
	roles = []string{
		"student",
		"teacher",
	}
	names = []string{
		"Emad",
		"Reza",
		"Ahmad",
		"Mina",
		"Jamshid",
		"Jafar",
		"Javad",
	}
	lnames = []string{
		"Ahmadi",
		"Rezaza",
		"Solmani",
		"Jafarian",
		"soltani",
	}
	courses = []string{
		"Diploma in Rural Healthcare",
		"Certificate in General Duty Assistant",
		"Diploma in Construction Management",
		"Diploma in Gemology",
		"Diploma in Photography",
		"Diploma in Food and Beverage Services",
	}
	times = []string{
		"2 months - 1 year",
		"1 year	",
		"1 - 1.5 year",
		"1 year",
		"3 years",
		"2 - 6 months",
	}
)

// seed tables
func main() {
	vip.Conf.New()
	postgres.DB.New()
	s := seed{}
	s.Roles()
	s.Users()
	s.Courses()
	s.UserRoles()
	s.UserCourses()
}

func (s *seed) Roles() {
	db := postgres.DB.GetDB()
	db.QueryRow("TRUNCATE TABLE roles RESTART IDENTITY CASCADE;")

	sqlStr := "INSERT INTO roles (name) VALUES"
	data := make([]string, 0)
	for _, dt := range roles {
		data = append(data, dt)
	}

	for _, row := range data {
		sqlStr += fmt.Sprintf("('%s'),", row)
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	err := db.QueryRow(sqlStr + ";")
	if err.Err() != nil {
		fmt.Println(err.Err())
	}
}

func (s *seed) Users() []user.User {
	rand.Seed(time.Now().UnixNano())
	db := postgres.DB.GetDB()
	db.QueryRow("TRUNCATE TABLE users RESTART IDENTITY CASCADE;")
	sqlStr := "INSERT INTO users (name, lname, identitiy, password,created_at) VALUES"
	for i := 1; i <= 20; i++ {
		name := names[rand.Intn(len(names))]
		lname := lnames[rand.Intn(len(lnames))]
		identitiy := rand.Int63n(10000000000000) + 10000000000
		password := hash.Generate(10)
		createdAt := date.GetNowString()
		s.users = append(s.users, user.User{
			ID:        int64(i),
			FirstName: name,
			LastName:  lname,
			Identitiy: fmt.Sprintf("%d", identitiy),
			Password:  password,
			Role:      &role.Role{Name: roles[rand.Intn(len(roles))]},
			CreatedAt: createdAt,
		})
		sqlStr += fmt.Sprintf("('%s','%s','%d','%s', '%s'),", name, lname, identitiy, password, createdAt)
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	err := db.QueryRow(sqlStr + ";")
	if err.Err() != nil {
		fmt.Println(err.Err())
	}

	return nil
}

func (s *seed) UserRoles() {
	db := postgres.DB.GetDB()
	db.QueryRow("TRUNCATE TABLE user_roles RESTART IDENTITY CASCADE;")
	sqlStr := "INSERT INTO user_roles (user_id, role_id, created_at) VALUES"
	roleID := 0

	for _, user := range s.users {
		if user.Role.Name == "student" {
			roleID = 1
		} else {
			roleID = 2
		}
		createdAt := date.GetNowString()

		sqlStr += fmt.Sprintf("('%d','%d', '%s'),", user.ID, roleID, createdAt)
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	err := db.QueryRow(sqlStr + ";")
	if err.Err() != nil {
		fmt.Println(err.Err())
	}
}

func (s *seed) UserCourses() {
	db := postgres.DB.GetDB()
	db.QueryRow("TRUNCATE TABLE user_course RESTART IDENTITY CASCADE;")
	sqlStr := "INSERT INTO user_course (user_id, course_id, created_at) VALUES"
	for _, user := range s.users {
		courseID := int64(rand.Intn(49) + 1)
		createdAt := date.GetNowString()

		sqlStr += fmt.Sprintf("('%d','%d','%s'),", user.ID, courseID, createdAt)
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	err := db.QueryRow(sqlStr + ";")
	if err.Err() != nil {
		fmt.Println(err.Err())
	}
}

func (s *seed) Courses() {
	rand.Seed(time.Now().UnixNano())
	db := postgres.DB.GetDB()
	db.QueryRow("TRUNCATE TABLE courses RESTART IDENTITY CASCADE;")
	sqlStr := "INSERT INTO courses (user_id, name, identitiy, valence, time, created_at) VALUES"
	counter := 50
	for i := 0; i < counter; i++ {
		user := s.users[rand.Intn(len(s.users))]
		if user.Role.Name == "teacher" {
			userID := user.ID
			name := courses[rand.Intn(len(courses))]
			valence := rand.Int63n(20) + 20
			time := times[rand.Intn(len(times))]
			identitiy := rand.Int63n(5000000000000) + 5000000000000
			createdAt := date.GetNowString()

			sqlStr += fmt.Sprintf("('%d','%s','%d','%d', '%s', '%s'),", userID, name, identitiy, valence, time, createdAt)
		} else {
			counter++
		}
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	err := db.QueryRow(sqlStr + ";")
	if err.Err() != nil {
		fmt.Println(err.Err())
	}
}
