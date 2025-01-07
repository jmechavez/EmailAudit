package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmechavez/EmailAudit/internal/domain"
)

type UserRepoDb struct {
	userDb *sql.DB
}

func (d UserRepoDb) FindAll() ([]domain.User, error) {
	findNameSql := "SELECT email_id, fname, lname, id_no, email, status FROM users"

	rows, err := d.userDb.Query(findNameSql)
	if err != nil {
		log.Println("Error while querying customer table" + err.Error())
		return nil, err
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
			return nil, err
		}
		users = append(users, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
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
