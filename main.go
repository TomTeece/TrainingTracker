package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tomteece/TrainingTracker/models"
)

func main() {
	models.InitDB()
	fmt.Println("listening on port 3001")
	http.HandleFunc("/Users", usersIndex)
	http.HandleFunc("/User", handleUser)
	http.HandleFunc("/Climbing", handleClimbing)
	http.ListenAndServe(":3001", nil)
}

func handleClimbing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var climbing models.Climbing
		_ = json.NewDecoder(r.Body).Decode(&climbing)

		id, err := models.AddClimbing(climbing)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(400), 400)
			return
		}

		json.NewEncoder(w).Encode(id)
		return
	}
	if r.Method == "GET" {
		var climbing models.Climbing
		_ = json.NewDecoder(r.Body).Decode(&climbing)
		fmt.Println(climbing.UserID)
		usrClimbing, err := models.ClimbingByUserID(climbing)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		json.NewEncoder(w).Encode(usrClimbing)
		return
	} else {
		http.Error(w, http.StatusText(405), 405)
		return
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	id, err := models.AddUser(user.FirstName, user.LastName)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	json.NewEncoder(w).Encode(id)

}

func usersIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	usrs, err := models.AllUsers()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(usrs)
}
