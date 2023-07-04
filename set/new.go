package set

func New[T any]() Nodes[T] {
	set := Nodes[T]{}
	set.Clear()

	return set
}
