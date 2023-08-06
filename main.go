package main

import (
	"anonymous-poll/database"
	"anonymous-poll/router"
	"fmt"
	"net/http"
)

func main() {
	r := router.Router()
	database.Init()
	fmt.Println("Server is getting started...")
	fmt.Println("Server listening on :3001")
	http.ListenAndServe(":3001", r)
}
