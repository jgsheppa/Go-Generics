package empty_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-generics/empty"
)

func TestSequence_NotEmpty(t *testing.T) {
	t.Parallel()
	want := false
	arr := empty.Sequence[string]{"hello", "world"}
	got := arr.Empty()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSequence_IsEmpty(t *testing.T) {
	t.Parallel()
	want := true
	arr := empty.Sequence[string]{}
	got := arr.Empty()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
