package auth

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{ID: 1, Username: "admin", Password: "password"},
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.Username == creds.Username && user.Password == creds.Password {
			w.Write([]byte("Login successful"))
			return
		}
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
