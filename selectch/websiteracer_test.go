package selectch

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//we are creating our own server, rather than trying to send a GET request to actual servers, this is helpful in testing as we don't want to be
//relying on external services.

func TestRacer(t *testing.T) {
	t.Run("compares speed of servers, returning the url of the fastest one", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()
		//prefixing a function call with defer, call the function at the end of the containing function

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server dosen't respond within 10 second", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 20*time.Microsecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

// http.HandlerFunc is a type of function, and we are sending it to httptest.NewServer as a paramneter by creating an anonymous function
// so esentially what a httptest.NewServer function needs, is a function that takes response writer and httprequest
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

//makeDelayedServer returns a pointer to a httptest.Server
