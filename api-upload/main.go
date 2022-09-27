package main

import (
	"fmt"

	database "citybiker-go-api/db"
)

func main() {
	message := database.Hello()
	database.LoadConfig()
	fmt.Println(message)
}
