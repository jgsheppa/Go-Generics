package dupes

func equals[T comparable](x, y T) bool {
	return x == y
}

func Dupes[T comparable](arr []T) bool {
	for _, i := range arr {
		for _, j := range arr {
			if equals(i, j) {
				return true
			}
			continue
		}
	}

	return false
}
