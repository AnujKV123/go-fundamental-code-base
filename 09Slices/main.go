package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Orange", "Grapes"}

	fmt.Printf("type of fruitList is %T\n", fruitList)
	fmt.Println("fruitList is: ", fruitList)

	fruitList = append(fruitList, "Mango", "Peach")
	fmt.Println("fruitList is: ", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println("fruitList is: ", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("fruitList is: ", fruitList)

	fruitList = append(fruitList[:1])
	fmt.Println("fruitList is: ", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 39
	highScores[1] = 98
	highScores[2] = 57
	highScores[3] = 96

	fmt.Println("highScores is: ", highScores)

	highScores = append(highScores, 95, 94, 93)
	fmt.Println("highScores is: ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("highScores is: ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slice based on index
	var courses = []string{"reactjs", "js", "go", "python", "java"}
	fmt.Println(courses)
	courses = append(courses[:2], courses[3:]...)
	fmt.Println(courses)
}
