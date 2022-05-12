package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep(duration time.Duration)
}

type DefaultSleep struct{}

func (d *DefaultSleep) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		sleeper.Sleep(time.Second)
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep(1 * time.Second)
	fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdout, &DefaultSleep{})
}
