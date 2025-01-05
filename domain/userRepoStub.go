package domain

type UserRepoStub struct {
	user []User
}

func (r UserRepoStub) FindAll() ([]User, error) {
	return r.user, nil
}

func NewUserRepoStub() UserRepoStub {
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
	return UserRepoStub{users}
}
