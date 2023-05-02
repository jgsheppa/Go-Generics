package stack

type Stack[E comparable] struct {
	SubStack map[E]E
	Order    []E
}

func (s *Stack[E]) Push(vals ...E) int {
	for _, val := range vals {
		s.SubStack[val] = val
		s.Order = append(s.Order, val)
	}
	return len(s.SubStack)
}

func (s *Stack[E]) Pop() (int, bool) {
	if s.Length() == 0 {
		return 0, false
	}
	last := s.Order[len(s.Order)-1]
	ok := s.Contains(last)
	if ok {
		delete(s.SubStack, last)
		s.Order = append(s.Order[:len(s.Order)-1], s.Order[len(s.Order):]...)
	}

	return len(s.SubStack), ok
}

func (s *Stack[E]) Contains(val E) bool {
	_, ok := s.SubStack[val]
	return ok
}

func (s *Stack[E]) Length() int {
	return len(s.SubStack)
}
