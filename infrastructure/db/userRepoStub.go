package db

import (
	"github.com/jmechavez/EmailAudit/internal/domain"
)

type UserRepoStub struct {
	user []domain.User
}

func (r UserRepoStub) FindAll() ([]domain.User, error) {
	return r.user, nil
}

func NewUserRepoStub() UserRepoStub {
	users := []domain.User{
		{
			EmailId: 240001,
			Fname:   "John Michael",
			Lname:   "Ez",
			IdNo:    2265,
			Email:   "jme@test.com",
			Status:  "0",
		},
		{
			EmailId: 240002,
			Fname:   "Joan Ryms",
			Lname:   "Ez",
			IdNo:    2266,
			Email:   "jre@test.com",
			Status:  "0",
		},
	}
	return UserRepoStub{users}
}
