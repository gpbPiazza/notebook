package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Aprendizado aqui, quando criamos um chann e escutamos o seu resultado ao final da função PAI
// A função PAI fica presa na linha que está escutando o resultado final do Chann pr
func main() {
	resultChan := make(chan []int)
	fmt.Println("AAA LIBERAMO A GO ROUTINE, PEGA TODOS OS NUMEROS PARES")
	go testChan("1,2,3,4,5,6,7,8,9,10,11,12.55,12,13,14,15", resultChan)
	fmt.Println("FUNÇÃO PAI ESTÁ ESCUTANDO O CHANN ANTES DE FINGER ALGUMA COISA")
	result := <-resultChan
	fmt.Println("FUNÇÃO PAI TERMIN OU DE ESCUTAR O CHANN  E AGORA VAI FINGER ALGUMA COISA", result)
	fmt.Println("BBBBB LIBERAMO A GO ROUTINE")

	for i := 0; i < 5; i++ {
		fmt.Println("finge que estamos fazendo outra coisa", i)
		time.Sleep(2 * time.Millisecond)
	}

	// result := <-resultChan
	// fmt.Println("FUNÇÃO PAI TERMINOU DE EXECUTAR, A GOROUTINE AINDA TA PROCESSANDO?")
	// fmt.Println("ESCUTANDO O CHANNEL RESULT", resultChan)
	close(resultChan)
	fmt.Println(result)
}

func testChan(text string, resultChannel chan []int) {
	fmt.Println("cccc tAMO NAA GO ROUTINE")

	eachStr := strings.Split(text, ",")
	var result []int
	for _, str := range eachStr {
		time.Sleep(2 * time.Second)
		val, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("DEU ERRO NO PARSE, QUERO QUE ACABE AQUI", err.Error())
			resultChannel <- result
			fmt.Println("CHANNEL DEPOIS NO ERRO", resultChannel)
			return
		}
		if val%2 == 0 {
			result = append(result, val)
		}
	}
	fmt.Println("ddd terminou a GO ROUTINE, ENVIANDO RESULTADO PARA O CHANNEL")
	fmt.Println("CHANNEL Antes", resultChannel)
	resultChannel <- result
	fmt.Println("CHANNEL DEPOIS", resultChannel)
}
