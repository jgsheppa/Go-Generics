package funky

type FuncMap[E, F any] map[string]func(E) F

func (f FuncMap[E, F]) Apply(fn string, param E) F {
	return f[fn](param)
}
