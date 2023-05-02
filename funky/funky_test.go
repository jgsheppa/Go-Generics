package funky_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-generics/funky"
)

func TestFunkyMap_Int(t *testing.T) {
	t.Parallel()

	fn := make(funky.FuncMap[int, int])
	fn["double"] = func(param int) int {
		return param * 2
	}

	want := 2
	got := fn.Apply("double", 1)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFunkyMap_String(t *testing.T) {
	t.Parallel()

	fn := make(funky.FuncMap[string, int])
	fn["length"] = func(param string) int {
		return len(param)
	}

	want := 6
	got := fn.Apply("length", "Vienna")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
