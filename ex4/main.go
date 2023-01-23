package main

import "fmt"

type Human struct {
	name   string
	gender string
	age    int
}

func (h Human) Info() {
	fmt.Printf("i am %s and i am a %s and i am %d yr old...\n", h.name, h.gender, h.age)
}

type HomoSapien interface {
	Info()
}

func main() {
	human := Human{"peppa", "female", 10}
	homosapien := Human{"george", "male", 8}

	var h HomoSapien
	h = human
	h.Info()
	h = homosapien
	h.Info()
}
