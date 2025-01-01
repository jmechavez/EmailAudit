package domain

type User struct {
	EmailId int
	Fname   string
	Lname   string
	IdNo    int
	Email   string
	Status  string
}

type userRepo interface {
	AllUsers(status string) ([]User, *errs.AppError)
}
