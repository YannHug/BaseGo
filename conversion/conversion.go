package conversion

import "fmt"

func Conversion() {
	var percentage float64 = 2.4
	fmt.Printf("Percentage: %f%%\n", percentage)
	fmt.Printf("Percentage: %d%%\n", int(percentage))
}
