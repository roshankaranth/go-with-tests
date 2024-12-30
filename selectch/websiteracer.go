package selectch

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

//the user calls the racer function, wheras we use the configurablRacer function for testing purposes.

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	//myVar := <- ch, this is a blocking call, as we are waiting for a value to be sent to the channel.
	//select allows waiting for multiple channels, and whichever  sends the value first wins and the code underneath the case is executed
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

	//select esentially listens on multiple channel (block) until it recieves from a channel, and when it does, it exectues the corresponding
	//case

	//time.After() returns a chan
}

func ping(url string) chan struct{} {
	//we create a channel, where we do not care what type we are sending to the channel, we just want to signal we are done and closing the channel works perfectly
	//struct{} type is the smallest data type available from memoery perspective(we are not allocating anything) as opposed to boolean.
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

//we use make to create channels instead of var ch chan struct{} because when using var the varaible is initialized to default value and in case of channels
//it will be initialized to nil values. and if we try to send to nil channels it will block forever, because we can't send to nil channels
