package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func check(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(200)
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "<h1>Health check</h1>")

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "200 OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	// http.HandleFunc("/", check)
	// http.HandleFunc("/healthy", check)
	// fmt.Println("Server starting...")
	// http.ListenAndServe(":8080", nil)

	e.Logger.Fatal(e.Start(":" + httpPort))
}
