package contains_test

import (
	"testing"
	"unicode"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-generics/contains"
)

func TestContainsFunc_Pass(t *testing.T) {
	t.Parallel()

	got := contains.ContainsFunc([]rune{'a', 'B', 'c'}, unicode.IsUpper)
	want := true

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestContainsFunc_Fail(t *testing.T) {
	t.Parallel()

	got := contains.ContainsFunc([]rune{'a', 'b', 'c'}, unicode.IsUpper)
	want := false

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
