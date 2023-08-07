package api

import (
	"anonymous-poll/database"
	"anonymous-poll/models"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllMyUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allUsers, err := GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Unable to fetch data from the database"})
		return
	}
	json.NewEncoder(w).Encode(allUsers)
}

func GetAllUsers() ([]models.User, error) {
	cur, err := database.Collection().Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var users []models.User
	for cur.Next(context.Background()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
