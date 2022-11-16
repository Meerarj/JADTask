package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Username string
	Password string
	IsAdmin  bool
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/echo", echoHandler)

	http.ListenAndServe(":5000", mux)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{} //initialize empty user

	//Parse json request body and use it to set fields on user
	//Note that user is passed as a pointer variable so that it's fields can be modified
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	//Marshal or convert user object back to json and write to response
	userJson, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(userJson)
}
