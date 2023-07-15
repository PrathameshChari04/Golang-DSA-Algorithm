package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["Python"] = "Django, Flask, FastAPI"
	languages["Javascript"] = "Node.js"
	languages["Ruby"] = "Ruby on Rails"
	languages["Java"] = "Spring-Boot"

	fmt.Println("List of all languages and their respective frameworks :", languages)

	///view a value using it key

	fmt.Println("Python Based Web Framework :", languages["Python"])

	///delete a value using it key

	delete(languages, "Ruby")

	fmt.Println("List of all languages and their respective frameworks :", languages)

	for key, value := range languages {
		fmt.Printf("Language : %v, Framework : %v\n", key,value)
	}

}