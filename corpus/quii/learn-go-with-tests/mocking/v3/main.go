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

// ConfigurableSleeper is an implementation of Sleeper with a defined delay
type ConfigurableSleeper struct {
	Duration time.Duration
}

// Sleep will pause execution for the defined Duration
func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.Duration)
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from 5 to out with a delay between count provided by Sleeper
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
