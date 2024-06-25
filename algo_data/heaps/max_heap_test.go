package heaps

import "testing"

func TestMaxHeap(t *testing.T) {
	heap := NewMaxHeap()

	heap.Insert(10)

	lastNode := heap.LastNode()
	if lastNode != 10 {
		t.Errorf("expected to last node to be 10 but got %d", lastNode)
	}

	heap.Insert(200)
	heap.Insert(100)
	heap.Insert(300)

	root := heap.RootNode()
	if root != 300 {
		t.Errorf("expected to root node to be 300 but got %d", root)
	}

	lastNode = heap.LastNode()
	if lastNode != 10 {
		t.Errorf("expected to lastNode to be 10 but got %d", lastNode)
	}

	heap.Insert(250)

	root = heap.RootNode()
	if root != 300 {
		t.Errorf("expected to root node to be 300 but got %d", root)
	}

	lastNode = heap.LastNode()
	if lastNode != 200 {
		t.Errorf("expected to lastNode to be 200 but got %d", lastNode)
	}
	// tree looks like this in this state of test
	//		300
	//		/\
	//	250	100
	//	/ \
	// 10 200

	rootNode := heap.RootNode()
	deletedNode := heap.Delete()
	if deletedNode != rootNode {
		t.Errorf("expected to always delete the root node but got %d", deletedNode)
	}

	rootNode = heap.RootNode()
	if rootNode != 250 {
		t.Errorf("expected to the new rootNode to be 250 but got %d", rootNode)
	}
}
