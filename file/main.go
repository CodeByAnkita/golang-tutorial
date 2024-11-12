package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("error while creating file", err)
			return
		}
		defer file.Close()

		//write content in file
		// fmt.Fprintf(file, "Hello, World!\n")
		content := "Hello Ankita"
		// _,errors:= io.WriteString(file, content+"\n")
		byte, errors := io.WriteString(file, content+"\n")
		fmt.Println("byte written: ", byte)

		if errors != nil {
			fmt.Println("error while writing to file", err)
			return
		}

		fmt.Println("successfully created file")
	**/
	//open file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("error while opening file", err)
		return
	}
	defer file.Close()

	// Create a buffer to read the file content
	buffer := make([]byte, 1024)

	// Read the file content into the buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file", err)
			return
		}
		// process the read content
		fmt.Printf("%s", buffer[:n])
		// fmt.Println(string(buffer[:n]))
	}
	// Read the entire file into a byte slice
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file ", err)
		return
	}
	fmt.Println(string(content))

}
