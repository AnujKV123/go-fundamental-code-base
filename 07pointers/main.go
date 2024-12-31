package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("value of pointer is", ptr)
	a := 2

	var myPtr = &a
	fmt.Println("value of actual pointer is", myPtr)
	fmt.Println("value of actual pointer is", *myPtr)

	*myPtr = *myPtr * 20
	fmt.Println("new value is", a)

	var b *int = &a
	fmt.Println(*b)
	*b = 10
	fmt.Println(a)
}
