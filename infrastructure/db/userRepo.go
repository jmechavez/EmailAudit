package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/jmechavez/EmailAudit/errors"
	"github.com/jmechavez/EmailAudit/infrastructure/logger"
	"github.com/jmechavez/EmailAudit/internal/domain"
)

type UserRepoDb struct {
	userDb *sqlx.DB
}

func (d UserRepoDb) AddUser(u domain.User) (*domain.User, *errors.AppError) {
	insertUserSql := "INSERT INTO users (fname, lname, id_no, email, status ,email_action) VALUES ($1, $2, $3, $4, $5, $6) RETURNING email_id"

	// Execute the query and scan the returned ID
	var emailId int64
	err := d.userDb.QueryRow(
		insertUserSql,
		u.Fname,
		u.Lname,
		u.IdNo,
		u.Email,
		u.Status,
		u.EmailAction,
	).Scan(&emailId)
	if err != nil {
		logger.Error("Error while creating account: " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected Database Error")
	}

	u.EmailId = emailId

	return &u, nil
}

func (d UserRepoDb) FindAll(status string) ([]domain.User, *errors.AppError) {
	// var rows *sql.Rows
	var err error
	var users []domain.User

	if status == "" {
		findUserSql := "SELECT email_id, fname, lname, id_no, email, status, email_action FROM users"
		err = d.userDb.Select(&users, findUserSql)
		// rows, err = d.userDb.Query(findNameSql)
	} else {
		findUserSql := "SELECT email_id, fname, lname, id_no, email, status, email_action FROM users WHERE status = $1"
		err = d.userDb.Select(&users, findUserSql, status)
		// rows, err = d.userDb.Query(findNameSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected Database Error")
	}

	// err = sqlx.StructScan(rows, &users)
	// if err != nil {
	// 	logger.Error("Error while scanning customers " + err.Error())
	// 	return nil, errors.NewUnExpectedError("Error Parsing Client Data")
	// }
	// replace into sqlx
	// var users []domain.User
	// for rows.Next() {
	// 	var c domain.User
	// 	err := rows.Scan(
	// 		&c.EmailId,
	// 		&c.Fname,
	// 		&c.Lname,
	// 		&c.IdNo,
	// 		&c.Email,
	// 		&c.Status,
	// 	)
	// 	if err != nil {
	// 		logger.Error("Error while scanning customers " + err.Error())
	// 		return nil, errors.NewUnExpectedError("Error Parsing Client Data")
	// 	}
	// 	users = append(users, c)
	// }

	return users, nil
}

func (d UserRepoDb) ByUserNum(id string) (*domain.User, *errors.AppError) {
	findUserSql := "SELECT email_id, fname, lname, id_no, email, status, email_action FROM users WHERE email_id = $1"

	var users domain.User

	err := d.userDb.Get(&users, findUserSql, id)
	// rows, err := d.userDb.Query(findUserSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("User not found")
		}
		logger.Error("Error while querying user table: " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	// defer rows.Close()

	// Check if any rows are returned
	// if rows.Next() {
	// 	err = rows.Scan(
	// 		&user.EmailId,
	// 		&user.Fname,
	// 		&user.Lname,
	// 		&user.IdNo,
	// 		&user.Email,
	// 		&user.Status,
	// 	)
	// 	if err != nil {
	// 		log.Println("Error while scanning user: " + err.Error())
	// 		return nil, errors.NewUnExpectedError("Error Parsing User Data")
	// 	}
	// } else {
	// 	// Handle no rows found
	// 	return nil, errors.NewNotFoundError("User not found")
	// }
	//
	return &users, nil
}

func NewUserRepoDb(userPostDb *sqlx.DB) UserRepoDb {
	return UserRepoDb{userPostDb}
}
