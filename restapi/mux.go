package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Name struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
	Content string `json:"Content"`
}
type Names []Name

func allNames(w http.ResponseWriter, r *http.Request) {
	a := Names{
		Name{Name: "Meera", Address: "tky", Content: "Hello World"},
	}
	fmt.Println("Endpoint hit:All articles endpoint")
	fmt.Fprintf(w, "%s", a)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}
func handleRequest() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/names", allNames).Methods("GET")
	return r
}
func main() {
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":8080"),
		Handler: handleRequest(),
	}
	err := httpServer.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("unexpected :%v", err)
		os.Exit(1)
	}
}
