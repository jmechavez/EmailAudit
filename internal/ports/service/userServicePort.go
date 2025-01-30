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
	NewUser(dto.NewUserRequest) (*dto.NewUserResponse, *errors.AppError)
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

func (r DefaultUserService) NewUser(
	rq dto.NewUserRequest,
) (*dto.NewUserResponse, *errors.AppError) {
	err := rq.Validate()
	if err != nil {
		return nil, err
	}
	u := domain.User{
		EmailId:     0,
		Fname:       rq.Fname,
		Lname:       rq.Lname,
		IdNo:        rq.IdNo,
		Email:       rq.Email,
		Status:      "1",
		EmailAction: rq.EmailAction,
	}
	newUser, err := r.repo.AddUser(u)
	if err != nil {
		return nil, err
	}
	response := newUser.ToNewAccountResponseDTO()
	return &response, nil
}

func NewUserService(repository domain.UserRepo) DefaultUserService {
	return DefaultUserService{repository}
}
