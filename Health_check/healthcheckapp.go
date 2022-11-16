package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}
func check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>Health check</h1>")
	// fmt.Fprint(w, "<h1>200 OK<h1>")
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health_check", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":30000", nil)

}
