package main

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
)

type Customer struct {
	name      string
	role      string
	email     string
	phone     string
	contacted bool
}

func getAllCustomers(customerMap map[uint32]string) {
	for id, customer := range customerMap {
		fmt.Println()
	}
}

func getCustomer(name string) {}

func addCustomer(customerMap map[uint32]string) {}

func updateCustomer(name string) {}

func removeCustomer(name string) {}

func main() {
	c1 := Customer{"John Doe", "Buyer", "johndoe@gmail.com", "123-456-7890", true}
	c2 := Customer{"Jane Doe", "Payer", "janedoe@gmail.com", "987-654-3210", false}

	customerMap := map[uint32]Customer{
		1: c1,
		2: c2,
	}

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", nil)
}
