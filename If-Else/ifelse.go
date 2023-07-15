package main

import "fmt"

func main() {

	fmt.Println("If - Else")

	if  7%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 4; num%2 == 0 {
		fmt.Printf(" %v is Even", num)
	} else {
		fmt.Println("Odd")
	}

	
}