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

// this allows us to use the real sleeper and the mock(spy sleeper), in main and in our tests respectively.
type sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// the configurableSleeper has two fields, one is duration of the the type time.Duration and another is a function. (function as a field)
type SpyTime struct {
	durationSlept time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

//the configurable Sleeper implements the sleeper interface. In this method it calls on the sleep function which is a field in the struct.

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// SpyTime also implemetns the Sleeper interface.
type SpyCountdownOperation struct {
	Calls []string
}

//this struct implements two interfaces i.e Writer and sleeper interface. So we can send this struct, as both parameter, and record for both
//actions.

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
	//over here in the case of sleeper we are passing in the default sleep function. Then the sleeper calls it method, which then calls the sleep function passed in.
}

// countdown function has a dependency on sleep. so we use DI, to use use a mock of time.sleep instead of the real time.sleep.
// thus allowing us to spy on the calls.
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
