package lrucache

import (
	"testing"
)

func TestCache(t *testing.T) {
	// ["LRUCache","put","put","get","put","get","put","get","get","get"]
	// [[2],       [1,1], [2,2],[1], [3,3],[2],  [4,4],[1],  [3],   [4]]

	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	val := cache.Get(1)
	if val != 1 {
		t.Errorf("To key 1 should returen val 1: want: %d got: %d", 1, val)
	}

	cache.Put(3, 3)

	val = cache.Get(2)
	if val != -1 {
		t.Errorf("after get key 1 and put key 3 the key 2 should be the LRU and be removed from cache. want: -1 got: %d", val)
	}

	cache.Put(4, 4)

	val = cache.Get(1)
	if val != -1 {
		t.Errorf("after puted 4 and 3 keys shoudl remove key 1 from cache: want: %d got: %d", -1, val)
	}

	val = cache.Get(3)
	if val != 3 {
		t.Errorf("To key 3 should returen val 3: want: %d got: %d", 3, val)
	}

	val = cache.Get(4)
	if val != 4 {
		t.Errorf("To key 4 should returen val 3: want: %d got: %d", 4, val)
	}
}
