// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	tar "archive/tar"
)

// MakeHeaderChan returns a new open channel
// (simply a 'chan *tar.Header' that is).
//
// Note: No 'Header-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myHeaderPipelineStartsHere := MakeHeaderChan()
//	// ... lot's of code to design and build Your favourite "myHeaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myHeaderPipelineStartsHere <- drop
//	}
//	close(myHeaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeHeaderBuffer) the channel is unbuffered.
//
func MakeHeaderChan() chan *tar.Header {
	return make(chan *tar.Header)
}

// ChanHeader returns a channel to receive all inputs before close.
func ChanHeader(inp ...*tar.Header) chan *tar.Header {
	out := make(chan *tar.Header)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanHeaderSlice returns a channel to receive all inputs before close.
func ChanHeaderSlice(inp ...[]*tar.Header) chan *tar.Header {
	out := make(chan *tar.Header)
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

// ChanHeaderFuncNil returns a channel to receive all results of act until nil before close.
func ChanHeaderFuncNil(act func() *tar.Header) <-chan *tar.Header {
	out := make(chan *tar.Header)
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

// ChanHeaderFuncNok returns a channel to receive all results of act until nok before close.
func ChanHeaderFuncNok(act func() (*tar.Header, bool)) <-chan *tar.Header {
	out := make(chan *tar.Header)
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

// ChanHeaderFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanHeaderFuncErr(act func() (*tar.Header, error)) <-chan *tar.Header {
	out := make(chan *tar.Header)
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

// JoinHeader sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinHeader(out chan<- *tar.Header, inp ...*tar.Header) chan struct{} {
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

// JoinHeaderSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinHeaderSlice(out chan<- *tar.Header, inp ...[]*tar.Header) chan struct{} {
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

// JoinHeaderChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinHeaderChan(out chan<- *tar.Header, inp <-chan *tar.Header) chan struct{} {
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

// DoneHeader returns a channel to receive one signal before close after inp has been drained.
func DoneHeader(inp <-chan *tar.Header) chan struct{} {
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

// DoneHeaderSlice returns a channel which will receive a slice
// of all the Headers received on inp channel before close.
// Unlike DoneHeader, a full slice is sent once, not just an event.
func DoneHeaderSlice(inp <-chan *tar.Header) chan []*tar.Header {
	done := make(chan []*tar.Header)
	go func() {
		defer close(done)
		HeaderS := []*tar.Header{}
		for i := range inp {
			HeaderS = append(HeaderS, i)
		}
		done <- HeaderS
	}()
	return done
}

// DoneHeaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneHeaderFunc(inp <-chan *tar.Header, act func(a *tar.Header)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *tar.Header) { return }
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

// PipeHeaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeHeaderBuffer(inp <-chan *tar.Header, cap int) chan *tar.Header {
	out := make(chan *tar.Header, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeHeaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeHeaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeHeaderFunc(inp <-chan *tar.Header, act func(a *tar.Header) *tar.Header) chan *tar.Header {
	out := make(chan *tar.Header)
	if act == nil {
		act = func(a *tar.Header) *tar.Header { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeHeaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeHeaderFork(inp <-chan *tar.Header) (chan *tar.Header, chan *tar.Header) {
	out1 := make(chan *tar.Header)
	out2 := make(chan *tar.Header)
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

// HeaderTube is the signature for a pipe function.
type HeaderTube func(inp <-chan *tar.Header, out <-chan *tar.Header)

// HeaderDaisy returns a channel to receive all inp after having passed thru tube.
func HeaderDaisy(inp <-chan *tar.Header, tube HeaderTube) (out <-chan *tar.Header) {
	cha := make(chan *tar.Header)
	go tube(inp, cha)
	return cha
}

// HeaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func HeaderDaisyChain(inp <-chan *tar.Header, tubes ...HeaderTube) (out <-chan *tar.Header) {
	cha := inp
	for i := range tubes {
		cha = HeaderDaisy(cha, tubes[i])
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
