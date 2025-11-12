package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	//"go/reader"
	"io/ioutil"
	"strings"

	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        uint32
	name      string
	role      string
	email     string
	phone     string
	contacted bool
}

// Undefined because "key" will CONSTANTLY CHANGE
var key uint32

// Undefined because "stringKey" will CONSTANTLY CHANGE & Key for "customerMapsForAPI" Map
var stringKey string

// Undefined because "updateKey" will CONSTANTLY CHANGE & Key for Updating "customerMap" Map
var updateKey uint32

var whatToInputStatements = []string{
	"Enter Customer Name: ",
	"Enter Customer Role: ",
	"Enter Customer Email: ",
	"Enter Customer Phone: ",
	"Enter Customer contacted: ",
	"Choose Customer to Update (via customer name): ",
}

// Map for Terminal & Other Functions
var customerMap = map[uint32]Customer{
	1: {1, "John Doe", "Buyer", "johndoe@gmail.com", "123-456-7890", true},
	2: {2, "Jane Doe", "Payer", "janedoe@gmail.com", "987-654-3210", false},
	3: {3, "Jill Dole", "Payer", "jilldole@gmail.com", "012-345-6789", true},
}

// Keys & Values MUST BE STRINGS Because JSON does NOT SUPPORT "uint32" AND/OR "structs"
// (Need to Make CUSTOM Unmarshal Function to Display "customerMap" onto API as JSON Response)
var customerMapsForAPI = []map[string]string{
	{
		"ID":        "1",
		"Name":      "John Doe",
		"Role":      "Buyer",
		"Email":     "johndoe@gmail.com",
		"Phone":     "123-456-7890",
		"Contacted": "true",
	},
	{
		"ID":        "2",
		"Name":      "Jane Doe",
		"Role":      "Payer",
		"Email":     "janedoe@gmail.com",
		"Phone":     "987-654-3210",
		"Contacted": "false",
	},
	{
		"ID":        "3",
		"Name":      "Jill Dole",
		"Role":      "Payer",
		"Email":     "jilldole@gmail.com",
		"Phone":     "012-345-6789",
		"Contacted": "true",
	},
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

func doesCustomerExist(customerNotFound bool, userInput string) Customer {
	// Checks if "userInput" Exists
	for mapKey, customer := range customerMap {
		// Defines "key" for "delete()" command to Indicate Which Customer to Remove
		key = mapKey

		// MUST USE "strings.Compare(userInput, customer.name) == 0", Using "userInput == customer.name" Defines "userInput" & "customer.name" as NOT EQUAL EVEN THOUGH THEY ARE EQUAL
		if strings.Compare(userInput, customer.name) == 0 {
			customerNotFound = false

			// Defines "updateKey" for "updateCustomer()" function When Adding New Customer to "customerMap"
			updateKey = mapKey

			// Defines "stringKey" for "getCustomer()" function to Indicate Which Customer Data to Display on API
			stringKey = strconv.FormatUint(uint64(mapKey), 10)

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

	// Defines "key" & "stringKey" to allow "addCustomer()" function to Add New Customer to "customerMap" & "customerMapsForAPI"
	if addingNewCustomer {
		key += 1
	} else {
		key = updateKey
	}
	stringKey = strconv.FormatUint(uint64(key), 10)

	// Adds or Updates Customer Info
	switch customerInfoStrings[4] {
	case "true":
		customerMap[key] = Customer{
			key, customerInfoStrings[0], customerInfoStrings[1], customerInfoStrings[2], customerInfoStrings[3], true,
		}

		customerInput := map[string]string{
			"ID":        stringKey,
			"Name":      customerInfoStrings[0],
			"Role":      customerInfoStrings[1],
			"Email":     customerInfoStrings[2],
			"Phone":     customerInfoStrings[3],
			"Contacted": "true",
		}
		customerMapsForAPI[key] = customerInput

		return true
	case "false":
		customerMap[key] = Customer{
			key, customerInfoStrings[0], customerInfoStrings[1], customerInfoStrings[2], customerInfoStrings[3], false,
		}

		customerInput := map[string]string{
			"ID":        stringKey,
			"Name":      customerInfoStrings[0],
			"Role":      customerInfoStrings[1],
			"Email":     customerInfoStrings[2],
			"Phone":     customerInfoStrings[3],
			"Contacted": "false",
		}
		customerMapsForAPI[key] = customerInput

		return true
	default:
		fmt.Println("\nCustomer contacted must be either \"true\" or \"false\".")
		return false
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This project involves building the backend (i.e. server-side portion) of a CRM application. The backend will allow the user make HTTP requests to the Postman server to perform CRUD operations (Create, Read, Update, and Delete). The mock customer data will allow user to perform CRUD operations in Postman application.")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// Organizes Terminal Output by Keeping "whatToInputStatements" Strings Aligned to Left Side of Terminal
	// (CANNOT do anything about Initial Two New Lines when server starts)
	fmt.Println("\n")

	// Displays "customerMap" onto Terminal IN "ID" ORDER
	// 1. Create Slice of Structs to Hold "customerMap" Values
	var sortedMap []Customer
	for _, v := range customerMap {
		sortedMap = append(sortedMap, v)
	}
	// 2. Sort Slice
	sort.Slice(sortedMap, func(i, j int) bool { return sortedMap[i].ID < sortedMap[j].ID })
	// 3. Print Sorted Slice of "customerMap" ("sortedMap")
	for _, data := range sortedMap {
		// MUST LIST EVERY "data" Struct Type SPECIFICALLY to Display Customer in "customerMap"
		fmt.Println("ID:", data.ID)
		fmt.Println("Name:", data.name)
		fmt.Println("Role:", data.role)
		fmt.Println("Email:", data.email)
		fmt.Println("Phone:", data.phone)
		fmt.Println("Contacted:", data.contacted, "\n")
	}

	// Returns "customerMapsForAPI" as JSON Back to User in API Response
	// 1. Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// 2. Keep track of new entry -> Holds UNMARSHALED Data & MUST BE "map[string]string" or Entry will NOT UNMARSHAL CORRECTLY
	var newEntry_test_3 map[string]string

	// 3. Read the request -> Reads all Data from "reader"
	reqBody, readRequestError := ioutil.ReadAll(r.Body)
	if readRequestError != nil {
		fmt.Print("Failed to Read Request")
	}

	// 4. Parse JSON body
	json.Unmarshal(reqBody, &newEntry_test_3)

	// 5. Add new entry to "customerMapsForAPI"
	for _, customerData := range customerMapsForAPI {
		for k, v := range newEntry_test_3 {
			// Responds with conflict if entry exists
			if _, ok := customerData[k]; ok {
				w.WriteHeader(http.StatusConflict)
			} else {
				// Responds with OK if entry does not already exist
				customerData[k] = v
				w.WriteHeader(http.StatusCreated)
			}
		}
	}

	// 6. Returns "customerMapsForAPI"
	json.NewEncoder(w).Encode(customerMapsForAPI)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	// Checks if Customer Exists
	customerExistence := doesCustomerExist(true, inputCustomerInfo(0))
	if customerExistence != (Customer{}) {
		fmt.Print(customerExistence, "\n")
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

	// Encodes Customer as JSON
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(customerMapsForAPI[key])
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	// 1. Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// (2.) Holds New Customer Data in "map[string]string" Form
	var newCustomer map[string]string

	// (3.) Obtains Body of POST Request from API (Input from API)
	reqBody, _ := ioutil.ReadAll(r.Body)

	// 4. Parse JSON body
	json.Unmarshal(reqBody, &newCustomer)

	// (5.) Adds New Customer Data to "customerMapsForAPI"
	customerMapsForAPI = append(customerMapsForAPI, newCustomer)

	// 6. Returns "customerMapsForAPI"
	// SORT OF Checks if Addition is Successful (error can occur when choosing "contacted" boolean)
	json.NewEncoder(w).Encode(customerMapsForAPI)
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
		w.WriteHeader(http.StatusNotFound)
	}
}

func removeCustomer(w http.ResponseWriter, r *http.Request) {
	customerNotFound := true

	// Parse Path Parameters
	vars := mux.Vars(r)

	// Obtains "id" from Handle Function Path ("/customers/{id}")
	id := vars["id"]

	for _, customerData := range customerMapsForAPI {
		fmt.Print("Inside FOR LOOP")
		//for index, customer := range customerData {
		//fmt.Print("id:", id)
		//fmt.Print("\nindex:", index)
		//fmt.Print("\ncustomer:", customer)
		//fmt.Print("\ncustomerData:", customerData)
		//fmt.Print("\ncustomerData[id]:", customerData[id])
		//fmt.Print("\ncustomerData[\"ID\"]:", customerData["ID"], "\n")
		// Checks if Customer Exists
		if customerData["ID"] == id {
			fmt.Print("Inside IF STATEMENT")

			intId, _ := strconv.Atoi(id)

			// Removes Chosen Customer
			customerMapsForAPI = append(customerMapsForAPI[:intId], customerMapsForAPI[intId+1:]...)

			// Organizes Terminal Output by Preventing "print statement" & Result of Postman request from Being On the Same Line
			fmt.Println("\n")

			customerNotFound = false

			w.WriteHeader(http.StatusAccepted)
		}
		//}
	}

	if customerNotFound {
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	// Accesses "index.html" as Default File
	//fileServer := http.FileServer(http.Dir("./static"))
	//http.Handle("/", fileServer)

	// Calls Functions as Handler Functions
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", removeCustomer).Methods("DELETE")

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", router)
}
