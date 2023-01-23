package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	msgQ := make(chan string)
	done := make(chan bool)

	id := 0
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	go func() {
		numJob := rand.Intn(20)
		fmt.Println("num of jobs : ", numJob)
		for i := 0; i < numJob; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				id = i
				//time.Sleep(time.Second * 1)
				ticker := time.NewTicker(time.Second * 1)
				<-ticker.C
				msg := fmt.Sprintf("msg id %d with timestamp : %s", (id + 1), time.Now().String())
				msgQ <- msg
			}()
		}
		wg.Wait()
		close(done)
	}()

	for {
		select {
		case rMsg := <-msgQ:
			fmt.Println("received msg is : ", rMsg)
		case <-done:
			return
		}
	}
}
