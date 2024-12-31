package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

func main() {

	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)
	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
}
