package main

import "fmt"

func main() {
	// no inheritance in golang; no super or parent class
	Anuj := User{"Anuj", "anuj@go.dev", true, 24}
	fmt.Println(Anuj)
	fmt.Printf("Anuj details are: %+v\n", Anuj)
	fmt.Printf("Name is %v and email is %v", Anuj.Name, Anuj.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
