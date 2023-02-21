package main

import (
	"fmt"
	"sync"
)

type Limiter struct {
	pool chan struct{}
	wg   sync.WaitGroup
}

func NewLimiter(n int) *Limiter {
	return &Limiter{pool: make(chan struct{}, n), wg: sync.WaitGroup{}}
}

func (l *Limiter) Run(task func()) {
	l.wg.Add(1)
	go func() {
		l.pool <- struct{}{}
		task()
		<-l.pool
		l.wg.Done()
	}()
}

func (l *Limiter) Wait() {
	fmt.Println("Waiting to get some free space in the pool...")
	l.wg.Wait()
}

func main() {
	limiter := NewLimiter(10)
	for i := 0; i < 1_00; i++ {
		i := i
		limiter.Run(func() {
			fmt.Println("Hello ", i)
		})
	}

	limiter.Wait()
	fmt.Println("done!")
}