package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIHandler struct {
	odd  chan (int)
	even chan (int)
}

func (a *APIHandler) SayHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := r.URL.Query()
	w.Write([]byte(fmt.Sprintf("<h1>%s from %s</h1>", vars["name"], value["location"])))
}
func (a *APIHandler) Print(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%s", vars["intValue"])
	val := 0
	switch {
	case val%2 == 0:
		fmt.Println("Even")
		a.even <- val

	case val%2 != 0:
		fmt.Println("Odd")
		a.odd <- val
	default:
		fmt.Println(" ")
	}
}
func (a *APIHandler) ReadfromChannel() {
	for {

		select {
		case <-a.odd:
			fmt.Println("given number is odd")
		case <-a.even:
			fmt.Println("given number is even")
		}
	}
}
func main() {
	a := &APIHandler{}
	go a.ReadfromChannel()
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", a.SayHello)
	r.HandleFunc("/number/{intValue}", a.Print)
	fmt.Printf("Server started at :8080....")
	http.ListenAndServe(":8080", r)
}
