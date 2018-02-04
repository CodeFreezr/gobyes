package main
import "sync"

func callLocked(lock *sync.Mutex, f func()) {
	lock.Lock()
	defer lock.Unlock()
	f()
}


func main() {
	var broken func()
	var lock sync.Mutex
	defer func() { recover() }()
	callLocked(&lock, broken)
}
