package linkedlist

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func New(val int) *Node {
	return &Node{
		Val: val,
	}
}

// SliceToList uses the inputed slice to build a new linkedList
// SliceToList returns the head of the linkedList
func SliceToList(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}

	head := &Node{
		Val: vals[0],
	}
	for i := 1; i < len(vals); i++ {
		head.Add(vals[i])
	}
	return head
}

// Add concat new Node to the last node in the list and return the head of the list
func (n *Node) Add(val int) *Node {
	newNode := &Node{Val: val}
	iterator := n
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = newNode
	return n
}

func (n *Node) FindMiddleNode() *Node {
	endNode := n
	middleNode := n
	for endNode != nil && endNode.Next != nil {
		endNode = endNode.Next.Next
		middleNode = middleNode.Next
	}

	return middleNode
}

func (n *Node) DeleteMiddleNode() *Node {
	endNode := n
	middleIndex := 0
	for endNode != nil && endNode.Next != nil {
		endNode = endNode.Next.Next
		middleIndex++
	}

	if middleIndex == 0 {
		n.Next = nil
		return n
	}

	preivousIndexFromMiddle := middleIndex - 1

	iterator := n
	index := 0
	for index < preivousIndexFromMiddle {
		iterator = iterator.Next
		index++
	}
	iterator.Next = iterator.Next.Next

	return n
}

func (n *Node) SortByOddIndexes() *Node {
	if n.Next == nil {
		return n
	}

	odd := n
	oddHead := n
	even := n.Next
	evenHead := n.Next

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = even.Next

		even.Next = even.Next.Next
		even = odd.Next
	}
	odd.Next = evenHead

	return oddHead
}

func (n *Node) Print() {
	if n == nil {
		fmt.Println("empty list")
	}

	currentNode := n
	for currentNode != nil {
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
	}
}

func (n *Node) LastNode() *Node {
	if n == nil {
		return nil
	}

	currentNode := n
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	return currentNode
}

func (n *Node) Reverse() *Node {
	if n == nil || n.Next == nil {
		return n
	}

	current := n
	var prev *Node
	for current.Next != nil {
		nextNode := current.Next

		current.Next = prev

		prev = current
		current = nextNode
	}

	current.Next = prev

	return current
}

// selfDelete deletes the called node to be deleted from the linked list
// it will do nothing if N or n.Next is nil
func (n *Node) SelfDelete() {
	if n == nil || n.Next == nil {
		return
	}

	newVal := n.Next.Val

	n.Val = newVal
	n.Next = n.Next.Next
}
