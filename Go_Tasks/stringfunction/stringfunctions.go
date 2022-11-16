package main

import (
	"fmt"
	"strings"
)

type wordgame struct {
	str1 string
	str2 string
}

func (s wordgame) LowerCase() string {
	temp1 := strings.ToLower(s.str1)
	temp2 := strings.ToLower(s.str2)
	fmt.Printf("Lowercase Strings:%v,%v", temp1, temp2)
	return " "
}
func (s wordgame) UpperCase() string {
	temp1 := strings.ToUpper(s.str1)
	temp2 := strings.ToUpper(s.str2)
	fmt.Printf("Uppercase Strings:%v,%v", temp1, temp2)
	return " "
}
func main() {
	f := wordgame{"Hi", "MeeRa"}
	fmt.Println(f.LowerCase())
	fmt.Println(f.UpperCase())
}
