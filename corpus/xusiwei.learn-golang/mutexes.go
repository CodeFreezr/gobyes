package main

import (
	"fmt"
//	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var ops int64 = 0


	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for t := 0; t < 10000; t++ {
				key := state[t%5]  // rand.Intn(5)

				mutex.Lock()
				total += state[key]
				mutex.Unlock()

				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}


	for w := 0; w < 10; w++ {
		go func() {
			for t := 0; t < 10000; t++ {
				key := t % 5     // rand.Intn(5)
				val := t % 100   // rand.Intn(100)

				mutex.Lock()
				state[key] = val
				mutex.Unlock()

				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
