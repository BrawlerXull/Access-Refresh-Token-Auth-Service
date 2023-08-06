package api

import (
	"anonymous-poll/database"
	"anonymous-poll/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

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
