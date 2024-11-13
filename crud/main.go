package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Todo struct represents a typical to-do item structure with JSON tags
// to map the fields when encoding or decoding JSON.
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// performGetRequest sends a GET request to retrieve a single Todo item
// and displays the response.
func performGetRequest() {
	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	// Send GET request
	res, err := http.Get(myURL)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer res.Body.Close()

	// Check if status code indicates success
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in GET response:", res.Status)
		return
	}

	// Decode the response JSON into a Todo struct
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding GET response:", err)
		return
	}
	fmt.Println("GET Response Todo:", todo)
}

// performPostRequest sends a POST request to create a new Todo item
// and displays the response.
func performPostRequest() {
	// Define the Todo item to be posted
	todo := Todo{
		UserID:    11,
		Title:     "Prince Kumar",
		Completed: true,
	}

	// Marshal the Todo struct to JSON format
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	// Convert JSON data to a reader
	jsonReader := strings.NewReader(string(jsonData))
	myURL := "https://jsonplaceholder.typicode.com/todos"

	// Send POST request with JSON data
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer res.Body.Close()

	// Read and display the response
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("POST Response:", string(data))
	fmt.Println("Status Code:", res.Status)
}

// performUpdateRequest sends a PUT request to update an existing Todo item
// and displays the response.
func performUpdateRequest() {
	// Define the updated Todo item
	todo := Todo{
		UserID:    21,
		Title:     "Prince Kumar updated",
		Completed: true,
	}

	// Marshal the Todo struct to JSON format
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	// Convert JSON data to a reader
	jsonReader := strings.NewReader(string(jsonData))
	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	// Create a new PUT request
	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}

	// Set request header to indicate JSON content
	req.Header.Set("Content-Type", "application/json")

	// Use an HTTP client to send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}
	defer res.Body.Close()

	// Read and display the response
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("PUT Response:", string(data))
	fmt.Println("Status Code:", res.Status)
}


// performDeleteRequest sends a DELETE request to remove a Todo item
// and displays the response.
func performDeleteRequest() {
	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	// Create a new DELETE request
	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil {
		fmt.Println("Error creating DELETE request:", err)
		return
	}

	// Use an HTTP client to send the DELETE request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending DELETE request:", err)
		return
	}
	defer res.Body.Close()

	// Read and display the response (if any)
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("DELETE Response:", string(data))
	fmt.Println("Status Code:", res.Status)
}

// main function demonstrates the use of the performGetRequest,
// performPostRequest, and performUpdateRequest functions.
func main() {
	fmt.Println("Learning CRUD operations...")

	// Perform GET request
	fmt.Println("\nPerforming GET request:")
	performGetRequest()

	// Perform POST request
	fmt.Println("\nPerforming POST request:")
	performPostRequest()

	// Perform PUT request
	fmt.Println("\nPerforming PUT request:")
	performUpdateRequest()

	// Perform DELETE request
	fmt.Println("\nPerforming DELETE request:")
	performDeleteRequest()
}
