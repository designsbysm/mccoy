package set

func (s *Nodes[T]) Items() []T {
	var items []T

	for i := range s.items {
		items = append(items, i.(T))
	}

	return items
}
