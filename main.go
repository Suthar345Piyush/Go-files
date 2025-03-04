// CRUD operations in golang
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/100")
	if err != nil {
		fmt.Println("error getting:", err)
		return
	}
	defer res.Body.Close()

	//we have to put some extra validation as well

	if res.StatusCode != http.StatusOK {
		fmt.Println("error is getting response", res.Status)
		return
	}

	//well this is not right convention of writing this instead this we can write this like

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("error reading:", err)
	// 	return
	// }
	// fmt.Println("Data: ", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)

	if err != nil {
		fmt.Println("Error decoding: ", err)
		return
	}
	fmt.Println("Todo: ", todo)
}

func performPostRequest() {
	todo1 := Todo{
		UserID:    245,
		Title:     "Piyush suthar",
		Completed: true,
	}

	//coverting the Todo struct into json data

	jsonData, err := json.Marshal(todo1)
	if err != nil {
		fmt.Println("error marshalling: ", err)
		return
	}
	//convert json data into string

	var jsonReader io.Reader
	jsonString := string(jsonData)

	//converting this data into reader
	jsonReader = strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos/"

	// here we are sending post request
	// this application/json is basically telling to server that we are sending json data
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("error sending request: ", err)
		return
	}

	defer res.Body.Close()

	// Check if the status code indicates success

	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status:", res.Status)
		return
	}

	//we have to convert that data into readable form(using io and ReadAll function)

	//basically we are reading response

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading response: ", err)
		return
	}

	fmt.Println("response status:", res.Status)
	fmt.Println("Our data is: ", string(data))
}

// for updating we can use both put and patch
func performUpdateRequest() {
	todo := Todo{
		UserID:    635,
		Title:     "Piyush want be pro GO dev",
		Completed: true,
	}

	// struct to json data

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error marshalling:", err)
		return
	}

	//json->string
	jsonString := string(jsonData)

	var jsonReader io.Reader
	jsonReader = strings.NewReader(jsonString)

	const myURL = "https://jsonplaceholder.typicode.com/todos/100"

	// creating put request
	req, err := http.NewRequest(http.MethodPatch, myURL, jsonReader)
	if err != nil {
		fmt.Println("error in put req: ", err)
		return
	}

	//we are giving content type on different palce
	//this  Set method has key-value pair

	req.Header.Set("Content-type", "application/json")

	//sending a request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error in sending request: ", err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error reading respose: ", err)
		return
	}
	fmt.Println("response: ", string(data))
	fmt.Println("Response status ", res.Status)

}

func performDeleteRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//creating the delete request

	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("error creating delete request:", err)
		return
	}

	//sending the request
	//Do basically sends http request and receives the http response
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("response status: ", res.Status)

}

func main() {
	fmt.Println("learning the crud in go")
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()

}
