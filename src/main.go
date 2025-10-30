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

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	for _, customer := range customerMap {
		fmt.Println(customer)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	customerNotFound := true
	//var userInputInStringForm string

	// Takes User Input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Customer Name: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Trim(userInput, "\r\n")

	// TEST PRINT STATEMENT
	/*fmt.Println("User Input: ", userInput)
	fmt.Printf("User Input TYPE: %T\n", userInput)*/

	// Checks if "userInput" Exists
	for _, customer := range customerMap {
		// TEST PRINT STATEMENTS
		//fmt.Println("Customer: ", customer)
		fmt.Println("Customer Name: ", customer.name)
		fmt.Println("User Input: ", userInput)
		/*fmt.Println("User Input IN STRING FORM: ", userInputInStringForm)
		fmt.Printf("Customer Name TYPE: %T\n", customer.name)
		fmt.Printf("PRESENTLY User Input TYPE: %T\n", userInput)
		fmt.Printf("User Input IN STRING FORM TYPE: %T\n", userInputInStringForm)*/

		// MUST USE "strings.Compare(userInput, customer.name) == 0", Using "userInput == customer.name" Defines "userInput" & "customer.name" as NOT EQUAL EVEN THOUGH THEY ARE EQUAL
		if strings.Compare(userInput, customer.name) == 0 {
			// TEST PRINT STATEMENT
			fmt.Println("INSIDE if statement")

			customerNotFound = false
			w.WriteHeader(http.StatusAccepted)
			fmt.Print(customer)

			// TEST PRINT STATEMENT
			fmt.Println("EXITING if statement")

			break
		}
	}
	// TEST PRINT STATEMENT
	fmt.Println("EXITED for loop")

	// Displays if Customer was NOT FOUND
	if customerNotFound {
		// TEST PRINT STATEMENT
		fmt.Println("INSIDE Customer Not Found if statement")

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
}

func removeCustomer(w http.ResponseWriter, r *http.Request) {
	// Checks if Customer Exists
	http.HandleFunc("/customers/{id}", getCustomer)
}*/

func main() {
	// Accesses "index.html" as Default File
	//fileServer := http.FileServer(http.Dir("./static"))
	//http.Handle("/", fileServer)

	// Calls Functions as Handler Functions
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")
	/*router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", removeCustomer).Methods("DELETE")*/

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", router)
}
