package main

import "sync"

type Barrier struct {
	total      int
	count      int
	locker     *sync.Mutex
	condLocker *sync.Cond
}

func NewBarrier(size int) *Barrier {
	locker := &sync.Mutex{}
	condLock := sync.NewCond(locker)
	return &Barrier{
		count:      size,
		total:      size,
		locker:     locker,
		condLocker: condLock,
	}
}

func (b *Barrier) Wait() {
	b.locker.Lock()
	{
		b.count -= 1
		if b.count == 0 {
			b.count = b.total
			b.condLocker.Broadcast()
		} else {
			b.condLocker.Wait()
		}
	}
	b.locker.Unlock()
}
