package domain

import (
	"github.com/jmechavez/EmailAudit/errors"
	"github.com/jmechavez/EmailAudit/internal/dto"
)

type User struct {
	EmailId int    `json:"email_id" db:"email_id"`
	Fname   string `json:"fname"    db:"fname"`
	Lname   string `json:"lname"    db:"lname"`
	IdNo    int    `json:"id_no"    db:"id_no"`
	Email   string `json:"email"    db:"email"`
	Status  string `json:"status"   db:"status"`
}

func (u User) statusAsText() string {
	statusAsText := "active"
	if u.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (u User) ToDto() dto.UserResponse {
	return dto.UserResponse{
		EmailId: u.EmailId,
		Fname:   u.Fname,
		Lname:   u.Lname,
		IdNo:    u.IdNo,
		Email:   u.Email,
		Status:  u.statusAsText(), // call func statusAsText
	}
}

type UserRepo interface {
	FindAll(status string) ([]User, *errors.AppError)
	ByUserNum(id string) (*User, *errors.AppError)
}
