package lrucache

// this LRU cache are using Least Recently Used
// what is the criteria to define user? Get actions update used? Or just insert or update?
type Lrucache struct {
	capacity int
	stack    []string
	store    map[string]any
}

func New(capacity int) *Lrucache {
	return &Lrucache{
		capacity: capacity,
	}
}

func (lru *Lrucache) Put(key string, data any) {
	_, ok := lru.store[key]
	if !ok {
		if len(lru.stack) == lru.capacity {
			lru.deletLastInserted()
		}

		lru.stack = append(lru.stack, key)
		lru.store[key] = data
		return
	} else {
		// aqui o elemento existe no cache e precisa ser atualizad de posição e valor
		// teria que implementar um algoritmo de find  index e remove element from array
		// preciso colocar esse cara no topo e deletar a sua referência na stack
	}

	lru.store[key] = data
	lru.stack = append(lru.stack, key)
}

func (lru *Lrucache) Get(key string) any {
	return lru.store[key]
}

func (lru *Lrucache) deletLastInserted() {
	lastIndex := len(lru.stack) - 1
	lastElementInsertedKey := lru.stack[lastIndex]
	lru.stack = lru.stack[:lastIndex]
	delete(lru.store, lastElementInsertedKey)
}
