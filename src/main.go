package main

import (
	"encoding/json"
	"fmt"

	"io/ioutil"

	"net/http"

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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This project involves building the backend (i.e. server-side portion) of a CRM application. The backend will allow the user make HTTP requests to the Postman server to perform CRUD operations (Create, Read, Update, and Delete). The mock customer data will allow user to perform CRUD operations in Postman application.")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// Sets content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Returns WHOLE "customerMapsForAPI" as JSON Back to User in API Response
	json.NewEncoder(w).Encode(customerMapsForAPI)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	// Sets content type to JSON
	w.Header().Set("Content-Type", "application.json")

	customerNotFound := true

	// Parse Path Parameters
	vars := mux.Vars(r)

	// Obtains "id" from Handle Function Path ("/customers/{id}")
	id := vars["id"]

	for _, customerData := range customerMapsForAPI {
		// Checks if Customer Exists
		if customerData["ID"] == id {
			customerNotFound = false

			// Returns "customerMapsForAPI"
			json.NewEncoder(w).Encode(customerData)

			w.WriteHeader(http.StatusAccepted)
		}
	}

	if customerNotFound {
		w.WriteHeader(http.StatusNotFound)
	}
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
	// Set New Content Type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Holds New Customer Data in "Customer" Form
	//var newCustomer Customer
	var newCustomer map[string]string

	// Obtains Body of POST Request from API (Input from API)
	reqBody, _ := ioutil.ReadAll(r.Body)

	// Parse JSON body
	json.Unmarshal(reqBody, &newCustomer)

	customerNotFound := true

	// Parse Path Parameters
	vars := mux.Vars(r)

	// Obtains "id" from Handle Function Path ("/customers/{id}")
	id := vars["id"]

	for index, customerData := range customerMapsForAPI {
		// Checks if Customer Exists
		if customerData["ID"] == id {
			// Removes Chosen Customer
			customerMapsForAPI = append(customerMapsForAPI[:index], customerMapsForAPI[index+1:]...)

			// Adds New UPDATED Customer Data to "customerMapsForAPI"
			customerMapsForAPI = append(customerMapsForAPI, newCustomer)

			customerNotFound = false
		}
	}

	if customerNotFound {
		w.WriteHeader(http.StatusNotFound)
	} else {
		// Returns "customerMapsForAPI"
		json.NewEncoder(w).Encode(customerMapsForAPI)

		w.WriteHeader(http.StatusAccepted)
	}
}

func removeCustomer(w http.ResponseWriter, r *http.Request) {
	// Set New Content Type to JSON
	w.Header().Set("Content-Type", "application/json")

	customerNotFound := true

	// Parse Path Parameters
	vars := mux.Vars(r)

	// Obtains "id" from Handle Function Path ("/customers/{id}")
	id := vars["id"]

	for index, customerData := range customerMapsForAPI {
		// Checks if Customer Exists
		if customerData["ID"] == id {
			// Removes Chosen Customer
			customerMapsForAPI = append(customerMapsForAPI[:index], customerMapsForAPI[index+1:]...)

			customerNotFound = false
		}
	}

	if customerNotFound {
		w.WriteHeader(http.StatusNotFound)
	} else {
		// Returns "customerMapsForAPI" & Displays "customerMapsForAPI" on API
		json.NewEncoder(w).Encode(customerMapsForAPI)

		w.WriteHeader(http.StatusAccepted)
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
