// Go Program for 1-Dimensional Array

package main

import "fmt"

func main(){

	arr := [5]int{10,20,30,40,50}

	fmt.Println("List of Int Array")

	for i := 0; i < 5; i++ {
		fmt.Printf("%d \n", arr[i])
	
	}

	fmt.Println("List of String Array")

	s_arr := [5]string{"a","b","c","d","e"}
	
	for i := 0; i < 5; i ++ {
		fmt.Printf(s_arr[i])
	}
	

}