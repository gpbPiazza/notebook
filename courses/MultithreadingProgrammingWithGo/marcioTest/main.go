package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func zape(n int) {
	// fmt.Println("Oooi eu sou trhead de número", n)
	logger := log.Default()
	logger.Println("Oooi eu sou trhead de número", n)
	time.Sleep(3 * time.Second)
	logger.Println("Oooi eu sou trhead de número", n, "e esperei 3 segundos")
	wg.Done()
}

func main() {
	limiteOfThreads := 5
	fmt.Println("START TEST")
	wg.Add(limiteOfThreads)
	for i := 0; i < limiteOfThreads; i++ {
		go zape(i + 1)
	}
	wg.Wait()
	fmt.Println("END TEST")
}
