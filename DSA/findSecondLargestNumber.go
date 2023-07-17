package main

import ("fmt")


func secondLargest(data []int) int{
	max := -1
	secondMax := -1

	for _ , num := range data {
		if num > max {
			max = secondMax
			max = num
		} else if num < max && num > secondMax {
			secondMax = num
		}

	}
	return secondMax
}

func main() {
	arr := []int{12, 35, 1, 10, 34, 1}
	fmt.Println(arr)
	result := secondLargest(arr)

	if result != -1 {
		fmt.Println("Second Largest element is:", result )
	} else {
		fmt.Println("There are no Second largest elements in the array.")
	}
}