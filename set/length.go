package set

func (s Nodes[T]) Length() int {
	return len(s.items)
}
