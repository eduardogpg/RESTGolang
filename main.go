package main

import(
	"./models"
	"fmt"
)

func main() {
	models.CreateConnection()
	defer models.CloseConnection()
}