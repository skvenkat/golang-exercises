package main

import "testing"

func BenchmarkMain(b *testing.B) {
	for i := 0; i < 10; i++ {
		main()
	}
}
