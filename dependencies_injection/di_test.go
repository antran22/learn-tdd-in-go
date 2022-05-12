package dependencies_injection

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	t.Run("test normal greet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		err := Greet(&buffer, "Chris")

		expected := "Hello, Chris"
		actual := buffer.String()

		assert.NoError(t, err, "greet must not throw any error")

		assert.Equal(t, expected, actual, "string mismatched")
	})

}
