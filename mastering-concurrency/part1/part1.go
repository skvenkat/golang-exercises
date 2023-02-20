package part1

import (
	"fmt"
	"sync"
)

func Part1() {
	fmt.Println("Part1 - FanOut Pattern in goroutines")
	wg := sync.WaitGroup{}
	//use many goroutines to compute product of the two numbers
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			// capture the value of i and j
			// otherwise, the value of i  and j will be changed
			// before the goroutine is executed
			i := i
			j := j
			// add goroutine to the wait group
			wg.Add(1)
			go func() {
				// defer the decrement of the wait group
				defer wg.Done()
				// compute the product
				product := i * j
				//print the product
				fmt.Println(i, j, product)
			}()
		}
	}
	// ensure that goroutines are executed
	// otherwise, the program will exit
	// before all the goroutines are executed
	wg.Wait()
	fmt.Println("done!")
}
