package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmechavez/EmailAudit/errors"
	"github.com/jmechavez/EmailAudit/internal/domain"
)

type UserRepoDb struct {
	userDb *sql.DB
}

func (d UserRepoDb) FindAll(status string) ([]domain.User, *errors.AppError) {
	var rows *sql.Rows
	var err error

	if status == "" {
		findNameSql := "SELECT email_id, fname, lname, id_no, email, status FROM users"
		rows, err = d.userDb.Query(findNameSql)
	} else {
		findNameSql := "SELECT email_id, fname, lname, id_no, email, status FROM users WHERE status = $1"
		rows, err = d.userDb.Query(findNameSql, status)
	}

	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected Database Error")
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var c domain.User
		err := rows.Scan(
			&c.EmailId,
			&c.Fname,
			&c.Lname,
			&c.IdNo,
			&c.Email,
			&c.Status,
		)
		if err != nil {
			log.Println("Error while scanning customers " + err.Error())
			return nil, errors.NewUnExpectedError("Error Parsing Client Data")
		}
		users = append(users, c)
	}

	return users, nil
}

func NewUserRepoDb() UserRepoDb {
	connStr := "user=admin password=admin123 dbname=email_dir sslmode=disable"
	userDb, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return UserRepoDb{userDb}
}
