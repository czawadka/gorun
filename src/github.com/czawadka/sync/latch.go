package sync

import (
	"sync"
	"sync/atomic"
)

type CountDownLatch interface {
	CountDown()
	Await()
}

type rwMutexCountDownLatch struct {
	mutext sync.RWMutex
	count int32
}

func (l *rwMutexCountDownLatch) CountDown() {
	addResult := atomic.AddInt32(&l.count, -1)
	if addResult == 0 {
		l.mutext.Unlock() // release write lock which hold all routines awaiting on read lock
	}
}

func (l *rwMutexCountDownLatch) Await() {
	l.mutext.RLock()
}

func NewCountDownLatch(initialCount int32) CountDownLatch {
	l := rwMutexCountDownLatch{}
	l.mutext.Lock()
	l.count = initialCount
	return &l
}
