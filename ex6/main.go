package main

import (
	"fmt"
	"strconv"
)

func subCodeZZ() {

	code := ""
	i := 1
	for code != "FizzBuzzWhizzBang" {
		if ((i % 3) == 0) && ((i % 5) == 0) && ((i % 7) == 0) && ((i % 11) == 0) {
			code = "FizzBuzzWhizzBang"
		} else if ((i % 5) == 0) && ((i % 11) == 0) {
			code = "BuzzBang"
		} else if ((i % 3) == 0) && ((i % 7) == 0) {
			code = "FizzWhizz"
		} else if ((i % 3) == 0) && ((i % 5) == 0) {
			code = "FizzBuzz"
		} else if (i % 11) == 0 {
			code = "Whizz"
		} else if (i % 7) == 0 {
			code = "Bang"
		} else if (i % 5) == 0 {
			code = "Buzz"
		} else if (i % 3) == 0 {
			code = "Fizz"
		} else {
			code = strconv.Itoa(i)
		}
		fmt.Println(code)
		i++
	}
}

func main() {
	subCodeZZ()
}
