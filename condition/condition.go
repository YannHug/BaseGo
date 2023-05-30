package condition

import (
	"fmt"
	"gotest/data"
)

func Condition() {
	if data.Age > 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	switch data.Age {
	case 18:
		fmt.Println("You are 18")
	case 20:
		fmt.Println("You are 20")
	default:
		fmt.Println("You are not 18 or 20")
	}
}
