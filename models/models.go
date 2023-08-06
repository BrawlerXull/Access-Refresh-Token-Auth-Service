package models

import "time"

type User struct {
	UserName                   string                   `json:"username"`
	Email                      string                   `json:"email"`
	Password                   string                   `json:"password"`
	AccessToken                string                   `json:"access_token"`
	RefreshToken               string                   `json:"refresh_token"`
	AccessRefreshTokenPairList []AccessRefreshTokenPair `json:"access_refresh_token_pair"`
	ExpiryTimeDate             time.Time                `json:"expiry_time_date"`
}

type AccessRefreshTokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
