package main

import "fmt"

func main() {
	var name string = "Go"
	fmt.Println(name)
	fmt.Printf("Variable is of Type: %T \n", name)

	var number int = 100
	fmt.Println(number)
	fmt.Printf("Variable is of Type: %T \n ", number)

	var isActive bool = true 
	fmt.Println(isActive)
	fmt.Printf("Variable is of Type: %T \n ", isActive)

}