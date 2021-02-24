package main

import "fmt"

type Shaper interface {
	Area() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func main() {
	var r Shaper = Rect{
		width:  5.0,
		height: 4.0,
	}

	fmt.Printf("type of r is %T\n", r)
	fmt.Printf("value of r is %v\n", r)
	fmt.Printf("value of r area is %.2f\n\n", r.Area())
}
