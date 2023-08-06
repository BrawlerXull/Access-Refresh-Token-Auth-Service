package models

type User struct {
	UserName     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	AccessToken  string `json:"acccess_token"`
	RefreshToken string `json:"refresh_token"`
}
