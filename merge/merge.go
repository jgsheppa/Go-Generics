package merge

import "golang.org/x/exp/maps"

func Merge[T comparable, V any](s1, s2 map[T]V) map[T]V {
	m := make(map[T]V)
	maps.Copy(m, s1)
	maps.Copy(m, s2)

	return m
}

func BitfieldMerge[M ~map[K]V, K comparable, V any](ms ...M) M {
	result := M{}
	for _, m := range ms {
		maps.Copy[map[K]V](result, m)
	}
	return result
}
