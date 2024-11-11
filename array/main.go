package main

import (
    "fmt"
)

func main() {
    fmt.Println("We are array GoLang!")

    var name[5] string
        name[0] = "prince"
        name[1] = "bob"
        name[2] = "ankita"
        name[3] = "aakash"
        name[4] = "alice"

        fmt.Println("Names of Person is :", name)
    
		var numbers =[5]int{1, 2, 3, 4}
		fmt.Println("Numbers are :", numbers)


		fmt.Println("Length of Numbers array is :", len(numbers))
		fmt.Println("Value of name at 2 index is :",name[2])
		fmt.Printf("name Array is %q/n",name)

}

