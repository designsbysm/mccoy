package set

func (s *Nodes[T]) Add(i T) {
	if s.items == nil {
		s.Clear()
	}

	if _, ok := s.items[i]; !ok {
		s.items[i] = true
	}
}
