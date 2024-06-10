package linkedlist

import "testing"

func makeLinkedList() *LinkedList2 {
	node1 := NewNode2("vamo")
	node2 := NewNode2("gremio")
	node3 := NewNode2("fiu fiu")
	node4 := NewNode2("!!")

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	return NewLinkedList2(node1)
}

func itPanics(f func()) (isPaniced bool) {
	defer func() {
		val := recover()
		if val != nil {
			isPaniced = true
		}
	}()

	f()
	return isPaniced
}

func Test_LinkedList2_Search(t *testing.T) {
	linkedList := makeLinkedList()

	t.Run("return gremio When gremio is in the Search", func(t *testing.T) {
		val := linkedList.Search("gremio")
		if val != "gremio" {
			t.Errorf("expected gremio but got %v", val)
		}
	})

	t.Run("return vamo When vamo is in the Search", func(t *testing.T) {
		val := linkedList.Search("vamo")
		if val != "vamo" {
			t.Errorf("expected vamo but got %v", val)
		}
	})

	t.Run("return !! When !! is in the Search", func(t *testing.T) {
		val := linkedList.Search("!!")
		if val != "!!" {
			t.Errorf("expected !! but got %v", val)
		}
	})

	t.Run("nil when val is not found", func(t *testing.T) {
		val := linkedList.Search("mengudo")
		if val != nil {
			t.Error("expected val to be nil")
		}
	})
}

func Test_LinkedList2_Read(t *testing.T) {
	linkedList := makeLinkedList()

	t.Run("return  gremio When index is 1", func(t *testing.T) {
		val := linkedList.Read(1)

		if val != "gremio" {
			t.Errorf("expected gremio but got %v", val)
		}
	})

	t.Run("return vamo When index is 0", func(t *testing.T) {
		val := linkedList.Read(0)

		if val != "vamo" {
			t.Errorf("expected vamo but got %v", val)
		}
	})

	t.Run("panic when index is out of range", func(t *testing.T) {
		iPaniced := itPanics(func() { linkedList.Read(100) })

		if !iPaniced {
			t.Errorf("expected that paniced")
		}
	})
}

func Test_LinkedList2_InserAtIndex(t *testing.T) {
	t.Run("insert ZAP at index 2", func(t *testing.T) {
		node1 := NewNode2("vamo")
		node2 := NewNode2("gremio")
		node3 := NewNode2("fiu fiu")
		node4 := NewNode2("!!")

		node1.Next = node2
		node2.Next = node3
		node3.Next = node4

		list := NewLinkedList2(node1)

		val := list.Read(1)
		if val != "gremio" {
			t.Errorf("at index 1 has val gremio got %v", val)
		}

		list.InserAtIndex(2, "ZAP")

		val = list.Read(2)
		if val != "ZAP" {
			t.Errorf("expected index 2 has val ZAP got %v", val)
		}

		val = list.Read(1)
		if val != "gremio" {
			t.Errorf("at index 1 has val gremio got %v", val)
		}

		val = list.Read(3)
		if val != "fiu fiu" {
			t.Errorf("at index 3 has val fiu fiu got %v", val)
		}
	})

	t.Run("insert node at head", func(t *testing.T) {
		node1 := NewNode2("vamo")
		node2 := NewNode2("gremio")
		node3 := NewNode2("fiu fiu")
		node4 := NewNode2("!!")

		node1.Next = node2
		node2.Next = node3
		node3.Next = node4

		list := NewLinkedList2(node1)

		val := list.Read(0)
		if val != "vamo" {
			t.Errorf("at index 1 has val vamo got %v", val)
		}

		list.InserAtIndex(0, "poison_talk_dirty_to_me")

		val = list.Read(0)
		if val != "poison_talk_dirty_to_me" {
			t.Errorf("expected index 0 has val poison_talk_dirty_to_me got %v", val)
		}
	})

	t.Run("insert node at head", func(t *testing.T) {
		node1 := NewNode2("vamo")
		node2 := NewNode2("gremio")
		node3 := NewNode2("fiu fiu")
		node4 := NewNode2("!!")

		node1.Next = node2
		node2.Next = node3
		node3.Next = node4

		list := NewLinkedList2(node1)

		val := list.Read(3)
		if val != "!!" {
			t.Errorf("at index 4 has val !! got %v", val)
		}

		list.InserAtIndex(3, "key_west_all_of_the_lights")

		val = list.Read(3)
		if val != "key_west_all_of_the_lights" {
			t.Errorf("expected index 4 has val key_west_all_of_the_lights got %v", val)
		}

		val = list.Read(4)
		if val != "!!" {
			t.Errorf("at index 4 has val !! got %v", val)
		}
	})

	t.Run("insert in a none existent index should panics", func(t *testing.T) {
		node1 := NewNode2("vamo")
		node2 := NewNode2("gremio")
		node3 := NewNode2("fiu fiu")
		node4 := NewNode2("!!")

		node1.Next = node2
		node2.Next = node3
		node3.Next = node4

		list := NewLinkedList2(node1)

		panicked := itPanics(func() { list.InserAtIndex(10, "should_panics") })

		if !panicked {
			t.Error("expected to panicked")
		}
	})
}
