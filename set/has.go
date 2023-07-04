package set

func (s *Nodes[T]) Has(i T) bool {
	_, ok := s.items[i]

	return ok
}
