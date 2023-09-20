package main

import (
	"sync"
	"time"
)

type Train struct {
	ID     int
	Length int
	Front  int
}

type Intersection struct {
	ID       int
	Mutex    sync.Mutex
	LockedBy int
}

type Crossing struct {
	Position     int
	Intersection *Intersection
}

func MoveTrain(train *Train, distance int, crossings []*Crossing) {
	for train.Front < distance {
		train.Front += 1
		for _, cross := range crossings {
			if train.Front == cross.Position {
				cross.Intersection.Mutex.Lock()
				cross.Intersection.LockedBy = train.ID
			}
			back := train.Front - train.Length
			if back == cross.Position {
				cross.Intersection.Mutex.Unlock()
				cross.Intersection.LockedBy = -1
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
