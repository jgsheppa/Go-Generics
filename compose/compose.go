package compose

func Compose[T, U, V any](fn1 func(T) V, fn2 func(V) U, val T) U {
	return fn2(fn1(val))
}
