package main

import (
	"fmt"
	"strings"
)

type wordfunction interface {
	ToLowerCase() string
	ToUpperCase() string
}
type mywords struct {
	str1 []string
}

func (w wordfunction) ToLowerCase() string {
	res1:=strings.ToLower(w.str1)
	fmt.Printf("Lowercase:%v",res1)
	
}
func (w wordfunction) ToUpperCase() string{
	res2:=strings.ToUpper((w.str1))
	fmt.Printf(("Uppercase:%v",res2)
	
}
func main (){
	a:= wordfunction{"MeeRa"}
	fmt.Println(a.ToLower())
	fmt.Println(a.ToUpper())
}


