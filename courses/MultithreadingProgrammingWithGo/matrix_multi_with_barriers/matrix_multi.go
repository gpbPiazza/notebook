package main

import (
	"fmt"
	"math/rand"
	"time"
)

const matrixSize = 5

var (
	matrixA      = [matrixSize][matrixSize]int{}
	matrixB      = [matrixSize][matrixSize]int{}
	result       = [matrixSize][matrixSize]int{}
	workStart    = NewBarrier(matrixSize + 1)
	workComplete = NewBarrier(matrixSize + 1)
)

func generateRnadomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for column := 0; column < matrixSize; column++ {
			matrix[row][column] += rand.Intn(20) - 10
		}
	}
}

func multiply(row int) {
	for {
		workStart.Wait()
		for column := 0; column < matrixSize; column++ {
			for i := 0; i < matrixSize; i++ {
				result[row][column] += matrixA[row][i] * matrixB[i][column]
			}
		}
		workComplete.Wait()
	}
}

func main() {
	now := time.Now()
	fmt.Println("Start working...")

	for row := 0; row < matrixSize; row++ {
		go multiply(row)
	}

	for i := 0; i < 100; i++ {
		generateRnadomMatrix(&matrixA)
		generateRnadomMatrix(&matrixB)
		workStart.Wait()
		workComplete.Wait()
	}
	fmt.Println("Work done! Process took:", time.Since(now))

}
