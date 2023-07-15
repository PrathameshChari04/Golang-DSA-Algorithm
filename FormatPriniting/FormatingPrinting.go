package main

import "fmt"

func main(){
	fmt.Println("Formating Printing !")

	var name string = "Namaste"
	var isActive bool = true
	x := 5;
	//active := true

	fmt.Printf("%s\n",name)
	fmt.Printf("%T\n",isActive)
	fmt.Printf("%t \n",isActive)
	fmt.Printf("%d \n",x)
	fmt.Printf("%b \n",1)
	fmt.Printf("%c \n", 34)

}