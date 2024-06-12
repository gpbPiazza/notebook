package binarysearchtree

// to a tree to be binary search tree, she must satisfy:
// to be binary tree: is a tree that each node has zero, one ou two chieldren
// to be search tree:
// A nodes left decedents can only contain values that are less than the node itself. Likewise, a nodes right
// descedents only contain vakues that are greater than the node itself.
// The lest node must have less value than the father and the rigst ones greater

func Search(node *Node, target int) *Node {
	if node == nil || node.Val == target {
		return node
	} else if node.Val < target {
		return Search(node.Right, target)
	} else {
		return Search(node.Left, target)
	}
}

func Insert(node *Node, val int) {
	if node.Val < val {
		if node.Right == nil {
			node.Right = &Node{Val: val}
			return
		}
		Insert(node.Right, val)
	} else {
		if node.Left == nil {
			node.Left = &Node{Val: val}
			return
		}
		Insert(node.Left, val)
	}
}
