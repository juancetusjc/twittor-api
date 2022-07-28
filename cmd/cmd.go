package cmd

import "net/http"

/* cmd init*/
func cmd() {
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8083", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./pong.html")
}
