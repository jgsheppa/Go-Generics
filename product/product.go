package product

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float |
		constraints.Complex
}

func Product[T Number](first, second T) T {
	return first * second
}
