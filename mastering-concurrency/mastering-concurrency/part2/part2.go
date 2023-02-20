package part2

import (
	"fmt"
	"sync"
)

func Part2() {
	fmt.Println("Part1 - FanIn Pattern with goroutines")
	productChan := make(chan int)
	wg := sync.WaitGroup{}
	//use many goroutines to compute the product of two numbers
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			copyI, copyJ := i, j
			wg.Add(1)
			go func() {
				defer wg.Done()
				// send the product to an unbuffered channel
				// this is a common patter to in results from many go routines
				productChan <- copyI * copyJ
			}()
		}
	}

	// start a goroutine to wait for all the goroutines 
	// that are doing work to finish
	// then close the channel to signal the main() with
	// all the work is done
	go func() {
		wg.Wait()
		// closing the channel will cause the range loop to exit
		// this is a common pattern to signal the end of a channel
		close(productChan)
	}()

	// read the products from the channel
	// until the channel is closed
	for product := range productChan {
		fmt.Println(product)
	}

	fmt.Println("done!")
}