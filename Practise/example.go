package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1,2,3}

	fmt.Println(arr)

	slice_arr := []int{10,20}
	fmt.Println(slice_arr)

	slice_arr = append(slice_arr, 30,40)

	fmt.Println(slice_arr)

	slice_arr = append(slice_arr[1:3])

	fmt.Println(slice_arr)

	scores := make([]int, 3)

	scores[0] = 5

	fmt.Println(scores)

	scores = append(scores, 20,25)

	fmt.Println(scores)

	//scores[5] = 21

	fmt.Println(scores)

	// Deleting a value 
	var index int = 2

	scores = append(scores[:index])

	fmt.Println(scores)
}