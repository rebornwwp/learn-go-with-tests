package mocking

import (
	"fmt"
	"io"
	"time"
)

func Countdown(out io.Writer, sleep Sleeper) {
	for i := 3; i > 0; i-- {
		sleep.Sleep()
		fmt.Fprintln(out, i)
	}
	sleep.Sleep()
	fmt.Fprint(out, "Go!")
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}
