package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	content := "Hello guys, I am learning go lang! and this is my first file!"

	file, err := os.Create("./myGoFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./myGoFile.txt")
}

func readFile(fileName string) {
	// dataBytes, err := ioutil.ReadFile(fileName) Note: ioutil is deprecated
	dataBytes, err := os.ReadFile(fileName)

	chechNilErr(err)

	fmt.Println("Text data is: \n", string(dataBytes))
}

func chechNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
