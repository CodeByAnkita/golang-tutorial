package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("sayHello function ended successfully!")
}
func sayHi() {
	fmt.Println("Hi prince:)!")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("sayHi function ended successfully!")
}

func main() {
	fmt.Println("Leaning outgolang...")

	go sayHello()
	go sayHi()
	// wait for a moment to allow goroutines to finish
	time.Sleep(3000 * time.Millisecond)
}
