package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcStats(t *testing.T) {
	var testCases = []struct {
		arr  []int
		want Stats
	}{
		{[]int{6, 9, 15, -2, 92, 11},
			Stats{-2, 92, 6, 21.833334},
		},
		{[]int{164, 0, 167, -32, -49, 121, 4, 195, 59, 108, 169, 55, 30, 15, 20, 99, 46, 191, 39, 175, -19, 12, 2, 134, 66, 144, 103},
			Stats{-49, 195, 27, 74.740740},
		},
		{[]int{197, 136, 106, 102, 33, -36, 116, 97, 127, -31, -11, 119, -33, 40, 92, 163, 91, 100, -47, 123, 196, 94, -34, 177, -30, 82, 114, 90, -32, -3, 23, 7, 143, 88, 84, 117, 77, 69, 164, 130, 113, -23, 169, 140, 120, 172, 171, 81, 13, 26},
			Stats{-47, 197, 50, 80.440000},
		},
		{[]int{61, -17, 119, 185, 57, 83, 26, 92, 56, 112, 146, 36, 128, -35, 186, 105, 188},
			Stats{-35, 188, 17, 89.882352},
		},
		{[]int{176, -2, 98, -8, 36, -4, 27, 29, 59, 10, 52, -40, 154, 54, 140, 26, -6, 7, 94, 2, 174, 142, 138},
			Stats{-40, 176, 23, 59.04348},
		},
	}

	for i, tc := range testCases {
		testHeader := fmt.Sprintf("=====> Test case : %d <=====", i)
		t.Run(testHeader, func(t *testing.T) {
			tmpStats := calcStats(tc.arr)
			assert.Equal(t, tmpStats, tc.want)
		})
	}
}

func BenchCalcStats(b *testing.B) {
	arr := []int{94, 146, 64, 177, 81, 29, -45, 10, 43, -17, -47, 56, 98, 93, 124, 96, 41, -7, 15, 61, 192, 199, 190, 87, 180, 6, 83, 143, 156, 166, 110, 66, 48, 171, 115, 149, -36, 122, 80, 21, 198, 153, 131, -35, -31, -39, 9, 121, 63, 20}
	for i := 1; i <= 10; i++ {
		_ = calcStats(arr)
	}
}