package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's your name")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello Mr.", name)

	//for that we use  bufio package  with reader
	// new reader and  os operating  system and  stdin is  standard input

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello mr.", name)

}
