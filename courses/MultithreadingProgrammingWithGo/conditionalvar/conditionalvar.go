package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	money       = 100
	locker      = sync.RWMutex{}
	conditional = sync.NewCond(&locker)
	wg          = sync.WaitGroup{}
)

func expense() {
	for i := 0; i <= 100; i++ {
		locker.Lock()
		for money-20 < 0 {
			conditional.Wait()
		}
		money -= 20
		fmt.Println("expense see balance:", money)
		locker.Unlock()
	}

}

func income() {
	for i := 0; i <= 100; i++ {
		locker.Lock()
		money += 10
		fmt.Println("income see balance:", money)
		locker.Unlock()
		conditional.Signal()
		time.Sleep(1 * time.Millisecond)
	}

}

func main() {
	go income()
	go expense()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Result of transactions => total money:", money)
}
