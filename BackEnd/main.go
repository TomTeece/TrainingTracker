package main

import (
	"fmt"
	"net/http"

	"github.com/tomteece/TrainingTracker/BackEnd/models"
)

func main() {
	models.InitDB()
	http.HandleFunc("/Users", UsersIndex)
	http.HandleFunc("/User", handleUser)
	http.ListenAndServe(":3000", nil)
}

func UsersIndex(w http.ResponseWriter, r *http.Request) {

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
		fmt.Fprintf(w, "%d, %s, %s, /n", usr.ID, usr.FirstName, usr.LastName)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {

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
		fmt.Fprintf(w, "%d, %s, %s, /n", usr.ID, usr.FirstName, usr.LastName)
	}
}