package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"


func main() {

	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for j := 0; j < 20000; j++ {
				atomic.AddUint64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

