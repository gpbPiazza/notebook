package minstack

import "testing"

func TestCache_CASE1(t *testing.T) {
	// ["MinStack","push","push","push","getMin","pop","top","getMin"]
	// [[],[-2],[0],[-3],[],[],[],[]]

	minStack := NewMinStack()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	min := minStack.GetMin()
	if min != -3 {
		t.Errorf("want: -3 got: %d", min)
	}

	minStack.Pop()

	val := minStack.Top()
	if val != 0 {
		t.Errorf("want: -3 got: %d", 0)
	}

	min = minStack.GetMin()
	if min != -2 {
		t.Errorf("want: -2 got: %d", min)
	}
}

func TestMinStack_CASE2(t *testing.T) {
	// ["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
	// [[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]

	minStack := NewMinStack()

	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)

	val := minStack.Top()
	if val != 2147483647 {
		t.Errorf("want: 2147483647 got: %d", val)
	}

	minStack.Pop()

	min := minStack.GetMin()
	if min != 2147483646 {
		t.Errorf("want: 2147483646 got: %d", min)
	}

	minStack.Pop()

	min = minStack.GetMin()
	if min != 2147483646 {
		t.Errorf("want: 2147483646 got: %d", min)
	}

	minStack.Pop()

	minStack.Push(2147483647)

	val = minStack.Top()
	if val != 2147483647 {
		t.Errorf("want: 2147483647 got: %d", val)
	}

	min = minStack.GetMin()
	if min != 2147483647 {
		t.Errorf("want: 2147483647 got: %d", min)
	}

	minStack.Push(-2147483648)

	val = minStack.Top()
	if val != -2147483648 {
		t.Errorf("want: -2147483648 got: %d", val)
	}

	min = minStack.GetMin()
	if min != -2147483648 {
		t.Errorf("want: -2147483648 got: %d", min)
	}

	minStack.Pop()

	min = minStack.GetMin()
	if min != 2147483647 {
		t.Errorf("want: 2147483647 got: %d", min)
	}
}

func TestMinStack_CASE3(t *testing.T) {
	minStack := NewMinStack()

	minStack.Push(123)

	minStack.Pop()

	min := minStack.GetMin()
	if min != 0 {
		t.Errorf("want: 0 got: %d", min)
	}
}
