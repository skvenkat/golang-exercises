package main

import (
	"fmt"
)

type Arithmetic struct {
	x int
	y int
}

func (a *Arithmetic) add() int {
	return a.x + a.y
}

func (a *Arithmetic) sub() int {
	return a.x - a.y
}

func (a *Arithmetic) mul() int {
	return a.x * a.y
}

func (a *Arithmetic) div() int {
	if a.y == 0 {
		return 0
	}
	return a.x / a.y
}

func main() {
	fmt.Println("Hello, 世界")

	b := []byte("hello world!!!")
	fmt.Println(string(b))

	a1 := Arithmetic{1, 2}
	a2 := Arithmetic{9, 3}
	a3 := Arithmetic{0, 4}
	a4 := Arithmetic{8, 5}

	aList := []Arithmetic{a1, a2, a3, a4}

	for i, v := range aList {
		fmt.Printf("==== Iter-%d ====\n", i)
		fmt.Printf("Operand 1: %d | Operand 2: %d\n", v.x, v.y)
		fmt.Printf("Operator + : %d\n", v.add())
		fmt.Printf("Operator - : %d\n", v.sub())
		fmt.Printf("Operator * : %d\n", v.mul())
		fmt.Printf("Operator \\ : %d\n", v.div())
		fmt.Println("===================")
	}
}
