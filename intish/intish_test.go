package intish_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-generics/intish"
)

type MyInt int

func TestIsPositive_IsTrueFor1(t *testing.T) {
	t.Parallel()
	got := intish.IsPositive[MyInt](1)

	want := true
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsPositive_IsFalseForNegative1(t *testing.T) {
	t.Parallel()
	got := intish.IsPositive[MyInt](-1)

	want := false
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsPositive_IsFalseForZero(t *testing.T) {
	t.Parallel()
	got := intish.IsPositive[MyInt](0)

	want := false
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
