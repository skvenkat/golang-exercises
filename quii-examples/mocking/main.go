package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper allows you to put delays
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper is an implementation of Sleeper with defined delay
type ConfigurableSleeper struct {
	duration 	time.Duration
	sleep		func(time.Duration)
}

// Sleep will pause execution for the defined duration
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from 3 to out with a delay between count provided by Sleeper
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(out, "%d", i)
		sleeper.Sleep()
	}

	fmt.Fprintf(out, "%s", finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
