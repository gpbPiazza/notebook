package main

import (
	"container/list"
	"fmt"
)

func main() {

	// myConexations := newMangoSellerGrapht()
	// _ = findMangoSeller(myConexations)
	// Complexidade Big O notations
	// O(número de arestas/conexations)
	// Por adicionar as pessoas em uma lista de ja checadas O(número de nós)
	// Portanto a complexidade total fica O(número de arestas + número de nós) = O(V + A)
	// V número de vérticies/nós e A número de arestas

	grapth := newPathGrapht()
	path := shortPathBetweenVertices("jato", "gado", grapth)
	expected := []string{"jato", "gato", "gado"}

	for i, got := range path {
		if got != expected[i] {
			fmt.Println("Error!\nexpected %s \nbut got -> %s", expected[i], got)
		}
	}

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

func newMangoSellerGrapht() map[string][]string {
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

func shortPathBetweenVertices(start, end string, grapth map[string][]string) []string {
	results := make(map[int][]string, 0)
	conexations := grapth[start]
	//1 Pega as conexões
	//2 Concatena cada conexão no map de resultados
	//3 Vefica se o elemento concatenado é o final
	// 3.1 Caso for return resultado
	// 3.2 Caso não repete passo 1

	// 1. acessa node start
	// 2. percorre sua conexões
	// 3. concatena cada conexão em um slice de results
	// 4. Repete isso até acabar todas as conexões do grafo
	// 5. Com o map de possíveis paths feito encontre aquel que possuí start como primeiro e end como último elementos
	// 6. Dentre esses encontrados aquel que tiver o menor comprimento é o menor caminho

	for i, conexation := range conexations {
		results[i] = []string{start, conexation}
		if conexation == end {
			return results[i]
		}

	}
	return nil
}

func newPathGrapht() map[string][]string {
	grapth := make(map[string][]string, 0)

	grapth["jato"] = []string{"tato", "gato"}
	grapth["tato"] = []string{"chato", "gato"}
	grapth["gato"] = []string{"gado", "grato"}
	grapth["grato"] = []string{"gado"}
	grapth["gado"] = []string{}
	grapth["chato"] = []string{"gado"}

	fmt.Println(grapth)
	return grapth
}
