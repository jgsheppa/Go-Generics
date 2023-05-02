package compose_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-generics/compose"
)

func double(val int) int {
	return val * 2
}

func addOne(val int) int {
	return val + 1
}

func TestCompose_DoubleThenAddOne(t *testing.T) {
	t.Parallel()
	want := 3
	got := compose.Compose(double, addOne, 1)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCompose_AddOneThenDouble(t *testing.T) {
	t.Parallel()
	want := 4
	got := compose.Compose(addOne, double, 1)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
