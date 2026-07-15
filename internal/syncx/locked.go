package syncx

import "sync"

func Locked(l sync.Locker, fn func()) {
	l.Lock()
	defer l.Unlock()

	fn()
}
