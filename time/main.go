package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)
	fmt.Printf("type of currentTime %T\n", currentTime)

	formatted := currentTime.Format("02-01-2006, Monday, 15:04:05 ") //they uses this formate that is their launch date and time
	formattedTime := currentTime.Format("2006/01/02, 3:04 PM ")      //for am or pn you need write same as

	fmt.Println("Formatted time:", formatted)
	fmt.Printf("type of formatted %T\n", formattedTime)

	layout_str := "2006-01-02"
	dateStr := "2023-11-25"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted time from string:", formatted_time)

	//add 1 more day in currentTime
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("New date after adding 1 day:", new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("Formatted new date after adding 1 day:", formatted_new_date)

}
