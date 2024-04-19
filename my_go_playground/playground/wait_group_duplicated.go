package playground

import (
	"fmt"
	"sync"
	"time"
)

type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

func waitGroup1() {
	var wg sync.WaitGroup
	fmt.Println("waitGroup 1, valor ->", wg)
	for i := range []int{1, 2, 3, 4} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("waitGroup 1, goroutine ->", i)
			time.Sleep(1 * time.Second)
		}()
	}
	fmt.Println("waitGroup 1, waiting for all routines end")
	wg.Wait()
	fmt.Println("waitGroup 1, finish")
}

func waitGroup2() {
	var wg sync.WaitGroup
	fmt.Println("waitGroup 2, valor ->", wg)
	for i := range []int{1, 2, 3, 4} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("waitGroup 2, goroutine ->", i)
			time.Sleep(1 * time.Second)
		}()
	}
	fmt.Println("waitGroup 2, waiting for all routines end")
	wg.Wait()
	fmt.Println("waitGroup 2, finish")
}

func ExecuteWaitGrouDuplicated() {
	fmt.Println("waitGroup FATHER start")
	var wg sync.WaitGroup
	wg.Add(1)
	go waitGroup1()
	wg.Add(1)
	go waitGroup2()
	fmt.Println("waitGroup FATHER waiting")
	wg.Wait()
	fmt.Println("waitGroup FATHER finish")
}
