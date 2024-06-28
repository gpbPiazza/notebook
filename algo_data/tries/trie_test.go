package tries

import (
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := NewTrie()

	trie.Insert("ace")
	node, ok := trie.root.children[rune('a')]
	if !ok {
		t.Error("expected to have key a")
	}
	node, ok = node.children[rune('c')]
	if !ok {
		t.Error("expected to have key c")
	}
	node, ok = node.children[rune('e')]
	if !ok {
		t.Error("expected to have key e")
	}
	node, ok = node.children[rune('*')]
	if !ok {
		t.Error("expected to have key *")
	}
	if node != nil {
		t.Error("expected the last node to be nil'")
	}

	trie.Insert("act")
	node, ok = trie.root.children[rune('a')]
	if !ok {
		t.Error("expected to have key a")
	}
	node, ok = node.children[rune('c')]
	if !ok {
		t.Error("expected to have key c")
	}
	node, ok = node.children[rune('t')]
	if !ok {
		t.Error("expected to have key t")
	}
	node, ok = node.children[rune('*')]
	if !ok {
		t.Error("expected to have key *")
	}
	if node != nil {
		t.Error("expected the last node to be nil'")
	}

	trie.Insert("bad")
	node, ok = trie.root.children[rune('b')]
	if !ok {
		t.Error("expected to have key b")
	}
	node, ok = node.children[rune('a')]
	if !ok {
		t.Error("expected to have key a")
	}
	node, ok = node.children[rune('d')]
	if !ok {
		t.Error("expected to have key d")
	}
	node, ok = node.children[rune('*')]
	if !ok {
		t.Error("expected to have key *")
	}
	if node != nil {
		t.Error("expected the last node to be nil'")
	}

	trie.Insert("bat")
	node, ok = trie.root.children[rune('b')]
	if !ok {
		t.Error("expected to have key b")
	}
	node, ok = node.children[rune('a')]
	if !ok {
		t.Error("expected to have key a")
	}
	node, ok = node.children[rune('t')]
	if !ok {
		t.Error("expected to have key t")
	}
	node, ok = node.children[rune('*')]
	if !ok {
		t.Error("expected to have key *")
	}
	if node != nil {
		t.Error("expected the last node to be nil'")
	}

	trie.Insert("batter")
	node, ok = trie.root.children[rune('b')]
	if !ok {
		t.Error("expected to have key b")
	}
	node, ok = node.children[rune('a')]
	if !ok {
		t.Error("expected to have key a")
	}
	node, ok = node.children[rune('t')]
	if !ok {
		t.Error("expected to have key t")
	}
	node, ok = node.children[rune('t')]
	if !ok {
		t.Error("expected to have key t")
	}
	node, ok = node.children[rune('e')]
	if !ok {
		t.Error("expected to have key e")
	}
	node, ok = node.children[rune('r')]
	if !ok {
		t.Error("expected to have key e")
	}
	node, ok = node.children[rune('*')]
	if !ok {
		t.Error("expected to have key *")
	}
	if node != nil {
		t.Error("expected the last node to be nil'")
	}
}

func TestTrie_PrefixSearch(t *testing.T) {
	trie := NewTrie()

	trie.Insert("ace")
	trie.Insert("act")
	trie.Insert("bat")
	trie.Insert("batter")

	batNode := trie.SearchPrefix("bat")
	if batNode == nil {
		t.Error("expected to batNode not be nil")
	}
	_, ok := batNode.children[rune('*')]
	if !ok {
		t.Error("expected to batNode have a end of word")
	}

	_, ok = batNode.children[rune('t')]
	if !ok {
		t.Error("expected to batNodeChildren have another word starting with T")
	}
}
