package service

import (
	"github.com/jmechavez/EmailAudit/errors"
	"github.com/jmechavez/EmailAudit/internal/domain"
)

type UserService interface {
	GetAllUser() ([]domain.User, *errors.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepo
}

func (r DefaultUserService) GetAllUser() ([]domain.User, *errors.AppError) {
	return r.repo.FindAll()
}

func NewUserService(repository domain.UserRepo) DefaultUserService {
	return DefaultUserService{repository}
}
