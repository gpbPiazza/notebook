package main

import (
	"container/list"
	"fmt"
)

func main() {

	myConexations := newGrapht()

	_ = findMangoSeller(myConexations)
	// Complexidade Big O notations
	// O(número de arestas/conexations)
	// Por adicionar as pessoas em uma lista de ja checadas O(número de nós)
	// Portanto a complexidade total fica O(número de arestas + número de nós) = O(V + A)
	// V número de vérticies/nós e A número de arestas
}
func isMangoSeller(person string) bool {
	return person[0:1] == "M"
}

func findMangoSeller(peopleConexations map[string][]string) string {
	queue := list.New()
	// myConexations := peopleConexations["me"]
	queue.PushFront("me")

	alreadyCheck := make(map[string]bool, 0)

	for {
		element := queue.Front()
		if element == nil {
			return "there is no mango sellers in this conexations"
		}
		person, ok := element.Value.(string)
		if !ok {
			fmt.Println(element)
			return "element in queue is not a string"
		}
		fmt.Println("queue state ->", queue.Len())
		fmt.Println("checking ->", person)

		if alreadyCheck[person] {
			fmt.Println("already checked, remove from queue ->", person)
			queue.Remove(element)
			continue
		}

		if isMangoSeller(person) {
			fmt.Printf("the person %s is a mango seller", person)
			return person
		}
		fmt.Println("this person is not a mango seller ->", person)
		fmt.Println("queuing", peopleConexations[person])
		for _, conexations := range peopleConexations[person] {
			queue.PushBack(conexations)
		}
		alreadyCheck[person] = true
		queue.Remove(element)
	}
}

func newGrapht() map[string][]string {
	grapth := make(map[string][]string, 0)

	grapth["me"] = []string{"alice", "bob", "claire"}
	grapth["bob"] = []string{"anuj", "peggy", "Mango Seller Jilson"}
	grapth["alice"] = []string{"peggy"}
	grapth["anuj"] = []string{}
	grapth["peggy"] = []string{}
	grapth["thom"] = []string{}
	grapth["jonny"] = []string{}
	grapth["Mango Seller Jilson"] = []string{"claire", "thom"}
	fmt.Println(grapth)
	return grapth
}

type hashTable struct{}

func (ht hashTable) Get(key string) any {
	return ht
}

func newHashTable() hashTable {
	return hashTable{}
}
