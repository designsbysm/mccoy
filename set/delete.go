package set

func (s *Nodes[T]) Delete(i T) {
	delete(s.items, i)
}
