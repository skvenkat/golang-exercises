package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckLeapYear(t *testing.T) {
	var testCases = []struct {
		year   int
		isLeap bool
	}{
		{1897, false},
		{1900, false},
		{1927, false},
		{1947, false},
		{1965, false},
		{1977, false},
		{1981, false},
		{1990, false},
		{1996, true},
		{2000, true},
		{2006, false},
		{2010, false},
		{2014, false},
		{2016, true},
		{2020, true},
	}

	for _, tt := range testCases {
		testHeader := fmt.Sprintf("====> Year : %d <====", tt.year)
		t.Run(testHeader, func(t *testing.T) {
			assert.Equal(t, checkLeapYear(tt.year), tt.isLeap)
		})
	}

}
