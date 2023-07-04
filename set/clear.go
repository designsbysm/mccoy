package set

func (s *Nodes[T]) Clear() {
	s.items = make(map[any]bool)
}
