package middleware

import (
	"net/http"

	"github.com/juancetusjc/twittor-back/connectiondb"
)

/* Check connection DBs*/
func CheckConnection(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if connectiondb.CheckConnection() == 0 {
			http.Error(w, "Error connection to DBs", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
