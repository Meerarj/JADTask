package main

import "fmt"

var My_map = make(map[int]string)

func main() {
	// Creating a map
	// Using make() function
	fmt.Println(My_map)
	// As we already know that make() function
	// always returns a map which is initialized
	// So, we can add values in it
	fmt.Println(My_map)
	fmt.Println("1:Insertion\n 2:Deletion,\n 3:Modification")
	fmt.Println("Insert any no betweeen 1 to 3")
	var choice int
	fmt.Scanln(&choice)
	//fmt.Println(choice)
	switch choice {
	case 1:
		insertMap()
	case 2:
		deleteMap()
	case 3:
		modifyMap()
	}
}
func insertMap() {
	My_map[11] = "Dog"
	My_map[22] = "Cat"
	My_map[33] = "Lion"
	My_map[44] = "Tiger"
	fmt.Printf("After insertion of keys and values in map:%v", My_map)
}
func deleteMap() {
	insertMap()
	delete(My_map, 33)
	fmt.Printf("After Deleting the key 33:%v", My_map)
}
func modifyMap() {
	insertMap()
	My_map[3] = "Pig"
	My_map[4] = "Goat"
	fmt.Println("After Modification:%v", My_map)
}
