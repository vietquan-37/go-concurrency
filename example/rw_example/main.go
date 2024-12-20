package main

import "sync"

type ReadWriteMutex struct {
	readerCount int
	readerLock  sync.Mutex
	globalLock  sync.Mutex
}

func (l *ReadWriteMutex) ReaderLock() {
	l.readerLock.Lock()
	l.readerCount++
	if l.readerCount == 1 {
		l.globalLock.Lock()
	}
	l.readerLock.Unlock()
}
func (l *ReadWriteMutex) WriterLock() {
	l.globalLock.Lock()
}
