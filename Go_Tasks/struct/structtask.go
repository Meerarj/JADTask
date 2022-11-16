package main

import "fmt"

type word interface {
	concatenate()
}
type words struct {
	str1 string
	str2 string
}

func (w words) concatenate() string {
	return w.str1 + w.str2
}
func main() {

	t := words{"Fill", "Up"}
	fmt.Println("Newword is:", t.concatenate())
}
