package priorityqueue

func New[T any]() *Queue[T] {
	return &Queue[T]{
		nodes: nil,
	}
}
