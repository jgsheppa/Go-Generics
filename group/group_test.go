package group_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-generics/group"
)

func TestGroupContainsWhatIsAppendedToIt(t *testing.T) {
	t.Parallel()
	got := group.Group[string]{}
	got = append(got, "hello")
	got = append(got, "world")
	want := group.Group[string]{"hello", "world"}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestLenOfGroupIs2WhenItContains2Elements(t *testing.T) {
	t.Parallel()
	g := group.Group[int]{1, 2}
	want := 2
	got := group.Len(g)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}