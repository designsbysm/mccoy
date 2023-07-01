package linkedlist

func (n *T) Delete(index int) {
	// TODO: change to At
	found := n.Index(index - 1)
	next := found.tail

	if next != nil {
		found.tail = next.tail
	}
}
