package queues

type SliceQueue[T any] struct {
	store []T
}

func NewSliceQueue[T any]() *SliceQueue[T] {
	return &SliceQueue[T]{
		store: make([]T, 0),
	}
}

func (q *SliceQueue[T]) Enqueue(data T) {
	q.store = append(q.store, data)
}

func (q *SliceQueue[T]) Dequeue() T {
	if len(q.store) == 0 {
		return zeroValue[T]()
	}

	firstIn := q.store[0]
	q.store = q.store[1:]

	return firstIn
}

func zeroValue[T any]() T {
	var zero T
	return zero
}

func (q *SliceQueue[T]) Read() T {
	if len(q.store) == 0 {
		return zeroValue[T]()
	}

	return q.store[0]
}
