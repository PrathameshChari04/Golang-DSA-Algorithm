package main

import (
	"fmt"
)

// Fibonacci Function
func fibonacci(n int)int{
	if(n < 2){
		return n
	}
	return fibonacci(n-1) + fibonacci(n - 2)
}

func main() {
	fmt.Println("Fibonacci Problem !")
	result := fibonacci(7)
	fmt.Println(result)
}