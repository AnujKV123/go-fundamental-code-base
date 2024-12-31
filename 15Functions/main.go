package main

import "fmt"

func main() {

	result := adder(3, 4)
	fmt.Println("Result is: ", result)

	proResult, message := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Pro Result is: ", proResult)
	fmt.Println("Pro Result message is: ", message)

	greeter()
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "this is from proAdder"
}

func greeter() {
	fmt.Println("Hello World")
}
