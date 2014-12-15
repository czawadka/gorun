package sync

import (
	"sync"
)

type Latch interface {
	Release()
	Await()
}

type rwMutexLatch struct {
	rwMutex sync.RWMutex
}

func (l *rwMutexLatch) Release() {
	l.rwMutex.Unlock()
}

func (l *rwMutexLatch) Await() {
	l.rwMutex.RLock()
}

func NewLatch() Latch {
	l := rwMutexLatch{}
	l.rwMutex.Lock()
	return &l
}
