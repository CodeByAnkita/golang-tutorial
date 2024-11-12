package main

import "fmt"

// Define a struct named person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}
type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	State string
}
type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var Prince Person
	fmt.Println("Prince Person:", Prince)
	Prince.FirstName = "Prince"
	Prince.LastName = "Kapadiya"
	Prince.Age = 25
	fmt.Println("Prince Person:", Prince)

	Person1 := Person{
		FirstName: "Alice",
		LastName:  "Smith",
		Age:       30,
	}
	fmt.Println("Person1:", Person1)

	Person2 := Person{
		FirstName: "Bob",
		LastName:  "Johnson",
		Age:       28,
	}
	fmt.Println("Person2:", Person2)

	//new keyword
	var Person3 = new(Person)
	Person3.FirstName = "simaran"
	Person3.LastName = "kaps"
	Person3.Age = 27
	fmt.Println("person3:", Person3.FirstName)
	fmt.Println("Person3:", Person3)

	employee1 := Employee{
		Person_Details: Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       35,
		},
		Person_Contact: Contact{
			Email: "john@example.com",
			Phone: "+1 123-456-7890",
		},
		Person_Address: Address{
			House: 123,
			Area:  "Main Street",
			State: "NY",
		},
	}
	fmt.Println(employee1)

}
