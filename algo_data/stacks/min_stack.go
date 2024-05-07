package stacks

// https://leetcode.com/problems/min-stack/description/
type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinStack() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (ms *MinStack) Push(val int) {
	if len(ms.minStack) != 0 {
		minVal := ms.minStack[ms.minStackLastIndex()]
		if val <= minVal {
			ms.minStack = append(ms.minStack, val)
		}
	} else {
		ms.minStack = append(ms.minStack, val)
	}

	ms.stack = append(ms.stack, val)
}

func (ms *MinStack) Pop() {
	minVal := ms.minStack[ms.minStackLastIndex()]
	val := ms.stack[ms.stackLastIndex()]

	if minVal == val {
		ms.minStack = ms.minStack[:ms.minStackLastIndex()]
	}

	ms.stack = ms.stack[:ms.stackLastIndex()]
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return 0
	}
	return ms.stack[ms.stackLastIndex()]
}

func (ms *MinStack) GetMin() int {
	if len(ms.minStack) == 0 {
		return 0
	}
	return ms.minStack[ms.minStackLastIndex()]
}

func (ms *MinStack) minStackLastIndex() int {
	return len(ms.minStack) - 1
}

func (ms *MinStack) stackLastIndex() int {
	return len(ms.stack) - 1
}
