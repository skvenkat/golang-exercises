package main

import (
	"fmt"
	"sync"
	"time"
)

// options style function implementation
type Pizza struct {
	pepper         bool
	salt           bool
	cheese         bool
	extraCheese    bool
	chickenTopping bool
	paneerTopping  bool
}

type Options func(config *Pizza)

func BuildPizza(options ...Options) *Pizza {
	p := &Pizza{
		salt:   true,
		pepper: true,
	}

	for _, opt := range options {
		if opt != nil {
			opt(p)
		}
	}

	return p
}

func withPepper() func(config *Pizza) {
	return func(config *Pizza) {
		config.pepper = true
	}
}

func withCheese() func(config *Pizza) {
	return func(config *Pizza) {
		config.cheese = true
	}
}

func withChickenTopping() func(config *Pizza) {
	return func(config *Pizza) {
		config.chickenTopping = true
	}
}

func withExtraCheese() func(config *Pizza) {
	return func(config *Pizza) {
		config.extraCheese = true
	}
}

func orderPizza() {
	myPizza := BuildPizza(withExtraCheese(), withChickenTopping(), withCheese())
	fmt.Printf("my ordered pizza is ===> %+v", myPizza)
}
