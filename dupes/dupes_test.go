package dupes_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-generics/dupes"
)

func TestDupes_IntsHaveDuplicates(t *testing.T) {
	t.Parallel()
	got := dupes.Dupes([]int{1, 2, 1})
	want := true

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDupes_IntsHaveNoDuplicates(t *testing.T) {
	t.Parallel()
	got := dupes.Dupes([]int{1, 2, 3})
	want := false

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDupes_BoolsHaveDuplicates(t *testing.T) {
	t.Parallel()
	got := dupes.Dupes([]bool{true, false, true})
	want := true

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDupes_BoolsHaveNoDuplicates(t *testing.T) {
	t.Parallel()
	got := dupes.Dupes([]bool{true, false})
	want := false

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDupes_StringsHaveDuplicates(t *testing.T) {
	t.Parallel()
	got := dupes.Dupes([]string{"true", "false", "true"})
	want := true

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDupes_StringsHaveNoDuplicates(t *testing.T) {
	t.Parallel()
	got := dupes.Dupes([]string{"true", "false"})
	want := false

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
