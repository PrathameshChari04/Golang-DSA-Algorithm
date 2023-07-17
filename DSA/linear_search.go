package main

import (
	"fmt"
)

func linearSearch(data []int, key int) bool {
	for item,_ := range data {
		if item == key {
			return true
		}
	}
	return false
}


func main() {
	items := []int{1,2}
	fmt.Println(linearSearch(items,10))
}