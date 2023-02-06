package main

import "fmt"

const LRU_SIZE = 10

var LRU = make([]string, LRU_SIZE)

func lookupCurrentIndex(item string) int {
	for i := (len(LRU) - 1); i > 0; i-- {
		if LRU[i] == item {
			return i
		}
	}
	return -1
}

func updateCache(item string) {
	curIdx := lookupCurrentIndex(item)
	resetIdx := -2
	if curIdx > -1 {
		resetIdx = curIdx
	} else if len(LRU) >= LRU_SIZE {
		resetIdx = 0
	} else {
		LRU = append(LRU, item)
	}

	for i := resetIdx; i < len(LRU); i++ {
		if i+1 < len(LRU) {
			LRU[i] = LRU[i+1]
		}
	}
}

func main() {
	fmt.Println("Hello, 世界")

	a := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "k", "l"}

	for i, v := range a {
		updateCache(v)
		fmt.Printf("Iter %d : %v\n", i, LRU)
	}

}
