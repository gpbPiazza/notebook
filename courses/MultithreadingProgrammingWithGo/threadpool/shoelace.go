package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Point2D struct {
	x int
	y int
}

const (
	numberOfThreads = 5
)

// With amount of work constant and change the number ot threads will not result in a linear performance gain
// As said in Amdahl's law.
// Processing took: 563.942166ms 17 number of threads
// Processing took: 566.587375ms 16 number of threads
// Processing took: 659.999166ms 15 number of threads
// Processing took: 690.881208ms 10 number of threads
// Processing took: 1.004313209s 9 number of threads
// Processing took: 866.354625ms 8 number of threads
// Processing took: 840.43575ms  7 number of threads
// Processing took: 638.301875ms 6 number of threads
// Processing took: 892.49325m   5 number of threads
// Processing took: 716.986125ms 4 number of threads
// Processing took: 930.539667ms 3 number of threads
// Processing took: 805.170334ms 2 number of threads
// Processing took: 1.52431325s  1 number of threads

// We gonna have a pool/queu of itens to process with 1000 spaces, translated into
// conexationsPool variable, each worker will be connected with the pool and will process each request.
// We will have 10 woekers, translated in the const numberOfThreads

var (
	r           = regexp.MustCompile(`\((\d*),(\d*)\)`)
	waitGroup   = sync.WaitGroup{}
	locker      = sync.Mutex{}
	areasResult []float64
)

func calcArea(inputChan chan string) {
	for pointsStr := range inputChan {
		var points []Point2D

		for _, point := range r.FindAllStringSubmatch(pointsStr, -1) {
			x, _ := strconv.Atoi(point[1])
			y, _ := strconv.Atoi(point[2])
			points = append(points, Point2D{x: x, y: y})
		}

		var area float64
		for i := 0; i < len(points); i++ {
			neverWillBetOutOfRange := (i + 1) % len(points)
			a, b := points[i], points[neverWillBetOutOfRange]
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}

		area = math.Abs(area) / 2
		fmt.Println("Polygon area:", area)
		locker.Lock()
		areasResult = append(areasResult, area)
		locker.Unlock()
	}
	waitGroup.Done()
}

func main() {
	absPaht, err := filepath.Abs("./polygons.txt")
	if err != nil {
		log.Fatal(err)
	}
	dat, err := os.ReadFile(absPaht)
	if err != nil {
		log.Fatal(err)
	}
	polyTxt := string(dat)

	start := time.Now()
	masterInitWorker(polyTxt)
	end := time.Since(start)
	fmt.Println("Processing took:", end)
	// fmt.Println("Polygons areas:", areasResult)
}

func masterInitWorker(polyTxt string) {
	conexationsPool := make(chan string, 1000)
	initWorkers(conexationsPool)
	for _, line := range strings.Split(polyTxt, "\n") {
		conexationsPool <- line
	}
	close(conexationsPool)
	waitGroup.Wait()
}

func initWorkers(conexationsPool chan string) {
	for i := 0; i < numberOfThreads; i++ {
		go calcArea(conexationsPool)
	}
	waitGroup.Add(numberOfThreads)
}
