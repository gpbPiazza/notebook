package main

import (
	"sort"
	"sync"
	"time"
)

var (
	controller            = sync.Mutex{}
	conditionalController = sync.NewCond(&controller)
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

// Hierarchy lock works finding the tow intersections that need to be locked and always start first locking the
// lower intersection ID.
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

// Arbitrator implementations give to a arbitrator the decision if a thread can access the resource that she is asking for use
// if the resource is locked the thread sleep until the resource request its free.
func lockIntersectionsByArbitrator(id, reversedStart, reversedEnd int, crossings []*Crossing) {
	var intersectionsToLock []*Intersection
	for _, c := range crossings {
		if reversedEnd >= c.Position &&
			reversedStart <= c.Position &&
			c.Intersection.LockedBy != id {
			intersectionsToLock = append(intersectionsToLock, c.Intersection)
		}
	}

	controller.Lock()
	for !isAllUnlocked(intersectionsToLock) {
		// put the threads to sleep
		conditionalController.Wait() // remember when a conditionalVariable does Wait to a thread we release the controller.Lock to UnLock
	}

	for _, it := range intersectionsToLock {
		it.LockedBy = id
	}
	controller.Unlock()
}

func isAllUnlocked(intersections []*Intersection) bool {
	for _, it := range intersections {
		if it.LockedBy >= 0 {
			return false
		}
	}
	return true
}

func arbitratorUnlockResource(cross *Crossing) {
	// Arb
	controller.Lock() // by modifiend cross.Intersection.LockedBy we need to ensure that no one i
	// s locking the resource that we are releasing at the same time
	cross.Intersection.LockedBy = -1
	// release the threads from sleep
	conditionalController.Broadcast()
	controller.Unlock()
}

func MoveTrain(train *Train, distance int, crossings []*Crossing) {
	for train.Front < distance {
		train.Front += 1
		for _, cross := range crossings {
			if train.Front == cross.Position {
				reversedStart := cross.Position
				reversedEnd := cross.Position + trainLenght
				lockIntersectionsByArbitrator(train.ID, reversedStart, reversedEnd, crossings)
			}
			back := train.Front - train.Length
			if back == cross.Position {
				// With arbitrator implementation we dont lock the intersections using their own Mutex one by one any more
				// cross.Intersection.Mutex.Unlock()
				// cross.Intersection.LockedBy = -1

				arbitratorUnlockResource(cross)
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
