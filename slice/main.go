package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 3, 10,11,12,13,14,15,16,17,18)
	// fmt.Println(numbers)
	// fmt.Printf("Numbers has data type:%T\n", numbers)
	// fmt.Println("Length of numbers:", len(numbers))

	// name :=[]string{}
	// fmt.Println("name:",name)

	numbers := make([]int, 3, 5)

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 6)
	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 6)
	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 6)

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

}
