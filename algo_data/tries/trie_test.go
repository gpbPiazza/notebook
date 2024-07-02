package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestTrie_AllWordsInWordOf(t *testing.T) {
	trie := NewTrie()

	trie.Insert("zap")
	trie.Insert("zapper")
	trie.Insert("zap2")
	trie.Insert("zap3")
	trie.Insert("zap4")
	trie.Insert("zapDaGama")
	trie.Insert("aiai")
	trie.Insert("gremio")
	trie.Insert("mengudo")
	trie.Insert("mengo")
	trie.Insert("men")
	trie.Insert("mens")
	trie.Insert("mensalidade")

	zapNode := trie.SearchPrefix("zap")

	words := make([]string, 0)
	trie.AllWordsInWordOf(zapNode, &words, "zap")
	if len(words) != 6 {
		t.Error("expected to have 6 words starting with zap")
	}

	words2 := make([]string, 0)
	trie.AllWordsInWordOf(nil, &words2, "")
	if len(words2) != 13 {
		t.Error("expected to have 13 words")
	}
}

func TestTrie_CHAPTER17_EXERCISE2(t *testing.T) {
	trie := NewTrie()

	words := []string{
		"get",
		"go",
		"got",
		"gotten",
		"hall",
		"ham",
		"hammer",
		"hill",
		"zebra",
	}
	for _, word := range words {
		trie.Insert(word)
	}
	var got []string
	trie.AllWordsInWordOf(nil, &got, "")

	assert.ElementsMatch(t, got, words)
}

func TestTrie_TransversePrintln(t *testing.T) {
	trie := NewTrie()

	words := []string{
		"tag",
		"tan",
		"tankk",
		"tap",
		"today",
		"total",
		"we",
		"well",
		"went",
	}
	for _, word := range words {
		trie.Insert(word)
	}

	TransversePrintln(trie.root)
}

func TestTrie_AutocorrectWord(t *testing.T) {
	trie := NewTrie()

	words := []string{
		"tag",
		"tan",
		"tankk",
		"tap",
		"today",
		"total",
		"we",
		"well",
		"went",
	}
	for _, word := range words {
		trie.Insert(word)
	}

	t.Run("given a word not existing in the trie return all words with common prefix of TO", func(t *testing.T) {
		got := trie.AutocorrectWord("top")

		assert.ElementsMatch(t, []string{"today", "total"}, got)
	})

	t.Run("return the same word when word exist in the trie", func(t *testing.T) {
		got := trie.AutocorrectWord("total")

		assert.ElementsMatch(t, []string{"total"}, got)
	})

	t.Run("return all letter that starts with w", func(t *testing.T) {
		got := trie.AutocorrectWord("w")

		assert.ElementsMatch(t, []string{"we", "well", "went"}, got)
	})
}
