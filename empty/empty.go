package empty

type Sequence[T any] []T

func (s Sequence[T]) Empty() bool {
	return len(s) == 0
}
