package liskov_sub

import "fmt"

type human struct {
	name string
}

func (h human) getName() string {
	return h.name
}

type teacher struct {
	human
	degree string
	salary float64
}

type student struct {
	human
	grades map[string]int
}

type person interface {
	getName() string
}

type printer struct{}

func (pr printer) info(p person) {
	fmt.Println("Name: ", p.getName())
}

func LiskovSub() {
	h := human{"John"}
	t := teacher{human{"Marie"}, "BA", 3500.56}
	s := student{human{"Peter"}, map[string]int{"a": 23, "b": 34, "c": 32}}
	p := printer{}
	p.info(h)
	p.info(t)
	p.info(s)

}
