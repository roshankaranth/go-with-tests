package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

//there are two things to test
//1.if the right thing is being tested.
//2.if the write and sleep are in the right order.

func TestCountdown(t *testing.T) {
	t.Run("print 3 to GO!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &SpyCountdownOperation{})

		got := buffer.String()
		want := `3
2
1
GO!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &SpyCountdownOperation{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("wanted calles %v, got %v", want, spySleeperPrinter.Calls)
		}
	})

}

func TestConfigurableSpeaker(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()
	//overhere we are passing in the sleep method of the spytime struct, which is called when sleeper.Sleep() is called.
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
