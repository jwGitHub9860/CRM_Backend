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

// Did not define because it will CONSTANTLY CHANGE
var key uint32

// Did not define because it will CONSTANTLY CHANGE
var newCustomerName string

var customerMap = map[uint32]Customer{
	1: {"John Doe", "Buyer", "johndoe@gmail.com", "123-456-7890", true},
	2: {"Jane Doe", "Payer", "janedoe@gmail.com", "987-654-3210", false},
}

var whatToInputStatements = []string{
	"Enter Customer Name: ",
	"Enter Customer Role: ",
	"Enter Customer Email: ",
	"Enter Customer Phone: ",
	"Enter Customer contacted: ",
	"Choose Customer to Update (via customer name): ",
}

func inputCustomerInfo(inputPrintStatementNumber int) string {
	// Organizes Terminal Output by Keeping "whatToInputStatements" Strings Aligned to Left Side of Terminal
	// (CANNOT do anything about Initial Two New Lines when server starts)
	fmt.Println("\n")

	// Takes User Input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(whatToInputStatements[inputPrintStatementNumber])
	userInput, _ := reader.ReadString('\n')
	return strings.Trim(userInput, "\r\n")
}

func chooseCustomerInfo() bool {
	// TESTING CODE
	fmt.Println("Inside 'chooseCustomerInfo()' function")

	// Initializes Strings to EMPTY because Strings will CONSTANTLY CHANGE
	customerInfoStrings := [5]string{}

	for i := 0; i < 5; i++ {
		// TESTING CODE
		fmt.Println("Inside 'chooseCustomerInfo()' FOR LOOP")

		customerInfoStrings[i] = inputCustomerInfo(i)

		// Defines "newCustomerName" for "addCustomer()" function
		newCustomerName = customerInfoStrings[0]
	}

	// TESTING CODE
	fmt.Println("customerInfoStrings[4]:", customerInfoStrings[4])

	// Adds or Updates Customer Info
	if customerInfoStrings[4] == "true" {
		// TESTING CODE
		fmt.Println("Inside 'chooseCustomerInfo()' For Loop IF STATEMENT (true)")

		customerMap[key] = Customer{customerInfoStrings[0], customerInfoStrings[1],
			customerInfoStrings[2], customerInfoStrings[3],
			true,
		}
		return true
	} else if customerInfoStrings[4] == "false" {
		// TESTING CODE
		fmt.Println("Inside 'chooseCustomerInfo()' For Loop IF STATEMENT (false)")

		customerMap[key] = Customer{customerInfoStrings[0], customerInfoStrings[1],
			customerInfoStrings[2], customerInfoStrings[3],
			false,
		}
		return true
	} else {
		return false
	}
}

func doesCustomerExist(customerNotFound bool, userInput string) Customer {
	// Checks if "userInput" Exists
	for mapKey, customer := range customerMap {
		// MUST USE "strings.Compare(userInput, customer.name) == 0", Using "userInput == customer.name" Defines "userInput" & "customer.name" as NOT EQUAL EVEN THOUGH THEY ARE EQUAL
		if strings.Compare(userInput, customer.name) == 0 {
			customerNotFound = false

			// Defines "key" for "delete()" command to Indicate Which Customer to Remove
			key = mapKey

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
	customerExistence := doesCustomerExist(true, inputCustomerInfo(0))
	if customerExistence != (Customer{}) {
		w.WriteHeader(http.StatusAccepted)
		fmt.Print(customerExistence)
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	additionSuccessful := chooseCustomerInfo()

	// Checks if Addition is Successful (error can occur when choosing "contacted" boolean)
	if additionSuccessful {
		// Checks if New Customer ALREADY Exists
		customerExistence := doesCustomerExist(true, newCustomerName)

		if customerExistence != (Customer{}) {
			// Deletes Newly Added Customer if it's DUPLICATE
			delete(customerMap, key)

			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusAccepted)
		}
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	// Choose Customer Name to Choose which Customer to Update
	chosenCustomerName := inputCustomerInfo(5)

	// Checks if Customer Exists
	customerExistence := doesCustomerExist(true, chosenCustomerName)
	if customerExistence != (Customer{}) {
		updateSuccessful := chooseCustomerInfo()

		// Checks if Update is Successful (error can occur when choosing "contacted" boolean)
		if updateSuccessful {
			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}

func removeCustomer(w http.ResponseWriter, r *http.Request) {
	// Saves Customer Name that User Chose
	chosenCustomerName := inputCustomerInfo(0)

	// Checks if Customer Exists
	customerExistence := doesCustomerExist(true, chosenCustomerName)
	if customerExistence != (Customer{}) {
		w.WriteHeader(http.StatusAccepted)

		delete(customerMap, key)

		// Organizes Terminal Output by Preventing "print statement" & Result of Postman request from Being On the Same Line
		fmt.Println("\n")
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}

func main() {
	// Accesses "index.html" as Default File
	//fileServer := http.FileServer(http.Dir("./static"))
	//http.Handle("/", fileServer)

	// Calls Functions as Handler Functions
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", removeCustomer).Methods("DELETE")

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", router)
}
