package dependencyinjection

import (
	"fmt"
	"io"
)

func Greeting(w io.Writer, name string) {
	fmt.Fprintf(w, "hello, %s", name)
}
