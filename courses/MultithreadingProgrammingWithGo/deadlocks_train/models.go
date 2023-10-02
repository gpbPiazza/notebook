package main

import (
	"sort"
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
	PriorityLock int
	Position     int
	Intersection *Intersection
}

func lockIntersectionsByPriority(id, reversedStart, reversedEnd int, crossings []*Crossing) {
	var intersectionsToLock []*Intersection
	for _, c := range crossings {
		if reversedEnd >= c.Position &&
			reversedStart <= c.Position &&
			c.Intersection.LockedBy != id {
			intersectionsToLock = append(intersectionsToLock, c.Intersection)
		}
	}

	sort.Slice(intersectionsToLock, func(i, j int) bool {
		return intersectionsToLock[i].ID > intersectionsToLock[j].ID
	})

	for _, it := range intersectionsToLock {
		it.Mutex.Lock()
		it.LockedBy = id
	}
}

func MoveTrain(train *Train, distance int, crossings []*Crossing) {
	for train.Front < distance {
		train.Front += 1
		for _, cross := range crossings {
			if train.Front == cross.Position {
				reversedStart := cross.Position
				reversedEnd := cross.Position + trainLenght
				lockIntersectionsByPriority(train.ID, reversedStart, reversedEnd, crossings)
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
