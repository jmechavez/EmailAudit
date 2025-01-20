package domain

import "github.com/jmechavez/EmailAudit/errors"

type User struct {
	EmailId int    `json:"email_id" db:"email_id"`
	Fname   string `json:"fname"    db:"fname"`
	Lname   string `json:"lname"    db:"lname"`
	IdNo    int    `json:"id_no"    db:"id_no"`
	Email   string `json:"email"    db:"email"`
	Status  string `json:"status"   db:"status"`
}

type UserRepo interface {
	FindAll(status string) ([]User, *errors.AppError)
	ByUserNum(id string) (*User, *errors.AppError)
}
