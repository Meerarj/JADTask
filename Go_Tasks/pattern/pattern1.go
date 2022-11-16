package main

import "fmt"

func main() {
	var rows int = 4
	for i := 1; i <= rows; i++ {
		k := 0
		for space := 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if k == 1*i {
				break
			}
		}
		fmt.Println("")
	}

}
