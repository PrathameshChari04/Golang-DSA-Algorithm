/*

Given an array Arr of N positive integers. Your task is to find the elements whose value is equal to that of its index value ( Consider 1-based indexing ).

Note: There can be more than one element in the array which have the same value as its index. You need to include every such element's index. Follows 1-based indexing of the array.

Example 1:

Input:
N = 5
Arr[] = {15, 2, 45, 12, 7}
Output: 2
Explanation: Only Arr[2] = 2 exists here.



*/

package main

import "fmt"

func valueEqualToIndex(data []int) []int {

	result := []int{}

	for i, num := range data {
		if num == (i + 1) {
			result = append(result, i+1)
		}
	}

	return result

}

func main() {
	arr := []int{1, 20, 3}

	result := valueEqualToIndex(arr)

	if len(result) > 0 {
		fmt.Println("Values equal to index:", result) // Output: Values equal
	} else {
		fmt.Println("No Result Found")
	}
}