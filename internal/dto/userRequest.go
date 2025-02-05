package dto

import (
	"strings"

	"github.com/jmechavez/EmailAudit/errors"
)

type NewUserRequest struct {
	Fname       string `json:"fname"`
	Lname       string `json:"lname"`
	Email       string `json:"email"`
	EmailAction string `json:"email_action"`
}

func (r NewUserRequest) Validate() *errors.AppError {
	if strings.ToLower(r.EmailAction) != "save" && strings.ToLower(r.EmailAction) != "create" {
		return errors.NewValidationError("Account type should be save & create")
	}
	return nil
}
