package routers

import (
	"encoding/json"
	"net/http"

	"github.com/juancetusjc/twittor-back/model"
)

/* Register to DBs*/
func Register(w http.ResponseWriter, r *http.Request) {
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Error read body "+err.Error(), 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "Invalid email", 400)
	}
	if len(u.Password) < 6 {
		http.Error(w, "Invalid password, min char(6)", 400)
	}
}
