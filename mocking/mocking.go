package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration  time.Duration
	sleepFunc func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleepFunc(s.duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)
}

func main() {
	sleeper := ConfigurableSleeper{
		duration:  1 * time.Second,
		sleepFunc: time.Sleep,
	}
	Countdown(os.Stdout, &sleeper)
}
