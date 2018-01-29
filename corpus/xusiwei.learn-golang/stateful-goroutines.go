package main

import (
	"fmt"
//	"math/rand"
	"sync/atomic"
//	"time"
)

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {

	const nReaders int = 100
	const nWriters int = 10

	var ops int64 = 0

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	monitorDone := make(chan bool)
	readerDone := make(chan int) // for reader work done notification
	writerDone := make(chan int) // for writer work done notification

	// as a monitor, receives:
	//  all read/write requests
	//  reader/writer work done notification
	// sends:
	//  monitor done after all readers and writers done
	go func() {
		var state = make(map[int]int) // all data now placed here
		var doneReaders int = 0
		var doneWriters int = 0

		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key] // reply requests value
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true // reply write success
			case /*rid :=*/ <-readerDone:
				doneReaders += 1
				if doneReaders == nReaders && doneWriters == nWriters {
					fmt.Println("state:", state)
					monitorDone <- true
					break
				}
			case /*wid :=*/ <-writerDone:
				doneWriters += 1
				if doneReaders == nReaders && doneWriters == nWriters {
					fmt.Println("state:", state)
					monitorDone <- true
					break
				}
			}
		}
	}()


	// readers
	for r := 0; r < nReaders; r++ {
		go func() {
			for t := 0; t < 10000; t++ {
				read := &readOp {
					key: t % 5,  // rand.Intn(5)
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
			readerDone <- r
		}()
	}


	// writers
	for w := 0; w < nWriters; w++ {
		go func() {
			for t := 0; t < 10000; t++ {
				write := &writeOp {
					key: t % 5,   // rand.Intn(5)
					val: t % 100, // rand.Intn(100)
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
			writerDone <- w
		}()
	}


	// wait monitor done
	<-monitorDone

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}
