package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	fmt.Println("Hello to-do-list server")
}
