package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tomteece/TrainingTracker/BackEnd/models"
)

func main() {
	models.InitDB()
	fmt.Println("listening on port 3000")
	http.HandleFunc("/Users", UsersIndex)
	http.HandleFunc("/User", handleUser)
	http.ListenAndServe(":3000", nil)
}

func UsersIndex(w http.ResponseWriter, r *http.Request) {
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
	for _, usr := range usrs {
		fmt.Fprintf(w, "%d, %s, %s, ", usr.ID, usr.FirstName, usr.LastName)
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
