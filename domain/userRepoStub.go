package domain

type userRepoStub struct {
	user []User
}

func (r userRepoStub) AllUser() ([]User, *errs.AppError) {
	return r.user, nil
}

func NewUserRepoStub() userRepoStub {
	users := []User{
		{
			240001,
			"John Michael",
			"Ez",
			2265,
			"jme@test.com",
			"0",
		},
		{
			240002,
			"Joan Ryms",
			"Ez",
			2266,
			"jre@test.com",
			"0",
		},
	}
	return NewUserRepoStub(users)
}
