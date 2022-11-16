package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func EmployeeDetails(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}

	var emp Employee
	json.Unmarshal(bodyBytes, &emp)
	json.NewEncoder(w).Encode(emp)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/employee", EmployeeDetails)
	r.HandleFunc("/newemployee", EmployeeDetails).Methods("POST")
	fmt.Printf("Server started at:8080")
	http.ListenAndServe(":8082", r)
}
