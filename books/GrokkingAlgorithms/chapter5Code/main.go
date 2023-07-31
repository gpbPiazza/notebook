package main

func main() {
}

type hashTable struct{}

func (ht hashTable) Get(key string) any {
	return ht
}

func newHashTable() hashTable {
	return hashTable{}
}
