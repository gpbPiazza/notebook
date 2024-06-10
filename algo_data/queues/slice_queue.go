package queues

type SliceQueue struct {
	store []any
}

func NewSliceQueue() *SliceQueue {
	return &SliceQueue{
		store: make([]any, 0),
	}
}

func (q *SliceQueue) Enqueue(data any) {
	q.store = append(q.store, data)
}

func (q *SliceQueue) Dequeue() any {
	if len(q.store) == 0 {
		return nil
	}

	firstIn := q.store[0]
	q.store = q.store[1:]

	return firstIn
}

func (q *SliceQueue) Read() any {
	if len(q.store) == 0 {
		return nil
	}

	return q.store[0]
}
