package main

import (
	"fmt"
	"log"
	"net/http"
)

func restService(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func handleRequests() {
	http.HandleFunc("/", restService)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
func main() {
	handleRequests()
}
