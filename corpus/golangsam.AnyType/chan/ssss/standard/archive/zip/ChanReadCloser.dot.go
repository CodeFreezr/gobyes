// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	zip "archive/zip"
)

// MakeReadCloserChan returns a new open channel
// (simply a 'chan *zip.ReadCloser' that is).
//
// Note: No 'ReadCloser-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadCloserPipelineStartsHere := MakeReadCloserChan()
//	// ... lot's of code to design and build Your favourite "myReadCloserWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadCloserPipelineStartsHere <- drop
//	}
//	close(myReadCloserPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadCloserBuffer) the channel is unbuffered.
//
func MakeReadCloserChan() chan *zip.ReadCloser {
	return make(chan *zip.ReadCloser)
}

// ChanReadCloser returns a channel to receive all inputs before close.
func ChanReadCloser(inp ...*zip.ReadCloser) chan *zip.ReadCloser {
	out := make(chan *zip.ReadCloser)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanReadCloserSlice returns a channel to receive all inputs before close.
func ChanReadCloserSlice(inp ...[]*zip.ReadCloser) chan *zip.ReadCloser {
	out := make(chan *zip.ReadCloser)
	go func() {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}()
	return out
}

// ChanReadCloserFuncNil returns a channel to receive all results of act until nil before close.
func ChanReadCloserFuncNil(act func() *zip.ReadCloser) <-chan *zip.ReadCloser {
	out := make(chan *zip.ReadCloser)
	go func() {
		defer close(out)
		for {
			res := act() // Apply action
			if res == nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanReadCloserFuncNok returns a channel to receive all results of act until nok before close.
func ChanReadCloserFuncNok(act func() (*zip.ReadCloser, bool)) <-chan *zip.ReadCloser {
	out := make(chan *zip.ReadCloser)
	go func() {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanReadCloserFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanReadCloserFuncErr(act func() (*zip.ReadCloser, error)) <-chan *zip.ReadCloser {
	out := make(chan *zip.ReadCloser)
	go func() {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// JoinReadCloser sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadCloser(out chan<- *zip.ReadCloser, inp ...*zip.ReadCloser) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}()
	return done
}

// JoinReadCloserSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadCloserSlice(out chan<- *zip.ReadCloser, inp ...[]*zip.ReadCloser) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}()
	return done
}

// JoinReadCloserChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadCloserChan(out chan<- *zip.ReadCloser, inp <-chan *zip.ReadCloser) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// DoneReadCloser returns a channel to receive one signal before close after inp has been drained.
func DoneReadCloser(inp <-chan *zip.ReadCloser) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}()
	return done
}

// DoneReadCloserSlice returns a channel which will receive a slice
// of all the ReadClosers received on inp channel before close.
// Unlike DoneReadCloser, a full slice is sent once, not just an event.
func DoneReadCloserSlice(inp <-chan *zip.ReadCloser) chan []*zip.ReadCloser {
	done := make(chan []*zip.ReadCloser)
	go func() {
		defer close(done)
		ReadCloserS := []*zip.ReadCloser{}
		for i := range inp {
			ReadCloserS = append(ReadCloserS, i)
		}
		done <- ReadCloserS
	}()
	return done
}

// DoneReadCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadCloserFunc(inp <-chan *zip.ReadCloser, act func(a *zip.ReadCloser)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *zip.ReadCloser) { return }
	}
	go func() {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}()
	return done
}

// PipeReadCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadCloserBuffer(inp <-chan *zip.ReadCloser, cap int) chan *zip.ReadCloser {
	out := make(chan *zip.ReadCloser, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeReadCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadCloserFunc(inp <-chan *zip.ReadCloser, act func(a *zip.ReadCloser) *zip.ReadCloser) chan *zip.ReadCloser {
	out := make(chan *zip.ReadCloser)
	if act == nil {
		act = func(a *zip.ReadCloser) *zip.ReadCloser { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeReadCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadCloserFork(inp <-chan *zip.ReadCloser) (chan *zip.ReadCloser, chan *zip.ReadCloser) {
	out1 := make(chan *zip.ReadCloser)
	out2 := make(chan *zip.ReadCloser)
	go func() {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}()
	return out1, out2
}

// ReadCloserTube is the signature for a pipe function.
type ReadCloserTube func(inp <-chan *zip.ReadCloser, out <-chan *zip.ReadCloser)

// ReadCloserDaisy returns a channel to receive all inp after having passed thru tube.
func ReadCloserDaisy(inp <-chan *zip.ReadCloser, tube ReadCloserTube) (out <-chan *zip.ReadCloser) {
	cha := make(chan *zip.ReadCloser)
	go tube(inp, cha)
	return cha
}

// ReadCloserDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReadCloserDaisyChain(inp <-chan *zip.ReadCloser, tubes ...ReadCloserTube) (out <-chan *zip.ReadCloser) {
	cha := inp
	for i := range tubes {
		cha = ReadCloserDaisy(cha, tubes[i])
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
