package tries

// A trie can have any number of child nodes
type Node struct {
	children map[rune]*Node
}

func NewNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
	}
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, char := range word {
		node, ok := currentNode.children[char]
		if ok {
			currentNode = node
		} else {
			newNode := NewNode()
			currentNode.children[char] = newNode
			currentNode = newNode
		}
	}
	currentNode.children[rune('*')] = nil
}

func (t *Trie) SearchPrefix(prefix string) *Node {
	currentNode := t.root

	for _, char := range prefix {
		node, ok := currentNode.children[char]
		if !ok {
			return nil
		}

		currentNode = node
	}

	return currentNode
}
