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
	NewUser(rq dto.NewUserRequest, idNo int) (*dto.NewUserResponse, *errors.AppError)
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
	idNo int,
) (*dto.NewUserResponse, *errors.AppError) {
	// Validate request
	if err := rq.Validate(); err != nil {
		return nil, err
	}

	// Create user domain object
	u := domain.User{
		EmailId:     0,
		Fname:       rq.Fname,
		Lname:       rq.Lname,
		IdNo:        idNo, // Use the provided idNo
		Email:       rq.Email,
		Status:      "active",
		EmailAction: rq.EmailAction,
	}

	// Add user to repository
	newUser, err := r.repo.AddUser(u)
	if err != nil {
		return nil, err
	}

	// Convert to response DTO
	response := newUser.ToNewAccountResponseDTO()
	return &response, nil
}

func NewUserService(repository domain.UserRepo) DefaultUserService {
	return DefaultUserService{repository}
}
