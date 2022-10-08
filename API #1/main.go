package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Class    *Class `json:"class"`
}

type Class struct {
	Name   string `json:"name"`
	Weapon string `json:"weapon"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)

			return
		}
	}

	json.NewEncoder(w).Encode(User{})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := User{
		ID: strconv.Itoa(rand.Intn(1000000)),
	}

	json.NewDecoder(r.Body).Decode(&user)

	users = append(users, user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range users {
		if item.ID == params["id"] {
			var user User

			json.NewDecoder(r.Body).Decode(&user)

			users[index] = user
			users[index].ID = item.ID
		}
	}

	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)

			json.NewEncoder(w).Encode(users)

			return
		}
	}

	json.NewEncoder(w).Encode(User{})
}

func main() {
	r := mux.NewRouter()

	users = []User{
		{
			ID:       "1",
			Username: "User1",
			Password: "password123",
			Class: &Class{
				Name:   "Warrior",
				Weapon: "Sword",
			},
		},
		{
			ID:       "2",
			Username: "User2",
			Password: "pass123",
			Class: &Class{
				Name:   "Mage",
				Weapon: "Staff",
			},
		},
	}

	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
