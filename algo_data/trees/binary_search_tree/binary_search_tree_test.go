package binarysearchtree

import (
	"testing"
)

//	     ______(50)______
//			/                \
//		(25)              (75)
//		/   \    	      / 	   \
//	(10)   (33)     (56)     (89)
//	/ \     /  \     /  \     / \
// (4)(11)(30)(40) (52)(61)(82)(95)

func NewBinarySearchTreeTest() *Node {
	root := &Node{
		Val: 50,
		Right: &Node{
			Val: 75,
			Right: &Node{
				Val:   89,
				Right: &Node{Val: 95},
				Left:  &Node{Val: 82},
			},
			Left: &Node{
				Val:   56,
				Right: &Node{Val: 52},
				Left:  &Node{Val: 61},
			},
		},
		Left: &Node{
			Val: 25,
			Right: &Node{
				Val:   33,
				Right: &Node{Val: 40},
				Left:  &Node{Val: 30},
			},
			Left: &Node{
				Val:   10,
				Right: &Node{Val: 11},
				Left:  &Node{Val: 4},
			},
		},
	}

	return root
}

func TestSearch(t *testing.T) {
	tree := NewBinarySearchTreeTest()
	t.Run("return node 4", func(t *testing.T) {
		got := Search(tree, 4)
		if got == nil || got.Val != 4 {
			t.Errorf("expected to found node with val 4 but got %v", got)
		}
	})

	t.Run("return node 95", func(t *testing.T) {
		got := Search(tree, 95)
		if got == nil || got.Val != 95 {
			t.Errorf("expected to found node with val 95 but got %v", got)
		}
	})

	t.Run("return nil when not found 1000", func(t *testing.T) {
		got := Search(tree, 1000)
		if got != nil {
			t.Errorf("expected to not found node 1000 but got %v", got)
		}
	})
}

func TestInsert(t *testing.T) {
	tree := NewBinarySearchTreeTest()
	t.Run("insert node 45", func(t *testing.T) {
		Insert(tree, 45)
		node40 := Search(tree, 40)
		if node40.Right == nil || node40.Right.Val != 45 {
			t.Errorf("expected to found node40 with new value inserted 45 but got %v", node40)
		}
	})
}
