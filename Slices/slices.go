package main

import (
	"fmt"
	"sort"
)

func main() {

	var languagesList = []string{"C", "C++", "C#", "Python"}

	fmt.Println("List of programming Languages : ", languagesList)

	languagesList = append(languagesList, "Javascript", "Java")

	fmt.Println(languagesList)

	languagesList = append(languagesList[0:])

	fmt.Println(languagesList)

	languagesList = append(languagesList[1:3])

	fmt.Println(languagesList)

	marks := make([]int,4)

	marks[0] = 20
	marks[1] = 22
	marks[2] = 25
	marks[3] = 30

	fmt.Println("List of Marks :", marks)

	marks = append(marks, 21, 20, 19)

	fmt.Println(marks)

	sort.Ints(marks)

	fmt.Println(" Sorted List of Marks : ", marks)

	fmt.Println("Is the List sorted : ", sort.IntsAreSorted(marks))

	// Remove elements from a slice

	var List = []string{"1","2","3","4","5"}

	fmt.Println(List)

	var index int = 2

	List = append(List[:index], List[index+1:]...)

	fmt.Println(List)

	




}