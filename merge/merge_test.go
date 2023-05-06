package merge_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-generics/merge"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	s1 := map[string]bool{
		"one": true,
	}

	s2 := map[string]bool{
		"two": true,
	}

	got := merge.Merge(s1, s2)
	want := map[string]bool{
		"one": true,
		"two": true,
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestMergeOverwrite(t *testing.T) {
	t.Parallel()

	s1 := map[string]bool{
		"one": true,
	}

	s2 := map[string]bool{
		"one": false,
		"two": true,
	}

	got := merge.Merge(s1, s2)
	want := map[string]bool{
		"one": false,
		"two": true,
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
