package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learning webrequest in golang")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("the error getting  get response", err)
		return
	}

	defer res.Body.Close()

	// type of  response , its come  with some  http response
	fmt.Printf("the  type  of  response %T\n", res)

	// reading the response of the  body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("the error getting for reading", err)
	}

	fmt.Println("the response is: ", string(data))

	// if  dont write string before the data , then we get data in array of  bytes form

	// how response look like
	// fmt.Println("the  response is: ", res)

	// resource management is  most important thing  in golang
}

// always error handling in golang , it is  usefull that our code should not break
