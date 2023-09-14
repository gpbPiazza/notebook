package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	fileMatches []string
	waitGroup   = sync.WaitGroup{}
	locker      = sync.Mutex{}
)

func concurrentSearch(root, filename string) {
	fmt.Println("Searching in", root)
	files, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, file := range files {
		if strings.Contains(strings.ToUpper(file.Name()), strings.ToUpper(filename)) {
			locker.Lock()
			fileMatches = append(fileMatches, filepath.Join(root, file.Name()))
			locker.Unlock()
		}

		if file.IsDir() {
			waitGroup.Add(1)
			go concurrentSearch(filepath.Join(root, file.Name()), filename)
		}
	}
	waitGroup.Done()
}

func linearSearch(root, filename string) {
	fmt.Println("Searching in", root)
	files, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, file := range files {
		if strings.Contains(strings.ToUpper(file.Name()), strings.ToUpper(filename)) {
			fileMatches = append(fileMatches, filepath.Join(root, file.Name()))
		}

		if file.IsDir() {
			linearSearch(filepath.Join(root, file.Name()), filename)
		}
	}
}

func searchController(root, fileName string) {
	waitGroup.Add(1)
	go concurrentSearch(root, fileName)
	waitGroup.Wait()
}

func benchMark() {
	root, filename := "/Users/olaisaac/projects/repos/notebook", "readme"

	linearStart := time.Now()
	linearSearch(root, filename)
	linearEnd := time.Since(linearStart)

	concurrentStart := time.Now()
	searchController(root, filename)
	concurrentEnd := time.Since(concurrentStart)

	if linearEnd > concurrentEnd {
		fmt.Println("concurrent search is faster!")
		fmt.Println("-> Linear executation Time:", linearEnd)
		fmt.Println("-> Concurrent executation Time:", concurrentEnd)
		return
	}
	fmt.Println("linear search is faster!")
	fmt.Println("-> Linear executation Time:", linearEnd)
	fmt.Println("-> Concurrent executation Time:", concurrentEnd)
}

func main() {
	benchMark()

	// searchController("/Users/olaisaac/projects/repos/notebook", "readme")
	// if len(fileMatches) == 0 {
	// 	fmt.Println("no files found")
	// }

	// for _, file := range fileMatches {
	// 	fmt.Println("Matched: ", file)
	// }
}
