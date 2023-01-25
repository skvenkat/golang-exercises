package main

import "fmt"

type Stats struct {
	min   int
	max   int
	count int
	avg   float32
}

func calcStats(arr []int) Stats {
	stats := Stats{count: len(arr)}

	for _, val := range arr {
		if val < stats.min {
			stats.min = val
		}

		if val > stats.max {
			stats.max = val
		}

		stats.avg += float32(val)
	}

	stats.avg = (stats.avg / float32(stats.count))

	return stats
}

func main() {
	arr := []int{6, 9, 15, -2, 92, 11}
	newStats := calcStats(arr)
	fmt.Printf("Calculated stats : %+v\n", newStats)
}
