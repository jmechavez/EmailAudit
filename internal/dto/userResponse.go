package dto

type UserResponse struct {
	EmailId     int64  `json:"email_id"`
	Fname       string `json:"fname"`
	Lname       string `json:"lname"`
	IdNo        int    `json:"id_no"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	EmailAction string `json:"email_action"`
}

type NewUserResponse struct {
	EmailId int64 `json:"email_id"`
}
