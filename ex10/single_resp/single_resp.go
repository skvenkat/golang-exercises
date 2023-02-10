package single_resp

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

type shape interface {
	getShape() string
	area() float32
}

type outPrinter struct{}

func (op outPrinter) toText(s shape) string {
	return fmt.Sprintf("The area of %s : %f", s.getShape(), s.area())
}

func SingleResp() {
	fmt.Println("=== Single Responsibility ===")
	c := circle{"Circle", 5}
	c.area()

	s := square{"Square", 2}
	s.area()

	op := outPrinter{}
	fmt.Println(op.toText(c))
	fmt.Println(op.toText(s))
}
