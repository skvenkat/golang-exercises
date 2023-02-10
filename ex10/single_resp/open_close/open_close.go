package open_close

import (
	"fmt"
	"math"
)

type circle struct {
	shape  string
	radius float32
}

func (c circle) getShape() string {
	return c.shape
}

func (c circle) area() float32 {
	return (math.Pi * c.radius * c.radius)
}

type square struct {
	shape   string
	sideLen float32
}

func (s square) getShape() string {
	return s.shape
}

func (s square) area() float32 {
	return s.sideLen * s.sideLen
}

type triangle struct {
	shape  string
	height float32
	base   float32
}

func (t triangle) getShape() string {
	return t.shape
}

func (t triangle) area() float32 {
	return ((t.base * t.height) / 2)
}

type shape interface {
	getShape() string
	area() float32
}

type outPrinter struct{}

func (op outPrinter) toText(s shape) string {
	return fmt.Sprintf("The area of %s : %f", s.getShape(), s.area())
}

type calculator struct {
	total float32
}

func (c calculator) sumAreas(shapes ...shape) float32 {
	var sum float32

	for _, shape := range shapes {
		sum += shape.area()

		/***
		switch shape.(type) {
		case circle:
			r := shape.(circle).radius
			sum += (math.Pi * r * r)
		case square:
			l := shape.(square).sideLen
			sum += (l * l)
		}
		***/
	}
	return sum
}

func OpenCloseResp() {
	fmt.Println("=== Open-Close Principle ===")

	c := circle{"Circle", 7}
	c.area()

	s := square{"Square", 5}
	s.area()

	t := triangle{"Triangle", 3, 3}
	t.area()

	op := outPrinter{}
	fmt.Println(op.toText(c))
	fmt.Println(op.toText(s))
	fmt.Println(op.toText(t))

	calc := calculator{}
	sum := calc.sumAreas(c, t, s)
	fmt.Printf("The total area of shapes : %s, %s and %s is : %f\n", c.getShape(), s.getShape(), t.getShape(), sum)

}
