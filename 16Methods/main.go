package main

import "fmt"

func main() {
	// no inheritance in golang; no super or parent class
	Anuj := User{"Anuj", "anuj@go.dev", true, 24}
	fmt.Println(Anuj)
	fmt.Printf("Anuj details are: %+v\n", Anuj)
	fmt.Printf("Name is %v and email is %v\n", Anuj.Name, Anuj.Email)

	Anuj.getStatus()
	Anuj.newEmail("anuj@go.dev.com")
	fmt.Printf("Name is %v and email is %v\n", Anuj.Name, Anuj.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("Status is: ", u.Status)
	// return "Status is: " + fmt.Sprint(u.Status)
}

func (u User) newEmail(email string) {
	u.Email = email
	fmt.Println("Email is: ", u.Email)
}
