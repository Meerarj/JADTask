package main

import (
	"fmt"
)

func main() {
	n := 4
	px := n
	py := n
	for i := 1; i <= n; i++ {
		for j := 0; j <= n*2; j++ {
			if (j == px) || (j == py) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		px = px - 1
		py = py + 1
		fmt.Println("\n")
	}
}
