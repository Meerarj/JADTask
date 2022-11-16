package main

import (
	"fmt"
	"strings"
)

func tolowercase(a []string) {

	for i, v := range a {

		var result string = strings.ToLower(v)
		a[i] = result
		fmt.Println(i, result)
		fmt.Printf(" i = %d v = %s \n", i, v)
	}

}

func main() {
	a := []string{"HI", "MEERA"}
	tolowercase(a)
	fmt.Println(a)
}
