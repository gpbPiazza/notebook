package lrucache

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func NewNode(key, val int) *Node {
	return &Node{Val: val, Key: key}
}

type LRUCache struct {
	capacity int

	cache map[int]*Node
	left  *Node // left.Next = LRU
	right *Node // right.Previus = MRU
}

func Constructor(capacity int) LRUCache {
	left := NewNode(0, 0)
	right := NewNode(0, 0)

	left.Next = right
	right.Prev = left
	return LRUCache{
		capacity: capacity,
		left:     left,
		right:    right,
		cache:    make(map[int]*Node),
	}
}

func (lru *LRUCache) Put(key int, data int) {
	node, exists := lru.cache[key]
	if exists {
		lru.removeNodeFromList(node)
	}
	newNode := NewNode(key, data)

	lru.insertAtRight(newNode)
	lru.cache[key] = newNode

	if len(lru.cache) > lru.capacity {
		lruElement := lru.left.Next
		lru.removeNodeFromList(lruElement)
		delete(lru.cache, lruElement.Key)
	}
}

func (lru *LRUCache) removeNodeFromList(node *Node) {
	prev, next := node.Prev, node.Next

	prev.Next = next
	next.Prev = prev
}

func (lru *LRUCache) insertAtRight(newNode *Node) {
	lastMRU := lru.right.Prev

	lru.right.Prev = newNode
	newNode.Next = lru.right
	lastMRU.Next = newNode
	newNode.Prev = lastMRU
}

func (lru *LRUCache) Get(key int) int {
	node, exists := lru.cache[key]
	if !exists {
		return -1
	}

	lru.removeNodeFromList(node)
	lru.insertAtRight(node)

	return node.Val
}
