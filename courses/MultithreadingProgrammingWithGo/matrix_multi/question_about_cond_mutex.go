package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mlock  = sync.Mutex{}
	condEx = sync.NewCond(&mlock)
)

func runChildThread() {
	mlock.Lock()
	fmt.Println("RunChildThread, lock acquired")
	condEx.Signal()
	fmt.Println("RunChildThread, Waiting")
	condEx.Wait()
	fmt.Println("RunChildThread, Running")
}

func RunMainThread() {
	mlock.Lock()
	fmt.Println("RunMainThread, lock acquired")
	go runChildThread()
	fmt.Println("RunMainThread, Waiting")
	condEx.Wait()
	fmt.Println("RunMainThread, Running")
	condEx.Signal()
	time.Sleep(10 * time.Second)
}

// Hi Dio,

// Thank you for signing up to the course. This is a tricky question designed to make you think about what happens when wait() and signal is called. Remember that rules are:

// 1. Only one thread can hold the lock at any point in time

// 2. wait() releases the lock and blocks the thread until a signal is called by another thread.

// 3. A signal() unblocks a thread that is currently waiting on a wait()

// 4. When a thread wakes up from a wait() it will try to acquire the lock back. If the lock is not available it will be blocked until it is.

// Here are the steps on what's happening in the program in this question:

// Time  MainThread                                         ChildThread
// ----  ----------                                         -----------
// 0:    Acquires mlock                                     Not started
// 1:    "RunMainThread, lock acquired"                     Not started
// 2:    Starts ChildThread                                 Tries to acquire mlock but has to wait
// 3:    "RunMainThread, Waiting"                           Waiting for mlock to be free
// 4:    Calls wait() -releasing mlock                      Acquires mlock
// 5:    Blocked on wait()                                  "RunChildThread, lock acquired"
// 6:    Wakes up, tries to acquire mlock but has to wait   calls signal()
// 7:    Waiting for mlock to be free                       "RunChildThread, Waiting"
// 8:    Waiting for mlock to be free                       Calls wait() -releasing mlock
// 9:    Acquires mlock                                     Blocked on wait()
// 10:   "RunMainThread, Running"                           Blocked on wait()
// 11:   calls signal()                                     Wakes up, tries to acquire mlock but has to wait
// 12:   Sleeps for 10 seconds, never releasing mlock       Waiting for mlock to be free
// 13:   Terminates
// Main thread terminating in golang terminates all threads
// Hope this helps!
