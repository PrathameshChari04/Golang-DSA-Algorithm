package main

import "fmt"

func main() {

	customer1 := Customer{"Customer1","customer1@email.com","9900990099",true,24}

	fmt.Println("Details : ", customer1)
	fmt.Printf(customer1.Username)
	fmt.Printf("Details : %+v\n", customer1)
	fmt.Printf("Customer1 Age is %v", customer1.Age)

}

type Customer struct {
	Username  string
	Email     string
	Mobile_no string
	Status    bool
	Age       int
}
