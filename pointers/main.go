package main

import (
	"fmt"
)
func modifyValueReference(num *int){
	*num= *num +20
}
func main() {
	//declare a variable and a Pointer
	var num int = 2
	var ptr *int
	//initialize the pointer with the address of the variable
	ptr = &num

	//short method
	// num:=2
	// num:="Prince"
	// ptr:=&num

	//print the value of the variable and the pointer
	fmt.Println("Value of num is:", num)
	// fmt.Println("Address of num is:", &num)
	// fmt.Println("Value of ptr is:", ptr)
	fmt.Println("Value pointed by ptr is:", *ptr)

	// //variable declaration and assignment on the same line
	// data := "ankita"
	// pointer := &data
	// fmt.Println("Value through pointer:", *pointer)

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
}

value :=10
modifyValueReference(&value)
fmt.Println("Value contains:", value)
}
