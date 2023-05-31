package lesMap

import (
	"fmt"
)

type Vector struct {
	x, y int
}

type User struct {
	name string
}

type Key struct {
	ID   int
	Name string
}

func Main() {
	//vecs := make(map[Vector]string)
	//fmt.Println(vecs)

	m := make(map[string]int)
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	m["hello"] = 5
	m["goodbye"] = 10
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	j, ok := m["helo"]
	fmt.Printf("j=%v, ok=%v\n", j, ok)

	m["beatles"] = 2
	if _, ok := m["beatles"]; ok {
		fmt.Println("Beatles key exist")
	}

	delete(m, "goodbye")
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	m3 := map[string]int{
		"Bob":     10,
		"Alice":   15,
		"Bobette": 20,
		"John":    7,
	}

	for name, age := range m3 {
		fmt.Printf("name=%v, age=%v\n", name, age)
	}

	mu := map[string]*User{
		"HR":  {"BOB"},
		"CEO": {"ALICE"},
	}
	fmt.Println(mu["HR"])

	hr := mu["HR"]
	hr.name = "John"
	fmt.Println(mu["HR"])

	mu["CFO"] = &User{"Bobette"}

	res := make(map[Key]string)
	res[Key{1, "AZE"}] = "file1"

	k := Key{2, "ERT"}
	res[k] = "file2"
}
