package print

import (
	"fmt"
	"io"
)

func PrintAnythingTo[T any](w io.Writer, v T) {
	fmt.Fprintln(w, v)
}