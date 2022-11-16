package main

import (
	"encoding/json"
	"fmt"
)
func main() {
type Fruits struct {
	Name   string 
	Size   string 
	Color  string 
}


	s := &Fruits{Name: "Apple", Size: "Large", Color: "Red"}
	e, _ := json.Marshal(s)
	t := `{Name: "Apple", Size: "Large", Color: "Red" }`
	f := &Fruits{}
	json.Unmarshal([]byte(t), f)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println(string(e))
	fmt.Println(string(f)
}
