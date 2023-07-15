package main

import ("fmt")

func linearSearch(data []int, key int )int {
	fmt.Println("Linear Search")

	for i:=0; i<len(data); i++ {
		if data[i] == key {
			return i
		}
	}
	return -1
}

func main() {
	// Sample Data
	var arr = [5]int{23, 46, 7}
	key := 7
	resultIndex := linearSearch(arr[:], key)
	if resultIndex!= -1 {	
		fmt.Printf("%d is present at index %d\n", key, resultIndex)
	} else {
		fmt.Printf("%d not found in array.\n", key)
	}
}

// Output: Linear search for the element with value '7'...