package tableau

import "fmt"

func Main() {
	var names [3]string
	names[0] = "Bob"
	names[2] = "Alice"
	fmt.Printf("names=%v (len=%v)\n", names, len(names))

	odds := [4]int{1, 3, 5, 7}
	fmt.Printf("odds=%v (len=%d)\n", odds, len(odds))
}
