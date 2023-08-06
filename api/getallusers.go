package api

import (
	"anonymous-poll/database"
	"anonymous-poll/models"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllMyUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allUsers := GetAllUsers()
	json.NewEncoder(w).Encode(allUsers)
}

func GetAllUsers() []models.User {
	cur, err := database.Collection().Find(context.Background(), bson.D{{}})
	checkError(err)
	var users []models.User

	for cur.Next(context.Background()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer cur.Close(context.Background())
	return users
}
