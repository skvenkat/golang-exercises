package part3

import (
	"fmt"
	"sync"
)


func Part3() {
	fmt.Println("Part3 - Synchronizing critical sections with Mutex")
	// create a map to store the products of two numbers
	// with the numbers in string form as a key
	productMap := map[string]int{}

	// create a mutex to protect the map
	// since only one goroutine can access the map at a time
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	// use many goroutines to compute the product of two numbers
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			copyI, copyJ := i, j
			wg.Add(1)
			go func() {
				defer wg.Done()
				// compute the product
				key := fmt.Sprintf("%d*%d", copyI, copyJ)
				value := copyI * copyJ

				// this is the critical section
				// only one goroutine can access the map at a time
				// this is a common pattern to protect a shared resource
				// with a  mutex
				mu.Lock()
				productMap[key] = value
				mu.Unlock()
			}()
		}
	}

	// wait for all routines to finish
	wg.Wait()

	// print the products from the map
	for key, value := range productMap {
		fmt.Printf("%s = %d\n", key, value)
	}

	fmt.Println("done! ", len(productMap))
}