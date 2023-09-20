package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = sync.Mutex{}
	lock2 = sync.Mutex{}
)

func blueRobot() {
	for {
		fmt.Println("Blue: Acquiring lock1")
		lock1.Lock()
		fmt.Println("Blue: Acquiring lock2")
		lock2.Lock()
		fmt.Println("Blue: both locks acquired")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Blue: both locks released")
	}
}

func redRobot() {
	for {
		fmt.Println("Red: Acquiring lock2")
		lock2.Lock()
		fmt.Println("Red: Acquiring lock1")
		lock1.Lock()
		fmt.Println("Red: both locks acquired")
		lock2.Unlock()
		lock1.Unlock()
		fmt.Println("Red: both locks released")
	}
}

func childThread() {
	for {
		fmt.Println("I'm alive")
	}
}

// a example of when the main threads finish the child will terminates too
func mainThread() {
	fmt.Println("Start child")
	go childThread()
	time.Sleep(10 * time.Second)
	fmt.Println("Done")
}

func main() {
	go redRobot()
	go blueRobot()
	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}
