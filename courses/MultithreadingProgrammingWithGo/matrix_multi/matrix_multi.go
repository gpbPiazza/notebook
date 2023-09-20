package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const matrixSize = 5

var (
	matrixA = [matrixSize][matrixSize]int{}
	matrixB = [matrixSize][matrixSize]int{}
	result  = [matrixSize][matrixSize]int{}
	rwLock  = sync.RWMutex{}
	cond    = sync.NewCond(rwLock.RLocker())
	wg      = sync.WaitGroup{}
)

func generateRnadomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for column := 0; column < matrixSize; column++ {
			matrix[row][column] += rand.Intn(20) - 10
		}
	}
}

func multiply(row int) {
	rwLock.RLock()
	for {
		wg.Done()
		cond.Wait() // after the first init of the go routine the execution line holds here until conditional.BroadCast
		for column := 0; column < matrixSize; column++ {
			for i := 0; i < matrixSize; i++ {
				result[row][column] += matrixA[row][i] * matrixB[i][column]
			}
		}
	}
}

// How this works?
// for each row to multply wer create a goroutine to do this work
// so if we have a 3x3 matrix to multply we gonna have 3 goroutines to multiply each row
// and we do this 100 times just to see the diference between this code vs the linear one
// for each new matrix generated the main function call the broadCast and each
// goroutine process his repective line.

func main() {
	now := time.Now()
	fmt.Println("Start working...")

	wg.Add(matrixSize)
	for row := 0; row < matrixSize; row++ {
		go multiply(row)
	}

	for i := 0; i < 100; i++ {
		wg.Wait()
		rwLock.Lock()
		generateRnadomMatrix(&matrixA)
		generateRnadomMatrix(&matrixB)
		wg.Add(matrixSize)
		rwLock.Unlock()
		cond.Broadcast()
	}
	fmt.Println("Work done! Process took:", time.Since(now))

}
