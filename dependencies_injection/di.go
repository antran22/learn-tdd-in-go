package dependencies_injection

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, name string) error {
	bytesWritten, err := fmt.Fprintf(w, "Hello, %s", name)
	if err != nil {
		return fmt.Errorf("%w: unable to greet", err)
	}
	if bytesWritten <= 0 {
		return fmt.Errorf("unable to greet")
	}
	return nil
}
