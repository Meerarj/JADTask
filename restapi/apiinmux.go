package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func main() {
	a := Handler{}

	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", a.SayHello)
	// r.HandleFunc("/number/{intValue}", a.Print)
	fmt.Printf("Server started at :8080....")
	http.ListenAndServe(":8080", r)

}

func (a *Handler) SayHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := r.URL.Query()
	w.Write([]byte(fmt.Sprintf("<h1>%s from %s</h1>", vars["name"], value["location"])))
}
