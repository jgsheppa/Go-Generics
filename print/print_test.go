package print_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"

	print "github.com/jgsheppa/go-generics/print"
)

func TestPrintAnythingTo_PrintsInputToSuppliedWriter(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	print.PrintAnythingTo(buf, "Hello, world")
	want := "Hello, world\n"
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}