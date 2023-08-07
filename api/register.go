package api

import (
	"anonymous-poll/database"
	"anonymous-poll/models"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/x-www-form-urlencode")
		w.Header().Set("Allow-Control-Allow-Methods", "POST")
		users, err := GetAllUsers()
		checkError(err)
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)
		if checkIfUserExists(user, users) {
			user.AccessToken = generateToken()
			user.RefreshToken = generateToken()
			expirationTime := time.Now().Add(1 * time.Hour)
			user.ExpiryTimeDate = expirationTime
			saveUser(user)
			value := []models.AccessRefreshTokenPair{}
			json.NewEncoder(w).Encode(map[string]interface{}{
				"username":                  user.UserName,
				"email":                     user.Email,
				"access_token":              user.AccessToken,
				"refresh_token":             user.RefreshToken,
				"password":                  user.Password,
				"access_refresh_token_pair": value,
				"expiry_time_date":          expirationTime,
			})
		} else {
			w.WriteHeader(http.StatusAlreadyReported)
			json.NewEncoder(w).Encode(map[string]string{"message": "user already exists"})
		}
	}
}

func saveUser(user models.User) {
	inserted, err := database.Collection().InsertOne(context.Background(), user)
	checkError(err)
	fmt.Println("Inserted 1 user in db with id: ", inserted.InsertedID)
}

func checkIfUserExists(user models.User, users []models.User) bool {
	fmt.Println("loop", users)
	for _, currentUser := range users {
		if currentUser.Email == user.Email || currentUser.UserName == user.UserName {
			return false
		}
	}
	return true
}

func generateToken() string {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	checkError(err)
	token := base64.StdEncoding.EncodeToString(tokenBytes)
	return token
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
