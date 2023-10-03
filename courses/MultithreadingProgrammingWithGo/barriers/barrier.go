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
	return &Barrier{size, size, locker, condLock}
}

func (b *Barrier) Wai() {
	b.locker.Lock()
	{
		b.count -= 1
		if b.count == 0 {
			b.resetCount()
			b.condLocker.Broadcast()
		}
		if b.count != 0 {
			b.condLocker.Wait()
		}
	}
	b.locker.Unlock()
}

func (b *Barrier) resetCount() {
	b.count = b.total
}
