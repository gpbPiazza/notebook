package queues

type Queue[T any] struct {
	store []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		store: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(data T) {
	q.store = append(q.store, data)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.store) == 0 {
		return q.zeroValue()
	}

	firstIn := q.store[0]
	q.store = q.store[1:]

	return firstIn
}

func (q *Queue[T]) Read() T {
	if len(q.store) == 0 {
		return q.zeroValue()
	}

	return q.store[0]
}

func (s *Queue[T]) zeroValue() T {
	var zero T
	return zero
}
