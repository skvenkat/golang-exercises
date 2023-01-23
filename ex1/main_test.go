package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	var addTestCases = []struct {
		x, y int
		want int
	}{
		{1, 2, 3},
		{9, 3, 12},
		{0, 4, 4},
		{8, 5, 13},
	}

	for _, tt := range addTestCases {
		testHeader := fmt.Sprintf("Operand1: %d | Operand2: %d", tt.x, tt.y)
		t.Run(testHeader, func(t *testing.T) {
			a := Arithmetic{tt.x, tt.y}
			got := a.add()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestSub(t *testing.T) {
	var subTestCases = []struct {
		x, y int
		want int
	}{
		{1, 2, -1},
		{9, 3, 6},
		{0, 4, -4},
		{8, 5, 3},
	}

	for _, tt := range subTestCases {
		testHeader := fmt.Sprintf("Operand1: %d | Operand2: %d", tt.x, tt.y)
		t.Run(testHeader, func(t *testing.T) {
			b := Arithmetic{tt.x, tt.y}
			got := b.sub()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestMul(t *testing.T) {
	var mulTestCases = []struct {
		x, y int
		want int
	}{
		{1, 2, 2},
		{9, 3, 27},
		{0, 4, 0},
		{8, 5, 40},
	}

	for _, tt := range mulTestCases {
		testHeader := fmt.Sprintf("Operand1: %d | Operand2: %d", tt.x, tt.y)
		t.Run(testHeader, func(t *testing.T) {
			c := Arithmetic{tt.x, tt.y}
			got := c.mul()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestDiv(t *testing.T) {
	var divTestCases = []struct {
		x, y int
		want int
	}{
		{1, 2, 0},
		{9, 3, 3},
		{0, 4, 0},
		{8, 5, 1},
	}

	for _, tt := range divTestCases {
		testHeader := fmt.Sprintf("Operand1: %d | Operand2: %d", tt.x, tt.y)
		t.Run(testHeader, func(t *testing.T) {
			d := Arithmetic{tt.x, tt.y}
			got := d.div()
			assert.Equal(t, got, tt.want)
		})
	}
}
