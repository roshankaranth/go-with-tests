package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "GO!"
const countdownStart = 3
const write = "write"
const sleep = "sleep"

type sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

type SpyCountdownOperation struct {
	Calls []string
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *SpyCountdownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperation) Write(p []byte) (n int, e error) {
	s.Calls = append(s.Calls, write)
	return
}

//spies are kind of mocks which can record how dependecy is used. They can record the argument
//sent in, how many times it has been called.

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func Countdown(w io.Writer, Sleeper sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		Sleeper.Sleep()
	}

	// for i := countdownStart; i > 0; i-- {
	// 	Sleeper.Sleep()
	// }

	// for i := countdownStart; i > 0; i-- {
	// 	fmt.Fprintln(w, i)
	// }

	//wrong implementation

	fmt.Fprint(w, finalWord)
}

//io.Writer interface is the defacto way of writing data somewhere.
