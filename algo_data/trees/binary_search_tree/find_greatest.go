package binarysearchtree

func FindGreatest(node *Node, greatestVal int) int {
	if node == nil {
		return greatestVal
	}

	FindGreatest(node.Left, greatestVal)
	if greatestVal < node.Val {
		greatestVal = node.Val
	}
	return FindGreatest(node.Right, greatestVal)
}

func FindGreatestBook(node *Node) int {
	if node.Right != nil {
		return FindGreatestBook(node)
	} else {
		return node.Val
	}
}
