package tries

import "fmt"

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

// allWordsInWordOf returns a slice of words starting by the given word
// given nil node will assume to start from the root trie
// The result will be concatened in the slice words by reference
func (t *Trie) AllWordsInWordOf(node *Node, words *[]string, word string) {
	if node == nil {
		node = t.root
	}

	for char, val := range node.children {
		if char == rune('*') {
			*words = append(*words, word)
		} else {
			t.AllWordsInWordOf(val, words, word+string(char))
		}
	}
}

func TransversePrintln(node *Node) {
	if node == nil {
		return
	}

	for key, child := range node.children {
		TransversePrintln(child)
		fmt.Println(string(key))
	}
}

// Autoccorect word returns a suggestions fix typo
// by the given word. It Will return the first word that shares the bigger amount of
// the prefixes with the existing words in the trie.
func (t *Trie) AutocorrectWord(word string) []string {
	searchCommonPrefix := func(word string) (string, *Node) {
		currentNode := t.root
		var commonPrefix string
		for _, char := range word {
			node, ok := currentNode.children[char]
			if !ok {
				return commonPrefix, currentNode
			}
			commonPrefix += string(char)
			currentNode = node
		}

		return commonPrefix, currentNode
	}

	commonPrefix, commonNode := searchCommonPrefix(word)

	var result []string
	t.AllWordsInWordOf(commonNode, &result, "")

	for i := range result {
		word := result[i]
		result[i] = commonPrefix + word
	}

	return result
}
