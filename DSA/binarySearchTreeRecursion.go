package main

import ("fmt")

func binarySearch(arr []int, s int, e int, target int) int {
	if s > e {
		return -1
	}
	m := s + (e - s) / 2

	if(arr[m] == target) {
		return m
	}
	// Left Hand Side
	if (target < arr[m]){
		binarySearch(arr,s,m -1,target)
	}
	return binarySearch(arr,m+1,e,target)
}

func main(){
	fmt.Println("Binary Search Tree !")

	arr := []int{1,2,3,4}
	target := 8
	fmt.Println(binarySearch(arr,0,len(arr) - 1,target))
}