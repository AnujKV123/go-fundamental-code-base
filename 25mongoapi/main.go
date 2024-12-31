package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnujKV123/golang-mongodb-api-implementation/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Listening on port 5000 ....")
}
