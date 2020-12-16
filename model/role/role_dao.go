package role

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/emadghaffari/api-teacher/database/postgres"
	"github.com/emadghaffari/api-teacher/utils/date"
	"github.com/emadghaffari/res_errors/errors"
	log "github.com/sirupsen/logrus"
)

var (
	insertQuery         = "INSERT INTO roles(name) VALUES($1);"
	assignUserRoleQuery = "INSERT INTO user_roles(user_id, role_id) VALUES($1, $2);"
	selectQuery         = "SELECT id, name FROM roles WHERE name = $1;"
)

// Insert New Role
func (r *Role) Insert() errors.ResError {
	db := postgres.DB.GetDB()

	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("users").Columns("name").
		Values(r.Name, date.GetNowString()).
		Suffix("RETURNING id").ToSql()

	if err != nil {
		log.Error(fmt.Sprintf("Error in Insert Role StatementBuilder: %s", err))
		return errors.HandlerInternalServerError(err.Error(), err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Error(fmt.Sprintf("Error in Insert Role BeginDB: %s", err))
		return errors.HandlerInternalServerError(err.Error(), err)
	}

	if err := tx.QueryRow(query, args...).Scan(&r.ID); err != nil {
		log.Error(fmt.Sprintf("Error in Insert Role QueryRow: %s", err))
		return errors.HandlerInternalServerError(err.Error(), err)
	}
	tx.Commit()

	return nil
}

// GetByName : Get a Role By Name
func (r *Role) GetByName() errors.ResError {
	db := postgres.DB.GetDB()
	stms, err := db.Prepare(selectQuery)
	if err != nil {
		log.Error(fmt.Sprintf("Error in Select Role: %s", err))

		return errors.HandlerInternalServerError(err.Error(), err)
	}
	defer stms.Close()
	result := stms.QueryRow(r.Name)

	if err := result.Scan(
		&r.ID,
		&r.Name,
	); err != nil {
		log.Error(fmt.Sprintf("Error in Select Role: %s", err))

		return errors.HandlerInternalServerError(err.Error(), err)
	}

	return nil
}

// Assign Role to User
func (r *Role) Assign(id int64) errors.ResError {
	if r.ID > 0 && id > 0 {
		db := postgres.DB.GetDB()

		query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
			Insert("user_roles").Columns("user_id", " role_id", "created_at").
			Values(id, r.ID, date.GetNowString()).
			ToSql()

		if err != nil {
			log.Error(fmt.Sprintf("Error in user_roles StatementBuilder: %s", err))
			return errors.HandlerInternalServerError(err.Error(), err)
		}
		tx, err := db.Begin()
		if err != nil {
			log.Error(fmt.Sprintf("Error in user_roles BeginDB: %s", err))
			return errors.HandlerInternalServerError(err.Error(), err)
		}

		if err := tx.QueryRow(query, args...).Err(); err != nil {
			log.WithFields(log.Fields{
				"user_id": id,
				"role_id": r.ID,
			}).Error(fmt.Sprintf("Faild Assign Role To User: %s", err))
			return errors.HandlerInternalServerError(err.Error(), err)
		}
		tx.Commit()

		return nil
	}
	return errors.HandlerBadRequest("roleID or userID not valid!")
}
