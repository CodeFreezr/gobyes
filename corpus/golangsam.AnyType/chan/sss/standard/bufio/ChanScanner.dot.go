// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	bufio "bufio"
)

// MakeScannerChan returns a new open channel
// (simply a 'chan *bufio.Scanner' that is).
//
// Note: No 'Scanner-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myScannerPipelineStartsHere := MakeScannerChan()
//	// ... lot's of code to design and build Your favourite "myScannerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myScannerPipelineStartsHere <- drop
//	}
//	close(myScannerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeScannerBuffer) the channel is unbuffered.
//
func MakeScannerChan() (out chan *bufio.Scanner) {
	return make(chan *bufio.Scanner)
}

// ChanScanner returns a channel to receive all inputs before close.
func ChanScanner(inp ...*bufio.Scanner) (out <-chan *bufio.Scanner) {
	cha := make(chan *bufio.Scanner)
	go func(out chan<- *bufio.Scanner, inp ...*bufio.Scanner) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanScannerSlice returns a channel to receive all inputs before close.
func ChanScannerSlice(inp ...[]*bufio.Scanner) (out <-chan *bufio.Scanner) {
	cha := make(chan *bufio.Scanner)
	go func(out chan<- *bufio.Scanner, inp ...[]*bufio.Scanner) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanScannerFuncNok returns a channel to receive all results of act until nok before close.
func ChanScannerFuncNok(act func() (*bufio.Scanner, bool)) (out <-chan *bufio.Scanner) {
	cha := make(chan *bufio.Scanner)
	go func(out chan<- *bufio.Scanner, act func() (*bufio.Scanner, bool)) {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// ChanScannerFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanScannerFuncErr(act func() (*bufio.Scanner, error)) (out <-chan *bufio.Scanner) {
	cha := make(chan *bufio.Scanner)
	go func(out chan<- *bufio.Scanner, act func() (*bufio.Scanner, error)) {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// JoinScanner sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinScanner(out chan<- *bufio.Scanner, inp ...*bufio.Scanner) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *bufio.Scanner, inp ...*bufio.Scanner) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinScannerSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinScannerSlice(out chan<- *bufio.Scanner, inp ...[]*bufio.Scanner) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *bufio.Scanner, inp ...[]*bufio.Scanner) {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinScannerChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinScannerChan(out chan<- *bufio.Scanner, inp <-chan *bufio.Scanner) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *bufio.Scanner, inp <-chan *bufio.Scanner) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneScanner returns a channel to receive one signal before close after inp has been drained.
func DoneScanner(inp <-chan *bufio.Scanner) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *bufio.Scanner) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneScannerSlice returns a channel which will receive a slice
// of all the Scanners received on inp channel before close.
// Unlike DoneScanner, a full slice is sent once, not just an event.
func DoneScannerSlice(inp <-chan *bufio.Scanner) (done <-chan []*bufio.Scanner) {
	cha := make(chan []*bufio.Scanner)
	go func(inp <-chan *bufio.Scanner, done chan<- []*bufio.Scanner) {
		defer close(done)
		ScannerS := []*bufio.Scanner{}
		for i := range inp {
			ScannerS = append(ScannerS, i)
		}
		done <- ScannerS
	}(inp, cha)
	return cha
}

// DoneScannerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneScannerFunc(inp <-chan *bufio.Scanner, act func(a *bufio.Scanner)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *bufio.Scanner) { return }
	}
	go func(done chan<- struct{}, inp <-chan *bufio.Scanner, act func(a *bufio.Scanner)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeScannerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeScannerBuffer(inp <-chan *bufio.Scanner, cap int) (out <-chan *bufio.Scanner) {
	cha := make(chan *bufio.Scanner, cap)
	go func(out chan<- *bufio.Scanner, inp <-chan *bufio.Scanner) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeScannerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeScannerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeScannerFunc(inp <-chan *bufio.Scanner, act func(a *bufio.Scanner) *bufio.Scanner) (out <-chan *bufio.Scanner) {
	cha := make(chan *bufio.Scanner)
	if act == nil {
		act = func(a *bufio.Scanner) *bufio.Scanner { return a }
	}
	go func(out chan<- *bufio.Scanner, inp <-chan *bufio.Scanner, act func(a *bufio.Scanner) *bufio.Scanner) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeScannerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeScannerFork(inp <-chan *bufio.Scanner) (out1, out2 <-chan *bufio.Scanner) {
	cha1 := make(chan *bufio.Scanner)
	cha2 := make(chan *bufio.Scanner)
	go func(out1, out2 chan<- *bufio.Scanner, inp <-chan *bufio.Scanner) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// ScannerTube is the signature for a pipe function.
type ScannerTube func(inp <-chan *bufio.Scanner, out <-chan *bufio.Scanner)

// ScannerDaisy returns a channel to receive all inp after having passed thru tube.
func ScannerDaisy(inp <-chan *bufio.Scanner, tube ScannerTube) (out <-chan *bufio.Scanner) {
	cha := make(chan *bufio.Scanner)
	go tube(inp, cha)
	return cha
}

// ScannerDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ScannerDaisyChain(inp <-chan *bufio.Scanner, tubes ...ScannerTube) (out <-chan *bufio.Scanner) {
	cha := inp
	for i := range tubes {
		cha = ScannerDaisy(cha, tubes[i])
	}
	return cha
}

/*
func sendOneInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
}

func sendTwoInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
	snd <- 2 // send a 2
}

var fun = func(left chan<- int, right <-chan int) { left <- 1 + <-right }

func main() {
	leftmost := make(chan int)
	right := daisyChain(leftmost, fun, 10000) // the chain - right to left!
	go sendTwoInto(right)
	fmt.Println(<-leftmost)
}
*/
