package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}

func main() {
	sleeper := ConfigurableSleeper{
		duration: 3,
		sleep:    time.Sleep,
	}
	Countdown(os.Stdout, &sleeper)
	fmt.Println()
	Countdown(os.Stdout, SleeperFunc(func() {
		time.Sleep(1 * time.Second)
	}))
}

const (
	countdownStart = 3
	finalWord      = "Go!"
)

type Sleeper interface {
	Sleep()
}

type SleeperFunc func()

func (s SleeperFunc) Sleep() {
	s()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
