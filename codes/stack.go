package main

type MinStack struct {
	min  []int
	data []int
}

func NewMinStack() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  make([]int, 0),
	}
}

func (ms *MinStack) Push(val int) {
	if len(ms.data) == 0 {
		ms.min = append(ms.min, val)
	}

	if val <= ms.min[len(ms.min)-1] {
		ms.min = append(ms.min, val)
	}

	ms.data = append(ms.data, val)
}

func (ms *MinStack) Pop() {
	lasMinIndex := len(ms.min) - 1
	lastDataIndex := len(ms.data) - 1

	topData := ms.data[lastDataIndex]
	topMin := ms.min[lasMinIndex]

	if topMin == topData {
		ms.min = ms.min[:lasMinIndex]
	}

	ms.data = ms.data[:lastDataIndex]
}

func (ms *MinStack) Top() int {
	lastIndex := len(ms.data) - 1
	return ms.data[lastIndex]
}

func (ms *MinStack) GetMin() int {
	lastIndex := len(ms.min) - 1
	return ms.min[lastIndex]
}
