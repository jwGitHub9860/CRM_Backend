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

var customerMap = map[uint32]Customer{
	1: Customer{"John Doe", "Buyer", "johndoe@gmail.com", "123-456-7890", true},
	2: Customer{"Jane Doe", "Payer", "janedoe@gmail.com", "987-654-3210", false},
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	for _, customer := range customerMap {
		fmt.Println(customer)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {}

func addCustomer(w http.ResponseWriter, r *http.Request) {}

func updateCustomer(w http.ResponseWriter, r *http.Request) {}

func removeCustomer(w http.ResponseWriter, r *http.Request) {}

func main() {
	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", nil)
}
