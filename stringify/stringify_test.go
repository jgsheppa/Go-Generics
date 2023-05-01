package stringify_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jgsheppa/go-generics/stringify"
)

type greeting struct{}

func (greeting) String() string {
	return "Howdy!"
}

func TestStringifyTo(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	stringify.StringifyTo(buf, greeting{})
	want := "Howdy!\n"
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
