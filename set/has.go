package set

func (s *Nodes[T]) Has(item T) bool {
	_, ok := s.items[item]

	return ok
}
