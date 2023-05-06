package contains

import "golang.org/x/exp/slices"

func ContainsFunc[T any](s []T, f func(r T) bool) bool {
	return slices.ContainsFunc(s, f)

}
