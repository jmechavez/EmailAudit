package service

import "github.com/jmechavez/EmailAudit/domain"

type UserService interface {
	GetAllUser() ([]domain.User, error)
}

type DefaultUserService struct {
	repo domain.UserRepo
}

func (r DefaultUserService) GetAllUser() ([]domain.User, error) {
	return r.repo.FindAll()
}

func NewUserService(repository domain.UserRepo) DefaultUserService {
	return DefaultUserService{repository}
}
