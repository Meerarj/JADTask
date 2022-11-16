package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:30000")
	if err != nil {
		log.Fatalf("HTTP GET request failed, %v\n", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Body read failed: %v\n", err)
	}
}
