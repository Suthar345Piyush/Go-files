package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.Create("text.txt")
	// if err != nil {
	// 	fmt.Println("error  while creating the file", err)
	// 	return
	// }

	// defer file.Close()

	// content := "hello i am learning golang"
	// byte, errors := io.WriteString(file, content+"\n")
	// fmt.Println("the data in bytes are: ", byte)
	// if errors != nil {
	// 	fmt.Println("error while writing file", errors)
	// 	return
	// }
	// fmt.Println("the  file creation is complete")

	// file, err := os.Open("text.txt")
	// if err != nil {
	// 	fmt.Println("error in opening the  file: ", err)
	// 	return
	// }

	// defer file.Close()

	// // create  a  buffer to read the  file content
	// // 1024 is  size of buffer
	// buffer := make([]byte, 1024)

	// read the  file content into the  buffer
	// EOF is end  of  file

	// for {
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("error while reading the  file", err)
	// 		return
	// 	}

	// 	// process the  read content
	// 	fmt.Println(string(buffer[:n]))
	// }

	// short cut to read the  data from other file
	// read the entire file into a  byte  slice
	// go has  function called as ReadFile with the  os
	//generally we not using this  os method , because it  is  not capable of reading  large file , it shows  error , it  is  used for  short file reading
	// for reading large  file content we use  buffer method

	content, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println("error while reading the  file", err)
		return
	}

	fmt.Println(string(content))

}
