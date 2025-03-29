package main

import "fmt"

type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLenght float64
}

type Shape interface {
	printArea() float64
}

func main() {
	t := Triangle{height: 10, base: 10}
	s := Square{sideLenght: 10}

	fmt.Println("Area of the triangle:", t.getArea())
	fmt.Println("Area of the square:", s.getArea())
}

func (t *Triangle) getArea() float64 {
	height, base := t.height, t.base

	return 0.5 * base * height
}

func (s *Square) getArea() float64 {
	sl := s.sideLenght

	return sl * sl
}
