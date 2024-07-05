package queues

type Queue[T any] interface {
	Dequeue() T
	Enqueue(val T)
	Read() T
}
