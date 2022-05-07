package hello

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertEqualOutput := func(t testing.TB, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("named hello world", func(t *testing.T) {
		got := Hello("Christ", English)
		want := "Hello, Christ"
		assertEqualOutput(t, got, want)
	})

	t.Run("unnamed hello world", func(t *testing.T) {
		got := Hello("", English)
		want := "Hello, World"

		assertEqualOutput(t, got, want)
	})

	t.Run("named in Spanish", func(t *testing.T) {
		got := Hello("Juan", Spanish)
		want := "Hola, Juan"

		assertEqualOutput(t, got, want)
	})

	t.Run("unnamed in Spanish", func(t *testing.T) {
		got := Hello("", Spanish)
		want := "Hola, Mundo"

		assertEqualOutput(t, got, want)
	})

	t.Run("named in French", func(t *testing.T) {
		got := Hello("Charles", French)
		want := "Bonjour, Charles"

		assertEqualOutput(t, got, want)
	})

	t.Run("unnamed in French", func(t *testing.T) {
		got := Hello("", French)
		want := "Bonjour, le Monde"

		assertEqualOutput(t, got, want)
	})
}
