package binarysearchtree

import (
	"testing"

	"github.com/gpbPiazza/notebook/algo_data/stdout"
	"github.com/stretchr/testify/assert"
)

//	     ______(50)______
//			/                \
//		(25)              (75)
//		/   \    	      / 	   \
//	(10)   (33)     (56)     (89)
//	/ \     /  \     /  \     / \
//
// (4)(11)(30)(40) (52)(61)(82)(95)
//	     \
//	     (12)

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
				Right: &Node{Val: 11, Right: &Node{Val: 12}},
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
	t.Run("insert node 45", func(t *testing.T) {
		tree := NewBinarySearchTreeTest()

		Insert(tree, 45)
		node40 := Search(tree, 40)
		if node40.Right == nil || node40.Right.Val != 45 {
			t.Errorf("expected to found node40 with new value inserted 45 but got %v", node40)
		}
	})
}

// func Test_My_Delete_Implementation(t *testing.T) {
// 	DeletTestCases(t, Delete)
// }

func Test_Book_Delete_Implementation(t *testing.T) {
	DeletTestCases(t, DeleteBookImplementation)
}

func DeletTestCases(t *testing.T, DeleteFunc func(node *Node, val int) *Node) {
	t.Run("delete node 4 who has no children", func(t *testing.T) {
		tree := NewBinarySearchTreeTest()

		DeleteFunc(tree, 4)
		node4 := Search(tree, 4)
		if node4 != nil {
			t.Errorf("expected to not found node4 but got %v", node4)
			return
		}

		node10 := Search(tree, 10)
		if node10.Left != nil {
			t.Errorf("expected to node 10 with left child to be nil but got %v", node10.Left)
			return
		}
	})

	t.Run("delete node 11 who has 1 child", func(t *testing.T) {
		tree := NewBinarySearchTreeTest()

		DeleteFunc(tree, 11)

		node11 := Search(tree, 11)
		if node11 != nil {
			t.Errorf("expected to not found node11 but got %v", node11)
			return
		}

		node10 := Search(tree, 10)
		if node10.Right == nil || node10.Right.Val != 12 {
			t.Errorf("expected to node 10 with right child to be 12 but got %v", node10.Right)
			return
		}
	})

	t.Run("delete node 25 who has 2 child", func(t *testing.T) {
		tree := NewBinarySearchTreeTest()

		DeleteFunc(tree, 25)

		node25 := Search(tree, 25)
		if node25 != nil {
			t.Errorf("expected to not found node25 but got %v", node25)
			return
		}

		node30 := Search(tree, 30)
		if node30 == nil {
			t.Error("expected to node 30 to be the successor of 25 but got nil")
			return
		}

		if node30.Left == nil || node30.Left.Val != 10 {
			t.Errorf("expected to successor node has left child to be node10 but got %v", node30.Left)
			return
		}

		if node30.Right == nil || node30.Right.Val != 33 {
			t.Errorf("expected to successor node has right child to be node33 but got %v", node30.Right)
			return
		}

		node33 := Search(tree, 33)
		if node33.Left != nil {
			t.Errorf("expected to node33 has left child to be nil  but got %v", node33.Left)
			return
		}
	})
}

func TestPrintNodes(t *testing.T) {
	t.Run("test Inorder transverse", func(t *testing.T) {
		output := stdout.String(func() { Inorder_TransverseAndPrint(NewBinarySearchTreeTest()) })
		assert.Equal(t, "4\n10\n11\n12\n25\n30\n33\n40\n50\n61\n56\n52\n75\n82\n89\n95\n", output)
	})

	t.Run("test Preorder transverse", func(t *testing.T) {
		output := stdout.String(func() { Preorder_TrasnverseAndPrint(NewBinarySearchTreeTest()) })
		assert.Equal(t, "50\n25\n10\n4\n11\n12\n33\n30\n40\n75\n56\n61\n52\n89\n82\n95\n", output)
	})
}
