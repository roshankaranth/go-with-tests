package helloworld

import "testing"

// grouping subtests.
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("", "Roshan")
		want := "Hello, Roshan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say, 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Spanish", "Roshan")
		want := "Hola, Roshan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("French", "Roshan")
		want := "Bonjour, Roshan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Hindi", func(t *testing.T) {
		got := Hello("Hindi", "Roshan")
		want := "Namaste, Roshan"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	//tells the compiler it is a helper function, so if a error occurs in this function, in the stack trace of error message it points to the line
	//which calls this function. Otherwise it would have pointed inside this function.
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
