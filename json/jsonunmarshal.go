package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", Hi)
	r.HandleFunc("/Hi/{name}", welcome)
	r.HandleFunc("/new", New).Methods("POST")
	fmt.Println("server running at 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func Hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))

}

func welcome(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintln(w, "Hi", params["name"])

}

func New(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error ", err)
	}

	var prs Person
	json.Unmarshal(body, &prs)

	json.NewEncoder(w).Encode(prs)

}
