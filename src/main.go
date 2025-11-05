package main

import (
	"bufio"
	"encoding/json"
	"fmt"

	//"go/reader"
	"io/ioutil"
	"strings"

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

// Undefined because "key" will CONSTANTLY CHANGE
var key uint32

// MUST BE EMPTY to Display "customerMap" on API in JSON
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

func chooseCustomerInfo(addingNewCustomer bool) bool {
	// Initializes Strings to EMPTY because Strings will CONSTANTLY CHANGE
	customerInfoStrings := [5]string{}

	for i := 0; i < 5; i++ {
		customerInfoStrings[i] = inputCustomerInfo(i)
	}

	// Checks if New Customer OR Updated Customer ALREADY Exists for "addCustomer()" function
	if doesCustomerExist(true, customerInfoStrings[0]) != (Customer{}) {
		fmt.Println("\nCustomer already exists.")
		return false
	}

	// Defines "key" to allow "addCustomer()" function to Add New Customer to "customerMap"
	if addingNewCustomer {
		key = 3
	}

	// Adds or Updates Customer Info
	if customerInfoStrings[4] == "true" {
		customerMap[key] = Customer{customerInfoStrings[0], customerInfoStrings[1],
			customerInfoStrings[2], customerInfoStrings[3],
			true,
		}
		return true
	} else if customerInfoStrings[4] == "false" {
		customerMap[key] = Customer{customerInfoStrings[0], customerInfoStrings[1],
			customerInfoStrings[2], customerInfoStrings[3],
			false,
		}
		return true
	} else {
		fmt.Println("\nCustomer contacted must be either \"true\" or \"false\".")
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
	// Organizes Terminal Output by Keeping "whatToInputStatements" Strings Aligned to Left Side of Terminal
	// (CANNOT do anything about Initial Two New Lines when server starts)
	fmt.Println("\n")

	// Displays "customerMap" onto Terminal
	for _, customer := range customerMap {
		fmt.Println(customer)
	}

	// Returns "customerMap" as JSON Back to User in API Response
	// 1. Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// TESTING CODE
	fmt.Println("PASSED Step 1")

	// 2. Keep track of new entry -> Holds UNMARSHALED Data & MUST BE "map[string]string" or Entry will NOT UNMARSHAL CORRECTLY
	var newEntry_test_3 map[string]string

	// TESTING CODE
	fmt.Println("PASSED Step 2")

	// 3. Read the request -> Reads all Data from "reader"
	reqBody, readRequestError := ioutil.ReadAll(r.Body)
	if readRequestError != nil {
		fmt.Print("Failed to Read Request")
	}

	// TESTING CODE
	fmt.Println("PASSED Step 3")

	// 4. Parse JSON body
	json.Unmarshal(reqBody, &newEntry_test_3)

	// TESTING CODE
	fmt.Println("PASSED Step 4")

	// 5. Add new entry to "customerMap"
	for k, v := range newEntry_test_3 {
		// Responds with conflict if entry exists
		if _, ok := initial_customer_data[k]; ok {
			w.WriteHeader(http.StatusConflict)
		} else {
			// Responds with OK if entry does not already exist
			initial_customer_data[k] = v
			w.WriteHeader(http.StatusCreated)
		}
	}

	// TESTING CODE
	fmt.Println("PASSED Step 5")

	// 6. Returns "customerMap"
	//json.NewEncoder(w).Encode(customerMap)
	json.NewEncoder(w).Encode(initial_customer_data)

	// TESTING CODE
	fmt.Println("PASSED Step 6")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	// Encodes Customer as JSON
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(customerMap)

	// Checks if Customer Exists
	customerExistence := doesCustomerExist(true, inputCustomerInfo(0))
	if customerExistence != (Customer{}) {
		fmt.Print(customerExistence, "\n")
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	// Checks if Addition is Successful (error can occur when choosing "contacted" boolean)
	if chooseCustomerInfo(true) {
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	// Checks if Customer Exists & "inputCustomerInfo(5)" -> Choose Customer Name to Choose which Customer to Update
	if doesCustomerExist(true, inputCustomerInfo(5)) != (Customer{}) {
		// Checks if Update is Successful (error can occur when choosing "contacted" boolean)
		if chooseCustomerInfo(false) {
			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}

func removeCustomer(w http.ResponseWriter, r *http.Request) {
	// "inputCustomerInfo(0)" -> Saves Customer Name that User Chose & Checks if Customer Exists
	if doesCustomerExist(true, inputCustomerInfo(0)) != (Customer{}) {
		delete(customerMap, key)

		// Organizes Terminal Output by Preventing "print statement" & Result of Postman request from Being On the Same Line
		fmt.Println("\n")

		w.WriteHeader(http.StatusAccepted)
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
