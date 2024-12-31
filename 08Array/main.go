package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 1
	// arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)
	fmt.Println(len(arr))

	var vegList = [6]string{"Potato", "Tomato", "Cucumber", "Carrot", "Beans"}
	fmt.Println("Vegy List is: ", vegList)
	fmt.Println("Vegy List Length is: ", len(vegList))

	vegList[0] = "Brinjal"
	fmt.Println("Vegy List is: ", vegList)
}
