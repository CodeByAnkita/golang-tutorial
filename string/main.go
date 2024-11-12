package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "apple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three four two two five three"
	count := strings.Count(str, "two")
	fmt.Println("Number of times 'two' appears in the string:", count)

	str = "      Hello,  go!   "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed string:", trimmed)

	str1 := "Prince"
	str2 := "Kapadiya"
	results := strings.Join([]string{str1, str2}, " ")
	fmt.Println("Combined string:", results)
}
