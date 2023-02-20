package part5

import (
	"fmt"
	"time"
)

// example of long running goroutine
func Part5() {
	fmt.Println("Part5 - Background Tasks with Goroutines")
	quit := make(chan struct{})
	count := 1
	
	// start a goroutine
	// it will run until the channel is closed
	// the quit channel is closed after 5 seconds
	go func() {
		// create a ticker
		ticker := time.NewTicker(
			1* time.Second,
		)

		for {
			// select will wait until one of the channels
			// is ready to be read from
			// if the quit channel is ready to be read from
			// the goroutine will return
			// if the ticker channel is ready to be read from
			// the goroutine will execute the task
			select {

			// read from the ticeker channel
			// the ticker channel will be ready to be read from 
			// every second
			case <- ticker.C:
				// do a task
				fmt.Println("Tick : ", count)
				count++
			case <- quit:
				// stop the goroutine by returning
				fmt.Println("Stopping.....")
				return
			}
		}
	}()

	// wait for 5 seconds
	time.Sleep(5 * time.Second)
	// stop the ticker 
	close(quit)
	fmt.Println("done!")
}
