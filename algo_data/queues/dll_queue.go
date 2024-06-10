package queues

type Node struct {
	Val  any
	Prev *Node
	Next *Node
}

func NewNode(val any) *Node {
	return &Node{Val: val}
}

type DoubleLinkedList struct {
	tail *Node
	head *Node
}

func NewDoubleLinkedList(head, tail *Node) *DoubleLinkedList {
	return &DoubleLinkedList{head: head, tail: tail}
}

func (dll *DoubleLinkedList) removeFromFront() any {
	if dll.head == nil {
		return nil
	}
	removedNode := dll.head
	dll.head = dll.head.Prev

	return removedNode.Val
}

func (dll *DoubleLinkedList) insertAtEnd(val any) {
	newNode := NewNode(val)
	if dll.tail == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.Next = dll.tail
		dll.tail.Prev = newNode
		dll.tail = newNode
	}
}

type DLLQueue struct {
	store *DoubleLinkedList
}

func NewDLLQueue() *DLLQueue {
	return &DLLQueue{
		store: NewDoubleLinkedList(nil, nil),
	}
}

func (dllq *DLLQueue) Enqueue(val any) {
	dllq.store.insertAtEnd(val)
}

func (dllq *DLLQueue) Dequeue() any {
	return dllq.store.removeFromFront()
}

func (dllq *DLLQueue) Read() any {
	if dllq.store == nil || dllq.store.head == nil {
		return nil
	}

	return dllq.store.head.Val
}
