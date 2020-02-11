package synch_lock

import (
	"strings"
	"sync"
)

type KeyLockPool struct {
	mgLock sync.RWMutex
	locks  []*Lock
}

func (l *KeyLockPool) GetLock(key string) bool {

	l.mgLock.Lock()
	defer l.mgLock.Unlock()

	// 是否可以拿到锁
	for _, lock := range l.locks {
		if lock.IsBlock(key) {
			return false
		}
	}

	// 申请一个锁
	for _, lock := range l.locks {
		if lock.IsFree() {
			lock.Take(key)
			return true
		}
	}

	// 锁被用完了
	return false
}

func (l *KeyLockPool) ReleaseLock(key string) {

	l.mgLock.Lock()
	defer l.mgLock.Unlock()

	for _, lock := range l.locks {
		lock.Release(key)
	}
}

func (l *KeyLockPool) IsAllFree() bool {

	l.mgLock.Lock()
	defer l.mgLock.Unlock()

	for _, lock := range l.locks {
		if !lock.IsFree() {
			return false
		}
	}
	return true
}

func (l *KeyLockPool) GetKeysStatus() string {
	var (
		keys = make([]string, len(l.locks))
	)

	l.mgLock.Lock()
	defer l.mgLock.Unlock()

	for i, lock := range l.locks {
		keys[i] = lock.Key
	}
	return strings.Join(keys, ",")
}

func NewKeyLockPool(size int) *KeyLockPool {

	var (
		l = KeyLockPool{
			locks: make([]*Lock, size),
		}
	)

	for i := 0; i < size; i++ {
		l.locks[i] = &Lock{}
	}

	return &l
}
