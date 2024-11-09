package main

import (
	"fmt"
	"learning/myutil"
)

func main() {
	fmt.Print("Welcome to golang")

	myutil.PrintMessage("hello world")

	var name string = "ankita"
	var version = 75
	fmt.Println(name)
	fmt.Println(version)

	var money int = 6800
	var currency = 852
	fmt.Println(money )
	fmt.Println( currency)

	var dimension float64 = 87.24
fmt.Printf("The value of dimension is: %.2f\n", dimension)

	var decided bool = false
	fmt.Println(decided)

	// var person = 23
	// fmt.Println("You are %d years old \n", person)

	const pi = 97.12
	fmt.Println(pi)

	person := "ankita kapadiya"
fmt.Println(person)
}
