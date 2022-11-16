package main

import (
	"fmt"
	"log"
	"os"
)

func fileRead() string {
	content, err := os.ReadFile("taskfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	return (string(content))
}
func main() {
	temp := fileRead()
	log.Println(temp)
}
