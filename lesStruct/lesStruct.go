package lesStruct

import (
	"fmt"
	"gotest/lesStruct/Embeded"
	"gotest/lesStruct/player"
	"gotest/lesStruct/user"
)

type Address struct {
	street, city string
}

type Person struct {
	Name string
	Age  int
	Addr Address
}

func Main() {
	var p Person
	p.Name = "Bob"
	p.Addr.city = "Lyon"

	var p1 player.Player
	p1.Name = "Paul"
	p1.Age = 10
	fmt.Printf("Player 1 : %v\n", p1)

	av := player.Avatar{"http://avatar.jpg"}
	fmt.Printf("Avatar : %v\n", av)

	p3 := player.Player{
		Name: "Paulo",
		Age:  15,
		Avatar: player.Avatar{
			Url: "http://url.com",
		},
	}
	fmt.Printf("Player 3 : %v\n", p3)

	p4 := player.New("Bobette")
	p4.Avatar = av
	fmt.Printf("Player 4 : %v\n", p4)

	u := user.User{
		Name:  "Bob",
		Email: "bob@golang.org",
	}
	fmt.Printf("User : %v\n", u)

	a := user.Admin{
		Level: 2,
		User: user.User{
			Email: "alice@golang.org",
		},
	}
	a.Name = "Alice"
	fmt.Printf("Admin : %v\n", a)

	Embeded.Main()

}
