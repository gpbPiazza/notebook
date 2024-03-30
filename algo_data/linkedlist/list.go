package linkedlist

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
func (l *Node) Add(val int) *Node {
	newNode := &Node{Val: val}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = newNode
	return l
}

func (l *Node) FindMiddleNode() *Node {
	endNode := l
	middleNode := l
	for endNode != nil && endNode.Next != nil {
		endNode = endNode.Next.Next
		middleNode = middleNode.Next
	}

	return middleNode
}

func (l *Node) DeleteMiddleNode() *Node {
	endNode := l
	middleIndex := 0
	for endNode != nil && endNode.Next != nil {
		endNode = endNode.Next.Next
		middleIndex++
	}

	if middleIndex == 0 {
		l.Next = nil
		return l
	}

	preivousIndexFromMiddle := middleIndex - 1

	iterator := l
	index := 0
	for index < preivousIndexFromMiddle {
		iterator = iterator.Next
		index++
	}
	iterator.Next = iterator.Next.Next

	return l
}
