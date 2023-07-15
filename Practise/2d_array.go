// Go program for 2-Dimensional Array

package main

import ("fmt")

func main() {
	fmt.Println("2D Array")

	arr := [2][2]int{{1,2},{3,4}}

	str_arr := [2][2]string{{"a","b"},{"c","d"}}

	for i:=0 ; i<2 ; i++ {
		for j:=0 ; j<2 ; j++ {
			fmt.Println(arr[i][j])
		}
	}

	for x:=0; x<2 ; x++ {
		for y:=0; y<2 ; y++ {
			fmt.Println(str_arr[x][y])
		}
	}

	var myarr[2] int; fmt.Println("Elements of the Array: ", str_arr,myarr)
}