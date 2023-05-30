package fonction

import (
	"fmt"
	"strings"
)

func Main() {
	printInfoParams("John", 30)
	fmt.Printf("Average result=%f\n", avg(16.3, 20.5))
	fmt.Printf("Sum resultat=%d\n", sumNamedReturn(10, 25, 3))
	lowerName, len := toLowerStr("AntiCorruption")
	fmt.Printf("%s (len=%d)\n", lowerName, len)
}
func printInfoParams(name string, age int) {
	fmt.Printf("Name=%s, Age=%d\n", name, age)
}

func avg(x, y float64) float64 {
	return (x + y) / 2
}

func sumNamedReturn(x, y, z int) (sum int) {
	sum = x + y + z
	return sum
}

func toLowerStr(name string) (string, int) {
	return strings.ToLower(name), len(name)
}
