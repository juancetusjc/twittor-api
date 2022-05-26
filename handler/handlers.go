package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8083", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./pong.html")
}

/*Handler config port and listener */
func Handerls() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	handlers := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handlers))

}
