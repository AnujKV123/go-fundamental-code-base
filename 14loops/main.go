package main

import "fmt"

func main() {

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Println("index is: ", index, "day is: ", index, day)
	}

	rogueValue := 1
	for rogueValue < 10 {

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		if rogueValue == 8 {
			break
		}

		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}
}
