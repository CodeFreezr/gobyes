// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	list "container/list"
)

// MakeElementSChan returns a new open channel
// (simply a 'chan []*list.Element' that is).
//
// Note: No 'ElementS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myElementSPipelineStartsHere := MakeElementSChan()
//	// ... lot's of code to design and build Your favourite "myElementSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myElementSPipelineStartsHere <- drop
//	}
//	close(myElementSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeElementSBuffer) the channel is unbuffered.
//
func MakeElementSChan() chan []*list.Element {
	return make(chan []*list.Element)
}

// ChanElementS returns a channel to receive all inputs before close.
func ChanElementS(inp ...[]*list.Element) chan []*list.Element {
	out := make(chan []*list.Element)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanElementSSlice returns a channel to receive all inputs before close.
func ChanElementSSlice(inp ...[][]*list.Element) chan []*list.Element {
	out := make(chan []*list.Element)
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

// ChanElementSFuncNil returns a channel to receive all results of act until nil before close.
func ChanElementSFuncNil(act func() []*list.Element) <-chan []*list.Element {
	out := make(chan []*list.Element)
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

// ChanElementSFuncNok returns a channel to receive all results of act until nok before close.
func ChanElementSFuncNok(act func() ([]*list.Element, bool)) <-chan []*list.Element {
	out := make(chan []*list.Element)
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

// ChanElementSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanElementSFuncErr(act func() ([]*list.Element, error)) <-chan []*list.Element {
	out := make(chan []*list.Element)
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

// JoinElementS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinElementS(out chan<- []*list.Element, inp ...[]*list.Element) chan struct{} {
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

// JoinElementSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinElementSSlice(out chan<- []*list.Element, inp ...[][]*list.Element) chan struct{} {
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

// JoinElementSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinElementSChan(out chan<- []*list.Element, inp <-chan []*list.Element) chan struct{} {
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

// DoneElementS returns a channel to receive one signal before close after inp has been drained.
func DoneElementS(inp <-chan []*list.Element) chan struct{} {
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

// DoneElementSSlice returns a channel which will receive a slice
// of all the ElementSs received on inp channel before close.
// Unlike DoneElementS, a full slice is sent once, not just an event.
func DoneElementSSlice(inp <-chan []*list.Element) chan [][]*list.Element {
	done := make(chan [][]*list.Element)
	go func() {
		defer close(done)
		ElementSS := [][]*list.Element{}
		for i := range inp {
			ElementSS = append(ElementSS, i)
		}
		done <- ElementSS
	}()
	return done
}

// DoneElementSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneElementSFunc(inp <-chan []*list.Element, act func(a []*list.Element)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a []*list.Element) { return }
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

// PipeElementSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeElementSBuffer(inp <-chan []*list.Element, cap int) chan []*list.Element {
	out := make(chan []*list.Element, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeElementSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeElementSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeElementSFunc(inp <-chan []*list.Element, act func(a []*list.Element) []*list.Element) chan []*list.Element {
	out := make(chan []*list.Element)
	if act == nil {
		act = func(a []*list.Element) []*list.Element { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeElementSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeElementSFork(inp <-chan []*list.Element) (chan []*list.Element, chan []*list.Element) {
	out1 := make(chan []*list.Element)
	out2 := make(chan []*list.Element)
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

// ElementSTube is the signature for a pipe function.
type ElementSTube func(inp <-chan []*list.Element, out <-chan []*list.Element)

// ElementSDaisy returns a channel to receive all inp after having passed thru tube.
func ElementSDaisy(inp <-chan []*list.Element, tube ElementSTube) (out <-chan []*list.Element) {
	cha := make(chan []*list.Element)
	go tube(inp, cha)
	return cha
}

// ElementSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ElementSDaisyChain(inp <-chan []*list.Element, tubes ...ElementSTube) (out <-chan []*list.Element) {
	cha := inp
	for i := range tubes {
		cha = ElementSDaisy(cha, tubes[i])
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
