// primary port
package service

import (
	"github.com/jmechavez/EmailAudit/errors"
	"github.com/jmechavez/EmailAudit/internal/domain"
	"github.com/jmechavez/EmailAudit/internal/dto"
)

type UserService interface {
	GetAllUser(status string) ([]dto.UserResponse, *errors.AppError)
	ByUserNum(id string) (*dto.UserResponse, *errors.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepo
}

func (r DefaultUserService) GetAllUser(status string) ([]dto.UserResponse, *errors.AppError) {
	switch status {
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = ""
	}

	u, err := r.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	var uResponse []dto.UserResponse
	for _, user := range u {
		uResponse = append(uResponse, user.ToDto())
	}

	return uResponse, nil
}

func (r DefaultUserService) ByUserNum(id string) (*dto.UserResponse, *errors.AppError) {
	u, err := r.repo.ByUserNum(id)
	if err != nil {
		return nil, err
	}

	uResponse := u.ToDto()

	return &uResponse, nil
}

func NewUserService(repository domain.UserRepo) DefaultUserService {
	return DefaultUserService{repository}
}
