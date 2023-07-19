package priorityqueue

func New[T any]() *Queue[T] {
	return &Queue[T]{
		length: 0,
		nodes:  nil,
	}
}
