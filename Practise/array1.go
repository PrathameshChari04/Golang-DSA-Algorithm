package main

import ("fmt")

func main() {
	arr := [...]int{10,20,30}

	fmt.Println("Intilized Array : ",arr)

	arr[0] = 200

	fmt.Println("Updated Array : ",arr)

	new_arr := [3]int{40,50,60}

	fmt.Println(arr==new_arr)

}