package intf_seg

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

type cube struct {
	square
}

func (c cube) volume() float64 {
	return math.Pow(c.length, 3)
}

type shape interface {
	area() float64
}

type object interface {
	volume() float64
}

func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}

func areaVolmeSum(shapes ...object) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.volume()
	}
	return sum
}

func IntfSeg() {
	s := square{5}
	c := cube{s}
	fmt.Printf("Square's area : %f\n", s.area())
	fmt.Printf("Cube's Volume : %f\n", c.volume())
}
