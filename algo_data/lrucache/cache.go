package lrucache

// this LRU cache are using Least Recently Used
// if get a data he is the most recently used
// if data is updated or inserted he is the most recently used
type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

func NewNode(key int, val int) *Node {
	return &Node{
		Key:  key,
		Val:  val,
		Next: nil,
		Prev: nil,
	}
}

type LRUCache struct {
	capacity int
	right    *Node // mRU - most recently used
	left     *Node // lRU - last recently used
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	left := NewNode(0, 0)
	right := NewNode(0, 0)
	right.Prev = left
	left.Next = right
	return LRUCache{
		capacity: capacity,
		right:    right,
		left:     left,
		cache:    make(map[int]*Node),
	}
}

func (lru *LRUCache) Put(key int, data int) {
	node, exist := lru.cache[key]
	newNode := NewNode(key, data)
	if exist {
		lru.removeFromList(node)
	}

	lru.cache[key] = newNode
	lru.insertAtRight(lru.cache[key])

	if lru.capacity < len(lru.cache) {
		// remove the last recently used from the list and from the cache
		lruElement := lru.left.Next
		lru.removeFromList(lruElement)
		delete(lru.cache, lruElement.Key)
	}
}

func (lru *LRUCache) Get(key int) int {
	node, exists := lru.cache[key]
	if exists {
		lru.makeNodeMRU(node)
		return node.Val
	}

	return -1
}
func (lru *LRUCache) makeNodeMRU(node *Node) {
	lru.removeFromList(node)
	lru.insertAtRight(node)
}

func (lru *LRUCache) removeFromList(node *Node) {
	prev, next := node.Prev, node.Next
	prev.Next = next
	next.Prev = prev
}

func (lru *LRUCache) insertAtRight(node *Node) {
	prev := lru.right.Prev
	lru.right.Prev = node
	prev.Next = node
	node.Next = lru.right
	node.Prev = prev
}
