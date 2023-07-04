package set

func (s *Nodes[T]) Delete(item T) {
	delete(s.items, item)
}
