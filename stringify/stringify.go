package stringify

import (
	"fmt"
	"io"
)

func StringifyTo[T fmt.Stringer](writer io.Writer, value T) {
	fmt.Fprintln(writer, value.String())
}
