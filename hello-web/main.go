package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>Welcome to go web development course</h1>"))
}

func main() {
	http.HandleFunc("/", home)

	log.Println("listening web server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
