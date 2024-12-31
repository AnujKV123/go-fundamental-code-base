package main

import "fmt"

const LoginToken string = "9f5nvy9u5cqut53nv939chn9yt94uv999rc9qut0e" //public
// note : in go we create public variable using capital letters

func main() {
	var username string = "Anuj"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 225
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var Val int = 225
	fmt.Println(Val)
	fmt.Printf("Variable is of type: %T \n", Val)

	var smallFloat float32 = 225.894994796749769549
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// defalt value and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	// implicit type/way
	var website = "google.com"
	fmt.Println(website)

	// no var style
	numberOfUsers := 3000000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	// fmt.Printf("Variable is of type:", LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
