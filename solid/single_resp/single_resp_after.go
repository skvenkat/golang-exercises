package single_resp

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) name() string {
	return "Circle"
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s square) name() string {
	return "Square"
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
	name() string
}

type outputter struct{}

func (o outputter) Text(s shape) string {
	return fmt.Sprintf("area of the %s : %f", s.name(), s.area())
}

func (o outputter) Json(s shape) string {
	res := struct {
		Name string  `json:"shape"`
		Area float64 `json:"area"`
	}{
		Name: s.name(),
		Area: s.area(),
	}

	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func SingleResp() {
	c := circle{3}
	s := square{6}
	out := outputter{}
	fmt.Println(out.Text(c))
	fmt.Println(out.Json(c))
	fmt.Println(out.Text(s))
	fmt.Println(out.Json(s))
}
