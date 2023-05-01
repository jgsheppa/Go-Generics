package greater

func IsGreater[T interface {
	Greater(val T) bool
}](x, y T) bool {
	return x.Greater(y)
}
