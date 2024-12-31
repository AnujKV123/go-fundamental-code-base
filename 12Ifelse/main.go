package main

import "fmt"

func main() {

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println("Negative number")
	} else if num > 10 {
		fmt.Println("Large number")
	} else {
		fmt.Println("Small number")
	}
}
