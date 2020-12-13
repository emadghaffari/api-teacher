package user

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/emadghaffari/api-teacher/database/postgres"
	"github.com/emadghaffari/api-teacher/model/role"
	"github.com/emadghaffari/api-teacher/utils/date"
	"github.com/emadghaffari/res_errors/errors"
	log "github.com/sirupsen/logrus"
)

var (
	loginQuery = "SELECT cp.id,cp.name,cp.lname,cp.identitiy,m.name as role FROM roles as m INNER JOIN user_roles as p ON m.id = p.role_id INNER JOIN users as cp ON p.user_id = cp.id WHERE cp.identitiy = $1 AND cp.password = $2 ;"
)

// Register meth create new User
func (c *User) Register() errors.ResError {
	db := postgres.DB.GetDB()

	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("users").Columns("name", " lname", " identitiy", " password", " created_at").
		Values(c.FirstName, c.LastName, c.Identitiy, c.Password, date.GetNowString()).
		Suffix("RETURNING id").ToSql()

	if err != nil {
		log.Error(fmt.Sprintf("Error in StatementBuilder: %s", err))
		return errors.HandlerInternalServerError(err.Error(), err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Error(fmt.Sprintf("Error in BeginDB: %s", err))
		return errors.HandlerInternalServerError(err.Error(), err)
	}
	defer tx.Commit()

	if err := tx.QueryRow(query, args...).Scan(&c.ID); err != nil {
		log.Error(fmt.Sprintf("Error in QueryRow: %s", err))
		tx.Rollback()
		return errors.HandlerInternalServerError(err.Error(), err)
	}

	return nil
}

// Login meth for Login Users
func (c *User) Login() errors.ResError {
	db := postgres.DB.GetDB()
	stms, err := db.Prepare(loginQuery)
	if err != nil {
		log.Error(fmt.Sprintf("Error in Select Role: %s", err))

		return errors.HandlerInternalServerError(err.Error(), err)
	}
	defer stms.Close()
	result := stms.QueryRow(c.Identitiy, c.Password)

	var rl string
	if err := result.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Identitiy, &rl); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.HandlerBadRequest("user Not Found!")
		}
		log.Error(fmt.Sprintf("Error in Scan: %s", err))
		return errors.HandlerInternalServerError(err.Error(), err)
	}
	c.Role = &role.Role{
		Name: rl,
	}

	return nil
}
