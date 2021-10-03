package routers

import (
	"encoding/json"
	"net/http"

	"github.com/4rufi/curso-go/db"
	"github.com/4rufi/curso-go/models"
)

/* REgistry - function for create a new user in db*/
func Registry(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w,"Error in data" + err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email can't be empty", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password can't be less of 6 caracters", http.StatusBadRequest)
		return
	}

	_, find, _ := db.CheckUserExist(t.Email)
	if find {
		http.Error(w, "User register with that email", http.StatusBadRequest)
		return
	}

	_, status, _ :=  db.InsertRegistry(t)
	if err != nil {
		http.Error(w, "Error when insert user registry"+ err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Error user can't registry", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}