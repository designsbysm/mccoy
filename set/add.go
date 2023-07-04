package set

func (s *Nodes[T]) Add(item T) {
	if s.items == nil {
		s.Clear()
	}

	if _, ok := s.items[item]; !ok {
		s.items[item] = true
	}
}
