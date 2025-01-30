package dto

import (
	"strings"

	"github.com/jmechavez/EmailAudit/errors"
)

type NewUserRequest struct {
	EmailId     int64  `json:"email_id"`
	Fname       string `json:"fname"`
	Lname       string `json:"lname"`
	IdNo        int    `json:"id_no"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	EmailAction string `json:"email_action"`
}

func (r NewUserRequest) Validate() *errors.AppError {
	if strings.ToLower(r.EmailAction) != "save" && strings.ToLower(r.EmailAction) != "create" {
		return errors.NewValidationError("Account type should be save & create")
	}
	return nil
}
