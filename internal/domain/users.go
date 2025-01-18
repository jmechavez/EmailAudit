package domain

import "github.com/jmechavez/EmailAudit/errors"

type User struct {
	EmailId int    `json:"email_id"`
	Fname   string `json:"fname"`
	Lname   string `json:"lname"`
	IdNo    int    `json:"id_no"`
	Email   string `json:"email"`
	Status  string `json:"status"`
}

type UserRepo interface {
	FindAll(status string) ([]User, *errors.AppError)
	ByUserNum(id string) (*User, *errors.AppError)
}
