package group

type Group[T any] []T

func Len[T any] (v Group[T]) int {
	return len(v)
}