// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	list "container/list"
)

// MakeListSChan returns a new open channel
// (simply a 'chan []*list.List' that is).
//
// Note: No 'ListS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myListSPipelineStartsHere := MakeListSChan()
//	// ... lot's of code to design and build Your favourite "myListSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myListSPipelineStartsHere <- drop
//	}
//	close(myListSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeListSBuffer) the channel is unbuffered.
//
func MakeListSChan() chan []*list.List {
	return make(chan []*list.List)
}

// ChanListS returns a channel to receive all inputs before close.
func ChanListS(inp ...[]*list.List) chan []*list.List {
	out := make(chan []*list.List)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanListSSlice returns a channel to receive all inputs before close.
func ChanListSSlice(inp ...[][]*list.List) chan []*list.List {
	out := make(chan []*list.List)
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

// ChanListSFuncNil returns a channel to receive all results of act until nil before close.
func ChanListSFuncNil(act func() []*list.List) <-chan []*list.List {
	out := make(chan []*list.List)
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

// ChanListSFuncNok returns a channel to receive all results of act until nok before close.
func ChanListSFuncNok(act func() ([]*list.List, bool)) <-chan []*list.List {
	out := make(chan []*list.List)
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

// ChanListSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanListSFuncErr(act func() ([]*list.List, error)) <-chan []*list.List {
	out := make(chan []*list.List)
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

// JoinListS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinListS(out chan<- []*list.List, inp ...[]*list.List) chan struct{} {
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

// JoinListSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinListSSlice(out chan<- []*list.List, inp ...[][]*list.List) chan struct{} {
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

// JoinListSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinListSChan(out chan<- []*list.List, inp <-chan []*list.List) chan struct{} {
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

// DoneListS returns a channel to receive one signal before close after inp has been drained.
func DoneListS(inp <-chan []*list.List) chan struct{} {
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

// DoneListSSlice returns a channel which will receive a slice
// of all the ListSs received on inp channel before close.
// Unlike DoneListS, a full slice is sent once, not just an event.
func DoneListSSlice(inp <-chan []*list.List) chan [][]*list.List {
	done := make(chan [][]*list.List)
	go func() {
		defer close(done)
		ListSS := [][]*list.List{}
		for i := range inp {
			ListSS = append(ListSS, i)
		}
		done <- ListSS
	}()
	return done
}

// DoneListSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneListSFunc(inp <-chan []*list.List, act func(a []*list.List)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a []*list.List) { return }
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

// PipeListSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeListSBuffer(inp <-chan []*list.List, cap int) chan []*list.List {
	out := make(chan []*list.List, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeListSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeListSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeListSFunc(inp <-chan []*list.List, act func(a []*list.List) []*list.List) chan []*list.List {
	out := make(chan []*list.List)
	if act == nil {
		act = func(a []*list.List) []*list.List { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeListSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeListSFork(inp <-chan []*list.List) (chan []*list.List, chan []*list.List) {
	out1 := make(chan []*list.List)
	out2 := make(chan []*list.List)
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

// ListSTube is the signature for a pipe function.
type ListSTube func(inp <-chan []*list.List, out <-chan []*list.List)

// ListSDaisy returns a channel to receive all inp after having passed thru tube.
func ListSDaisy(inp <-chan []*list.List, tube ListSTube) (out <-chan []*list.List) {
	cha := make(chan []*list.List)
	go tube(inp, cha)
	return cha
}

// ListSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ListSDaisyChain(inp <-chan []*list.List, tubes ...ListSTube) (out <-chan []*list.List) {
	cha := inp
	for i := range tubes {
		cha = ListSDaisy(cha, tubes[i])
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
