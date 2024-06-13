package binarysearchtree

import "fmt"

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

// Estudar para voltar aqui e implementar essa função por conta própria da meneira correta
// Atual função não deleta o objeto do sucessor da ávore, dessa maneira quebrando a estrutura da árvore.
// caso de test delete node 25
func Delete(node *Node, val int) *Node {
	if node == nil {
		return node
	} else if node.Left.Val == val {
		child, found := FindOneChild(node.Left)
		if found {
			node.Left = child
		} else if node.Left.Left != nil && node.Left.Right != nil {
			sucessor := FindSucessor(node.Left)
			sucessor.Left = node.Left.Left
			sucessor.Right = node.Left.Right
			node.Left = sucessor
		} else {
			node.Left = nil
		}
	} else if node.Right.Val == val {
		child, found := FindOneChild(node.Right)
		if found {
			node.Right = child
		} else if node.Right.Left != nil && node.Right.Right != nil {
			sucessor := FindSucessor(node.Right)
			sucessor.Left = node.Right.Left
			sucessor.Right = node.Right.Right
			node.Right = sucessor
		} else {
			node.Right = nil
		}
	} else if node.Val > val {
		node.Left = Delete(node.Left, val)
		return node
	} else {
		node.Right = Delete(node.Right, val)
		return node
	}
	return node
}

// FindOneChild only returns child node and found true if the Node has Only one children
func FindOneChild(node *Node) (*Node, bool) {
	if node.Left != nil && node.Right == nil {
		return node.Left, true
	}
	if node.Left == nil && node.Right != nil {
		return node.Right, true
	}

	return nil, false
}

func FindSucessor(node *Node) *Node {
	if node == nil || node.Right == nil {
		return node
	}

	node = node.Right
	for node.Left != nil {
		node = node.Left
	}

	return node
}

func DeleteBookImplementation(node *Node, val int) *Node {
	if node == nil {
		return node
	} else if val < node.Val {
		node.Left = DeleteBookImplementation(node.Left, val)
		return node
	} else if val > node.Val {
		node.Right = DeleteBookImplementation(node.Right, val)
		return node
	} else if val == node.Val {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			node.Right = lift(node.Right, node)
			return node
		}
	}
	return node
}

func lift(node, nodeToDelete *Node) *Node {
	if node.Left != nil {
		node.Left = lift(node.Left, nodeToDelete)
		return node
	} else {
		nodeToDelete.Val = node.Val
		return node.Right
	}
}

// Transverse the tree printing nodes in ASC order
func PrintNodesASC(node *Node) {
	if node == nil {
		return
	}

	PrintNodesASC(node.Left)
	fmt.Println(node.Val)
	PrintNodesASC(node.Right)
}
