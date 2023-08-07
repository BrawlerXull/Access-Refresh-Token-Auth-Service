package api

import (
	"anonymous-poll/models"
	"encoding/json"
	"net/http"
)

func ReAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Methods", "POST")

		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
			return
		}

		dbUser, err := GetOneUserByEmail(user.Email)
		checkError(err)
		if dbUser.AccessToken == user.AccessToken {
			dbUser.AccessToken = generateToken()
			dbUser.RefreshToken = generateToken()
			value := models.AccessRefreshTokenPair{
				AccessToken:  dbUser.AccessToken,
				RefreshToken: dbUser.RefreshToken,
			}
			dbUser.AccessRefreshTokenPairList = append(dbUser.AccessRefreshTokenPairList, value)
			updateUser(dbUser)

			json.NewEncoder(w).Encode(dbUser)
		} else {
			if dbUser.RefreshToken == user.RefreshToken {
				dbUser.AccessToken = generateToken()
				dbUser.RefreshToken = generateToken()
				value := models.AccessRefreshTokenPair{
					AccessToken:  dbUser.AccessToken,
					RefreshToken: dbUser.RefreshToken,
				}
				dbUser.AccessRefreshTokenPairList = append(dbUser.AccessRefreshTokenPairList, value)
				updateUser(dbUser)

				json.NewEncoder(w).Encode(dbUser)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"message": "Invalid User Access"})
			}
		}
	}
}
