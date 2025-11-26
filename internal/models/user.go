package models


type User struct {
	ID             int `db:"id"`
	Username       string `db:"username"`
	Password       string `db:"password"`
	Email          string `db:"email"`
	IsAdmin        bool   `db:"is_admin"`
	ProfilePicture string `db:"profile_picture"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type LoginResp struct {
	Token string `json:"token"`
}


type ForgotPasswordReq struct {
	Username string `json:"username"`
}