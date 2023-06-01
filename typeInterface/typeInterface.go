package typeInterface

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Color string
}

func (d Dog) Speak() string {
	return "WoufWouf"
}

type Cat struct {
	jumpHeight int
}

func (c Cat) Speak() string {
	return "MiaouMiaou"
}

func Main() {
	animals := []Animal{
		Dog{"White"},
		Cat{2},
	}

	for _, animal := range animals {
		//t, ok := animal.(Dog)
		//fmt.Printf("Type assertion value=%v, ok = %v\n", t, ok)
		describeAnimal(animal)
	}
}

func describeAnimal(a Animal) {
	switch v := a.(type) {
	case Dog:
		fmt.Printf("We have a dog, color=%v\n", v.Color)
	case Cat:
		fmt.Printf("We have a cat, jump=%v\n", v.jumpHeight)
	}
}
