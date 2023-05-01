package intish

type Intish interface {
	~int
}

func IsPositive[T Intish](num T) bool {
	return num > 0
}
