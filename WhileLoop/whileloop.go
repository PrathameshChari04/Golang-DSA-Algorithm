package main

import (
	"fmt"
)

func main() {
	fmt.Printf("While Looping \n")

	i := 1

	for i <= 10 {
		println(i)
		i++
	}

	for i:= 1; i < 5; i ++ {
		for j:= 1; j < i; j ++{
			fmt.Printf("*")
		}
		fmt.Print("\n")
	}
}