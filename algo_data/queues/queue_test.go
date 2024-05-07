package queues

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()

	got := queue.Dequeue()
	if got != 0 {
		t.Error("expected queue always return the zero value of seted type when is empty")
	}

	got = queue.Read()
	if got != 0 {
		t.Error("expected queue always return the zero value of seted type when is empty")
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
		t.Error("expected queue always return the zero value of seted type when is empty")
	}
}
