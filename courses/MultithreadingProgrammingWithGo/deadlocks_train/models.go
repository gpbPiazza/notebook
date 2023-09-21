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

func NewTrain(id int) *Train {
	return &Train{
		ID:     id,
		Length: trainLenght,
		Front:  0,
	}
}

type Intersection struct {
	ID       int
	Mutex    sync.Mutex
	LockedBy int
}

func NewIntersection(id int) *Intersection {
	return &Intersection{
		ID:       id,
		LockedBy: -1,
	}
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
