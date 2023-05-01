package dupes

func equals[T comparable](x, y T) bool {
	return x == y
}

func Dupes[T comparable](arr []T) bool {
	seen := map[T]bool{}
	for _, v := range arr {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}
