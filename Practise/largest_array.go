// Golang Program to Find Largest Element in an Array
package main

import "fmt"

func main() {

// taking an array variable
	var n [100]float64
	var total int
	fmt.Print("Enter number of elements: ")
	
	// taking user input
	fmt.Scanln(&total)
	for i := 0; i < total; i++ {
		fmt.Print("Enter the number : ")
		fmt.Scan(&n[i])
	}
	for j := 1; j < total; j++ {
		if n[0] > n[j] {
			n[0] = n[j]
		}
	}

	fmt.Print("The largest number is : ", n[0])
}
