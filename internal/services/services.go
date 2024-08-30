package services

import (
	"card-project/internal/database"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Unable read body: ", err)
	}

	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal("Unable unmarshal body: ", err)
	}

	fmt.Println("Add user: ", user)

	database.AddUser(user.FirstName, user.LastName)

	w.Write([]byte(body))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	fmt.Println("Delete user, id: ", userID)

	database.DeleteUser(userID)

	w.Write([]byte(userID))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Unable read body: ", err)
	}

	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal("Unable unmarshal body: ", err)
	}

	fmt.Println("Update user, id: ", userID)

	database.UpdateUser(userID, user.FirstName, user.LastName)

	w.Write([]byte(userID))
}
