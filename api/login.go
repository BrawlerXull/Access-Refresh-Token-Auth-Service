package api

import (
	"anonymous-poll/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/x-www-form-urlencode")
		w.Header().Set("Allow-Control-Allow-Methods", "POST")
		users := GetAllUsers()
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)
		if validateTheUser(&user, users) {
			user.AccessToken = generateToken()
			user.RefreshToken = generateToken()
			json.NewEncoder(w).Encode(map[string]string{"username": user.UserName, "email": user.Email, "access_token": user.AccessToken, "refresh_token": user.RefreshToken, "password": user.Password})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"message": "Invalid username or password"})
		}
	}
}

func validateTheUser(user *models.User, users []models.User) bool {
	fmt.Println("loop", users)
	for _, currentUser := range users {
		if currentUser.Email == user.Email && currentUser.Password == user.Password {
			user.UserName = currentUser.UserName
			return true
		}
	}
	return false
}
