package main

import "fmt"

func main() {
	var superheros [5]string

	superheros[0] = "Batman"
	superheros[1] = "Ironman"
	superheros[2] = "Superman"
	superheros[4] = "Wolverine"

	fmt.Println("List of Super-Heros : ", superheros)

	fmt.Println("List of Super-Heros : ", len(superheros))

	var languages = [4]string{"C","C++","C#","objective-c"}
	fmt.Println("Programming Languages : ", languages)
}