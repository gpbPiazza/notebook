package binarysearchtree

import (
	"testing"
)

func TestFindGreatest(t *testing.T) {
	t.Run("greatest val in the last level into most right node", func(t *testing.T) {
		tree := NewBinarySearchTreeTest()

		greatestVal := FindGreatest(tree, 0)
		if greatestVal == 0 {
			t.Error("expected to greatestVal not be 0")
		}

		if greatestVal != 95 {
			t.Errorf("expected to found node 95 as the greatest but got %d", greatestVal)
		}
	})

	t.Run("greatest val is the root node", func(t *testing.T) {
		tree := &Node{
			Val: 100,
			Left: &Node{
				Val:   56,
				Right: &Node{Val: 52},
				Left:  &Node{Val: 61},
			},
		}

		greatestVal := FindGreatest(tree, 0)
		if greatestVal == 0 {
			t.Error("expected to greatestVal not to be zero")
		}
		if greatestVal != 100 {
			t.Errorf("expected to greatestVal to be 100 but got %d", greatestVal)
		}
	})

	t.Run("greatest val is in the second level of the tree", func(t *testing.T) {
		tree := &Node{
			Val: 1,
			Right: &Node{
				Val: 100,
				Left: &Node{
					Val: 80,
					Left: &Node{Val: 60,
						Right: &Node{
							Val: 75},
						Left: &Node{
							Val: 55}}},
			},
		}

		greatestVal := FindGreatest(tree, 0)
		// greatestValBook := FindGreatestBook(tree)
		if greatestVal == 0 {
			t.Error("expected to greatestVal not to be zero")
		}
		if greatestVal != 100 {
			t.Errorf("expected to greatestVal to be 100 but got %d", greatestVal)
		}
	})

	t.Run("randomm tree case", func(t *testing.T) {
		randomSlice := RandomSlice(1000)
		tree := SliceToTree(randomSlice)
		greatestValTree := FindGreatest(tree, 0)
		greatestValSlice := FindGreatestSlice(randomSlice)

		if greatestValSlice != greatestValTree {
			t.Errorf("expected to greatestValTree and greatestValSlice to be equal but got greatestValSlice: %d and greatestValTree: %d",
				greatestValSlice, greatestValTree)
		}
	})
}
