// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"strings"
)

// MakeReplacerChan returns a new open channel
// (simply a 'chan *strings.Replacer' that is).
//
// Note: No 'Replacer-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReplacerPipelineStartsHere := MakeReplacerChan()
//	// ... lot's of code to design and build Your favourite "myReplacerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReplacerPipelineStartsHere <- drop
//	}
//	close(myReplacerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReplacerBuffer) the channel is unbuffered.
//
func MakeReplacerChan() (out chan *strings.Replacer) {
	return make(chan *strings.Replacer)
}

// ChanReplacer returns a channel to receive all inputs before close.
func ChanReplacer(inp ...*strings.Replacer) (out <-chan *strings.Replacer) {
	cha := make(chan *strings.Replacer)
	go func(out chan<- *strings.Replacer, inp ...*strings.Replacer) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanReplacerSlice returns a channel to receive all inputs before close.
func ChanReplacerSlice(inp ...[]*strings.Replacer) (out <-chan *strings.Replacer) {
	cha := make(chan *strings.Replacer)
	go func(out chan<- *strings.Replacer, inp ...[]*strings.Replacer) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanReplacerFuncNok returns a channel to receive all results of act until nok before close.
func ChanReplacerFuncNok(act func() (*strings.Replacer, bool)) (out <-chan *strings.Replacer) {
	cha := make(chan *strings.Replacer)
	go func(out chan<- *strings.Replacer, act func() (*strings.Replacer, bool)) {
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

// ChanReplacerFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanReplacerFuncErr(act func() (*strings.Replacer, error)) (out <-chan *strings.Replacer) {
	cha := make(chan *strings.Replacer)
	go func(out chan<- *strings.Replacer, act func() (*strings.Replacer, error)) {
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

// JoinReplacer sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReplacer(out chan<- *strings.Replacer, inp ...*strings.Replacer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *strings.Replacer, inp ...*strings.Replacer) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinReplacerSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReplacerSlice(out chan<- *strings.Replacer, inp ...[]*strings.Replacer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *strings.Replacer, inp ...[]*strings.Replacer) {
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

// JoinReplacerChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReplacerChan(out chan<- *strings.Replacer, inp <-chan *strings.Replacer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *strings.Replacer, inp <-chan *strings.Replacer) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneReplacer returns a channel to receive one signal before close after inp has been drained.
func DoneReplacer(inp <-chan *strings.Replacer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *strings.Replacer) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneReplacerSlice returns a channel which will receive a slice
// of all the Replacers received on inp channel before close.
// Unlike DoneReplacer, a full slice is sent once, not just an event.
func DoneReplacerSlice(inp <-chan *strings.Replacer) (done <-chan []*strings.Replacer) {
	cha := make(chan []*strings.Replacer)
	go func(inp <-chan *strings.Replacer, done chan<- []*strings.Replacer) {
		defer close(done)
		ReplacerS := []*strings.Replacer{}
		for i := range inp {
			ReplacerS = append(ReplacerS, i)
		}
		done <- ReplacerS
	}(inp, cha)
	return cha
}

// DoneReplacerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReplacerFunc(inp <-chan *strings.Replacer, act func(a *strings.Replacer)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *strings.Replacer) { return }
	}
	go func(done chan<- struct{}, inp <-chan *strings.Replacer, act func(a *strings.Replacer)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeReplacerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReplacerBuffer(inp <-chan *strings.Replacer, cap int) (out <-chan *strings.Replacer) {
	cha := make(chan *strings.Replacer, cap)
	go func(out chan<- *strings.Replacer, inp <-chan *strings.Replacer) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeReplacerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReplacerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReplacerFunc(inp <-chan *strings.Replacer, act func(a *strings.Replacer) *strings.Replacer) (out <-chan *strings.Replacer) {
	cha := make(chan *strings.Replacer)
	if act == nil {
		act = func(a *strings.Replacer) *strings.Replacer { return a }
	}
	go func(out chan<- *strings.Replacer, inp <-chan *strings.Replacer, act func(a *strings.Replacer) *strings.Replacer) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeReplacerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReplacerFork(inp <-chan *strings.Replacer) (out1, out2 <-chan *strings.Replacer) {
	cha1 := make(chan *strings.Replacer)
	cha2 := make(chan *strings.Replacer)
	go func(out1, out2 chan<- *strings.Replacer, inp <-chan *strings.Replacer) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// ReplacerTube is the signature for a pipe function.
type ReplacerTube func(inp <-chan *strings.Replacer, out <-chan *strings.Replacer)

// ReplacerDaisy returns a channel to receive all inp after having passed thru tube.
func ReplacerDaisy(inp <-chan *strings.Replacer, tube ReplacerTube) (out <-chan *strings.Replacer) {
	cha := make(chan *strings.Replacer)
	go tube(inp, cha)
	return cha
}

// ReplacerDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReplacerDaisyChain(inp <-chan *strings.Replacer, tubes ...ReplacerTube) (out <-chan *strings.Replacer) {
	cha := inp
	for i := range tubes {
		cha = ReplacerDaisy(cha, tubes[i])
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
