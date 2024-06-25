package heaps

// MaxHeap is a binary tree with array based implementation
// Where MaxHeap must satisfy the conditions:
// 1. The value of each node must be greater than each of his decendents (heap condition)
// 2. The tree must be complete. (We should I always insert at the right most spot,
// if there is any spot that is not the right most tree is no completed)
type MaxHeap struct {
	store []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		store: make([]int, 0),
	}
}

func (mh *MaxHeap) RootNode() int {
	return mh.store[0]
}

func (mh *MaxHeap) LastNode() int {
	return mh.store[len(mh.store)-1]
}

func (mh *MaxHeap) Insert(newVal int) {
	mh.store = append(mh.store, newVal)
	newNodeIndex := len(mh.store) - 1

	// newNodeIndex > 0 to avoind index out of range in parentIndex(0-1/2) = -1
	for newNodeIndex > 0 && mh.store[mh.parentIndex(newNodeIndex)] < mh.store[newNodeIndex] {
		mh.store[newNodeIndex], mh.store[mh.parentIndex(newNodeIndex)] =
			mh.store[mh.parentIndex(newNodeIndex)], mh.store[newNodeIndex]

		newNodeIndex = mh.parentIndex(newNodeIndex)
	}
}

func (mh *MaxHeap) Delete() int {
	deletedNode := mh.RootNode()
	mh.store[0] = mh.LastNode()

	trickledNodeIndex := 0
	for mh.hasGreaterChildren(trickledNodeIndex) {
		mh.store[mh.greaterChildrenIndex(trickledNodeIndex)], mh.store[trickledNodeIndex] =
			mh.store[trickledNodeIndex], mh.store[mh.greaterChildrenIndex(trickledNodeIndex)]

		trickledNodeIndex = mh.greaterChildrenIndex(trickledNodeIndex)
	}

	return deletedNode
}

func (mh *MaxHeap) hasGreaterChildren(index int) bool {
	rightChild := mh.store[mh.rightChildIndex(index)]
	leftChild := mh.store[mh.leftChildIndex(index)]
	return rightChild > mh.store[index] || leftChild > mh.store[index]
}

func (mh *MaxHeap) greaterChildrenIndex(index int) int {
	rightChild := mh.store[mh.rightChildIndex(index)]
	leftChild := mh.store[mh.leftChildIndex(index)]

	if rightChild > leftChild {
		return mh.rightChildIndex(index)
	} else {
		return mh.leftChildIndex(index)
	}
}

func (mh *MaxHeap) rightChildIndex(index int) int {
	return (index * 2) + 2
}

func (mh *MaxHeap) leftChildIndex(index int) int {
	return (index * 2) + 1
}

func (mh *MaxHeap) parentIndex(index int) int {
	return (index - 1) / 2
}
