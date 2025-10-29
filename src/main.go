package main

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
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

// CHECK ON: keep "name string" or not
func getCustomer(w http.ResponseWriter, r *http.Request) {
	customerNotFound := true
	var userInput string

	// Takes User Input
	fmt.Println("Enter Customer Name: ")
	fmt.Scanln(&userInput)

	// Checks if "userInput" Exists
	for _, customer := range customerMap {
		if customer.name == userInput {
			customerNotFound = false
			w.WriteHeader(http.StatusAccepted)
			fmt.Print(customer)
		}
	}

	// Displays if Customer was NOT FOUND
	if customerNotFound {
		w.WriteHeader(http.StatusConflict)
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	// Checks if Customer Exists
	http.HandleFunc("/customers/{id}", getCustomer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	// Checks if Customer Exists
	http.HandleFunc("/customers/{id}", getCustomer)
}

func removeCustomer(w http.ResponseWriter, r *http.Request) {
	// Checks if Customer Exists
	http.HandleFunc("/customers/{id}", getCustomer)
}

func main() {
	// Calls Functions as Handler Functions
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", removeCustomer).Methods("DELETE")

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", nil)
}
