package linkedlist

import "fmt"

type DLLNode struct {
	Val  any
	Prev *DLLNode
	Next *DLLNode
}

func NewDLLNode(val any) *DLLNode {
	return &DLLNode{Val: val}
}

type DoubleLinkedList struct {
	tail *DLLNode
	head *DLLNode
}

func NewDoubleLinkedList(head, tail *DLLNode) *DoubleLinkedList {
	head.Prev = tail
	tail.Next = head
	return &DoubleLinkedList{head: head, tail: tail}
}

func (dll *DoubleLinkedList) InsertAtHead(val any) {
	newNode := NewDLLNode(val)
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.Prev = dll.head
		dll.head.Next = newNode
		dll.head = newNode
	}
}

func (dll *DoubleLinkedList) InsertAtTail() {}

func (dll *DoubleLinkedList) Print() {
	if dll.tail == nil {
		fmt.Print("empty list")
	}

	currentNode := dll.tail
	for currentNode != nil {
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
	}
}
func (dll *DoubleLinkedList) ReversePrint() {
	if dll.head == nil {
		fmt.Print("empty list")
	}

	currentNode := dll.head
	for currentNode != nil {
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Prev
	}
}
