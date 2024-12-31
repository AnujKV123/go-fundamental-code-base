package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Rspons is of type: %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	// ioutil.ReadAll(response.Body) // ioutil is deprecated
	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
