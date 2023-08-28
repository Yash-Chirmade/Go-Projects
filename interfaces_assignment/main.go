package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sidelength float64
}
type triangle struct {
	height float64
	base   float64
}

func main() {
	ts := triangle{90,89}
	ss := square{90}
	printArea(ts)
	printArea(ss)
}
func printArea(s shape) {
	fmt.Print(s.getArea())
}
func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
