package open_close

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

type shape interface {
	area() float64
}

type calculator struct{}

func (a calculator) areaSum(shapes ...shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}

func OpenClose() {
	c := circle{9}
	s := square{7}
	t := triangle{4, 6}
	calc := calculator{}
	sumValue := calc.areaSum(c, s, t)
	fmt.Printf("The total area of shapes cirlce + sqaure + triangle = %f\n", sumValue)

}
