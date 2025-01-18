package service

import (
	"github.com/jmechavez/EmailAudit/errors"
	"github.com/jmechavez/EmailAudit/internal/domain"
)

type UserService interface {
	GetAllUser(status string) ([]domain.User, *errors.AppError)
	ByUserNum(id string) (*domain.User, *errors.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepo
}

func (r DefaultUserService) GetAllUser(status string) ([]domain.User, *errors.AppError) {
	switch status {
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = ""
	}
	return r.repo.FindAll(status)
}

func (r DefaultUserService) ByUserNum(id string) (*domain.User, *errors.AppError) {
	return r.repo.ByUserNum(id)
}

func NewUserService(repository domain.UserRepo) DefaultUserService {
	return DefaultUserService{repository}
}
