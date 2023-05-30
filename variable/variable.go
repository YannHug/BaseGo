package variable

import (
	"fmt"
	"gotest/data"
)

func Variable() {
	var age int = 20
	fmt.Println(age)

	var weight, height int = 80, 190
	fmt.Printf(" le poid est de %d et la taille de %d \n", weight, height)

	var name = "John"
	fmt.Println(name)

	email := "john@go.org"
	fmt.Println(email)

	fmt.Println(data.Age)
}
