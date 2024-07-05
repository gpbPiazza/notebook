package queues

type node[T any] struct {
	Val  T
	Prev *node[T]
	Next *node[T]
}

func newNode[T any](val T) *node[T] {
	return &node[T]{Val: val}
}

type doubleLinkedList[T any] struct {
	tail *node[T]
	head *node[T]
}

func newDoubleLinkedList[T any](head, tail *node[T]) *doubleLinkedList[T] {
	return &doubleLinkedList[T]{head: head, tail: tail}
}

func (dll *doubleLinkedList[T]) removeFromFront() T {
	if dll.head == nil {
		return zeroValue[T]()
	}
	removedNode := dll.head
	dll.head = dll.head.Prev

	return removedNode.Val
}

func (dll *doubleLinkedList[T]) insertAtEnd(val T) {
	newNode := newNode(val)
	if dll.tail == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.Next = dll.tail
		dll.tail.Prev = newNode
		dll.tail = newNode
	}
}

type dLLQueue[T any] struct {
	store *doubleLinkedList[T]
}

func newDLLQueue[T any]() *dLLQueue[T] {
	return &dLLQueue[T]{
		store: newDoubleLinkedList[T](nil, nil),
	}
}

func (dllq *dLLQueue[T]) Enqueue(val T) {
	dllq.store.insertAtEnd(val)
}

func (dllq *dLLQueue[T]) Dequeue() T {
	return dllq.store.removeFromFront()
}

func (dllq *dLLQueue[T]) Read() T {
	if dllq.store == nil || dllq.store.head == nil {
		return zeroValue[T]()
	}

	return dllq.store.head.Val
}
