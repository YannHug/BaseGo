package Embeded

import "fmt"

type Rect struct {
	Width, Height int
}

func (r Rect) Area() int {
	return r.Width * r.Height
}

func Main() {
	r := Rect{2, 4}
	fmt.Printf("Rect Area = %v\n", r.Area())
}
