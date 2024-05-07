package stacks

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack[int](3)
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	got := stack.Read()
	if got != -3 {
		t.Error("expected last element in the stack to be -3")
	}

	stack.Pop()

	got = stack.Read()
	if got != 0 {
		t.Error("expected last element in the stack to be 0")
	}

	stack.Pop()

	got = stack.Read()
	if got != -2 {
		t.Error("expected last element in the stack to be -2")
	}

	_ = stack.Pop()
	got = stack.Pop()
	if got != 0 {
		t.Error("expected stack always return the zero value of seted type when is empty")
	}

	got = stack.Read()
	if got != 0 {
		t.Error("expected stack always return the zero value of seted type when is empty")
	}
}
