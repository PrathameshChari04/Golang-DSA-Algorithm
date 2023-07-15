package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch-Case in Golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	println("Value of Dice : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Move 1 Spot")
	case 2:
		fmt.Println("Move 2 Spot")
	case 3:
		fmt.Println("Move 3 Spot")
	case 4:
		fmt.Println("Move 4 Spot")
	case 5:
		fmt.Println("Move 5 Spot")
	case 6:
		fmt.Println("Move 6 Spot")
	default:
		fmt.Println("End Game")
	}

}