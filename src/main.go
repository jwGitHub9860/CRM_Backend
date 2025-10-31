package main

import (
	"bufio"
	"strings"

	//"encoding/json"
	"fmt"
	//"go/reader"
	//"io/ioutil"
	"net/http"
	"os"

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
	1: {"John Doe", "Buyer", "johndoe@gmail.com", "123-456-7890", true},
	2: {"Jane Doe", "Payer", "janedoe@gmail.com", "987-654-3210", false},
}

func doesCustomerExist(customerNotFound bool) Customer {
	// Takes User Input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Customer Name: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Trim(userInput, "\r\n")

	// Checks if "userInput" Exists
	for _, customer := range customerMap {
		// MUST USE "strings.Compare(userInput, customer.name) == 0", Using "userInput == customer.name" Defines "userInput" & "customer.name" as NOT EQUAL EVEN THOUGH THEY ARE EQUAL
		if strings.Compare(userInput, customer.name) == 0 {
			customerNotFound = false
			return customer
		}
	}

	// Displays if Customer was NOT FOUND
	if customerNotFound {
		// Returns NULL VALUE for "struct"
		return Customer{}
	}

	// Fixes "missing return statement" Error
	return Customer{}
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	for _, customer := range customerMap {
		fmt.Println(customer)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	// Checks if Customer Exists
	customerExistence := doesCustomerExist(true)
	if customerExistence != (Customer{}) {
		w.WriteHeader(http.StatusAccepted)
		fmt.Print(customerExistence)
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}

/*func addCustomer(w http.ResponseWriter, r *http.Request) {
	// Checks if Customer Exists
	http.HandleFunc("/customers/{id}", getCustomer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	// Checks if Customer Exists
	http.HandleFunc("/customers/{id}", getCustomer)
}*/

func removeCustomer(w http.ResponseWriter, r *http.Request) {
	// Checks if Customer Exists
	http.HandleFunc("/customers/{id}", getCustomer)
}

func main() {
	// Accesses "index.html" as Default File
	//fileServer := http.FileServer(http.Dir("./static"))
	//http.Handle("/", fileServer)

	// Calls Functions as Handler Functions
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")
	/*router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")*/
	router.HandleFunc("/customers/{id}", removeCustomer).Methods("DELETE")

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", router)
}
