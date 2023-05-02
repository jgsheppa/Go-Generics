package stack_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-generics/stack"
)

func TestStack_Push(t *testing.T) {
	t.Parallel()

	s := stack.Stack[int]{
		make(map[int]int),
		[]int{},
	}

	want := 3
	got := s.Push(6, 43, 4)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestStack_Pop(t *testing.T) {
	t.Parallel()

	s := stack.Stack[int]{
		make(map[int]int),
		[]int{},
	}

	want := 2
	s.Push(6, 3, 4)
	got, ok := s.Pop()
	if !ok {
		t.Error(cmp.Diff(want, got))
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestStack_PopFail(t *testing.T) {
	t.Parallel()

	s := stack.Stack[int]{
		make(map[int]int),
		[]int{},
	}

	want := 0
	got, ok := s.Pop()
	if ok {
		t.Error(cmp.Diff(want, got))
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
