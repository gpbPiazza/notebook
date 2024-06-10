package queues

type Queue interface {
	Dequeue() any
	Enqueue(val any)
	Read() any
}
