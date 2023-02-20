package part4

import (
	"fmt"
	"sync"
)

func Part4() {
	fmt.Println("Part4 - Decoupling Data Transfer With Buffered Channels")
	iMax, jMax := 10, 10
	// create a buffered channel to store the product of two numbers
	// the channel has the exact capability of the number of products
	// that will be written to it

	productChan := make(chan int, iMax*jMax)
	wg := sync.WaitGroup{}
	// use as many goroutines to compute the sum of two numbers
	for i := 0; i < iMax; i++ {
		for j := 0; j< jMax; j++ {
			copyI, copyJ := i, j
			wg.Add(1)
			go func() {
				defer wg.Done()
				// send the product to an unbuffered channel
				productChan <- copyI * copyJ
			}()
		}
	}

	// wait here and allow all the goroutines to finish
	// their work and write to the unbuffered channel
	// this is only possible because the channel is bufferred
	// and has exact capacity of the number of the products
	// that will be written of it
	wg.Wait()
	// close the channel to signal that no more products
	// will be writtent to it
	// the close signal will appended to the end of the channel
	// and the channel will be stay open until all the products
	// are read from the channel
	close(productChan)

	// read the products from the channel 
	// until the channel is closed
	for product := range productChan {
		fmt.Println(product)
	}

	fmt.Println("done!")
}