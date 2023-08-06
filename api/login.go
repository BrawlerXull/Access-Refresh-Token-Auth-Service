package api

import (
	"anonymous-poll/database"
	"anonymous-poll/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/x-www-form-urlencode")
		w.Header().Set("Allow-Control-Allow-Methods", "POST")
		users := GetAllUsers()
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)
		if validateTheUser(&user, users) {
			expirationTime := time.Now().Add(1 * time.Hour)
			user.ExpiryTimeDate = expirationTime
			userObtained, err := GetOneUserByEmail(user.Email)
			checkError(err)
			user.AccessToken = generateToken()
			user.RefreshToken = generateToken()
			value := models.AccessRefreshTokenPair{
				AccessToken:  user.AccessToken,
				RefreshToken: user.RefreshToken,
			}
			userObtained.AccessRefreshTokenPairList = append(userObtained.AccessRefreshTokenPairList, value)
			updateUser(userObtained)
			fmt.Println(userObtained.AccessRefreshTokenPairList)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"username":                  user.UserName,
				"email":                     user.Email,
				"access_token":              user.AccessToken,
				"refresh_token":             user.RefreshToken,
				"password":                  user.Password,
				"access_refresh_token_pair": userObtained.AccessRefreshTokenPairList,
				"expiry_time_date":          expirationTime,
			})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"message": "Invalid username or password"})
		}
	}
}

func validateTheUser(user *models.User, users []models.User) bool {
	for _, currentUser := range users {
		if currentUser.Email == user.Email && currentUser.Password == user.Password {
			user.UserName = currentUser.UserName
			return true
		}
	}
	return false
}

func updateUser(newUserValue models.User) {
	existingUser, err := GetOneUserByEmail(newUserValue.Email)
	fmt.Println(existingUser)
	if err != nil {
		log.Fatal(err)
		return
	}
	existingUser = newUserValue
	err = updateUserInDatabase(existingUser)
	if err != nil {
		log.Fatal(err)
	}
}

func updateUserInDatabase(user models.User) error {
	_, err := database.Collection().UpdateOne(context.Background(), bson.M{"email": user.Email}, bson.M{"$set": user})
	return err
}

// func isExpired(user models.User) bool {
// 	if user.ExpiryTimeDate.Before(time.Now()) {
// 		return true
// 	} else {
// 		return false
// 	}
// }
