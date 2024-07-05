package queues

import "testing"

func Test_SliceQueue(t *testing.T) {
	queue := NewSliceQueue[int]()
	QueueTest(t, queue)
}

func Test_DLLQQueue(t *testing.T) {
	queue := newDLLQueue[int]()
	QueueTest(t, queue)
}

func QueueTest[T int](t *testing.T, queue Queue[int]) {
	got := queue.Dequeue()
	if got != 0 {
		t.Error("expected queue always return 0 when is empty, in Dequeue")
	}

	got = queue.Read()
	if got != 0 {
		t.Error("expected queue always return nil when is empty, in Read")
	}

	queue.Enqueue(10)
	queue.Enqueue(11)
	queue.Enqueue(12)
	queue.Enqueue(13)
	got = queue.Dequeue()
	if got != 10 {
		t.Error("expected queue always return the element 10 when dequeue")
	}

	got = queue.Dequeue()
	if got != 11 {
		t.Error("expected queue always return the element 11 when dequeue")
	}

	got = queue.Read()
	if got != 12 {
		t.Error("expected queue always return the element 12 when read")
	}

	got = queue.Read()
	if got != 12 {
		t.Error("expected queue always return the element 12 when reads twice")
	}

	got = queue.Dequeue()
	if got != 12 {
		t.Error("expected queue always return the element 12 when dequeue")
	}

	got = queue.Read()
	if got != 13 {
		t.Error("expected queue always return the element 13 when Read")
	}

	got = queue.Dequeue()
	if got != 13 {
		t.Error("expected queue always return the element 13 when dequeue")
	}

	got = queue.Read()
	if got != 0 {
		t.Error("expected queue always return 0 when is empty, in Read")
	}
}
