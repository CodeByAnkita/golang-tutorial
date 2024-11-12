package main

import (
	"fmt"
)

func main() {
	// name <-> grade
	studentGrades := make(map[string]int)

	fmt.Println("marks of bob : ", studentGrades["bob"])
	studentGrades["bob"] = 100
	studentGrades["prince"] = 90
	studentGrades["bob"] = 85
	studentGrades["ankita"] = 80
	studentGrades["david"] = 65
	studentGrades["Alice"] = 50
	studentGrades["Charile"] = 55

	fmt.Println("marks of bob : ", studentGrades["bob"])
	studentGrades["bob"] = 100
	fmt.Println("new marks of bob : ", studentGrades["bob"])

	// delete(studentGrades, "Charile")
	fmt.Println("marks of bob ", studentGrades["bob"])

	// Checking if a key exists
	grades, exists := studentGrades["ankita"]
	fmt.Println("Grades of ankita: ", grades)
	fmt.Println("ankitas exists : ", exists)
	// fmt.Println("Marks of David: ", studentGrades ["alice"])
	// Checking if a key exists
	Grades, Exists := studentGrades["ankita"]
	fmt.Println("Grades of Charile : ", Grades)
	fmt.Println("alice exists: ", Exists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n ", index, value)
	}

	person := map[string]int{
		"Alice": 90,
		"Bob":   80,
        "Charlie": 75,
	}

	for index, value := range person{
		fmt.Printf("............Key is %s and marks is %d\n ", index, value)
	}
}
