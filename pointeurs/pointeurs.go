package pointeurs

import "fmt"

func Main() {
	i := 1
	var p *int = &i

	fmt.Printf("i=%v\n", i)
	fmt.Printf("p=%v\n", p)
	fmt.Printf("*p=%v\n", *p)
	fmt.Println("--------------------------")

	s := "Bob"
	sPtr := &s
	s2 := *sPtr
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("--------------------------")

	*sPtr = "Alice"
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("--------------------------")

	UpdateVal(s)
	fmt.Printf("s=%v\n", s)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Println("--------------------------")

	//UpdatePtr(&s)
	UpdatePtr(sPtr)
	fmt.Printf("s=%v\n", s)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Println("--------------------------")

}

func UpdateVal(val string) {
	val = "value"
}

func UpdatePtr(ptr *string) {
	*ptr = "pointer"
}
