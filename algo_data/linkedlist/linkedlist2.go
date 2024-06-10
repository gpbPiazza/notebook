package linkedlist

type Node2 struct {
	Val  any
	Next *Node2
}

func NewNode2(val any) *Node2 {
	return &Node2{
		Val: val,
	}
}

type LinkedList2 struct {
	Head *Node2
}

func NewLinkedList2(head *Node2) *LinkedList2 {
	return &LinkedList2{
		Head: head,
	}
}

func (ll *LinkedList2) Read(index int) any {
	currentIndex := 0
	currentNode := ll.Head
	for currentIndex < index {
		currentIndex++
		currentNode = currentNode.Next
	}

	if currentNode == nil {
		return nil
	}

	return currentNode.Val
}

func (ll *LinkedList2) Search(val any) any {
	if ll.Head.Val == val {
		return ll.Head.Val
	}

	currentNode := ll.Head

	for currentNode != nil {
		if currentNode.Val == val {
			return currentNode.Val
		}

		currentNode = currentNode.Next
	}

	return nil
}

func (ll *LinkedList2) InserAtIndex(index int, val any) {
	newNode := NewNode2(val)
	if index == 0 {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	currentIndex := 0
	currentNode := ll.Head

	for currentIndex < (index - 1) {
		currentIndex++
		currentNode = currentNode.Next
	}

	newNode.Next = currentNode.Next
	currentNode.Next = newNode
}
