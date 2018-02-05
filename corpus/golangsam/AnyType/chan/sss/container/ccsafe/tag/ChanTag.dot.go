// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/GoLangsam/container/ccsafe/tag"
)

// MakeTagChan returns a new open channel
// (simply a 'chan *tag.TagAny' that is).
//
// Note: No 'Tag-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myTagPipelineStartsHere := MakeTagChan()
//	// ... lot's of code to design and build Your favourite "myTagWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myTagPipelineStartsHere <- drop
//	}
//	close(myTagPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeTagBuffer) the channel is unbuffered.
//
func MakeTagChan() (out chan *tag.TagAny) {
	return make(chan *tag.TagAny)
}

// ChanTag returns a channel to receive all inputs before close.
func ChanTag(inp ...*tag.TagAny) (out <-chan *tag.TagAny) {
	cha := make(chan *tag.TagAny)
	go func(out chan<- *tag.TagAny, inp ...*tag.TagAny) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanTagSlice returns a channel to receive all inputs before close.
func ChanTagSlice(inp ...[]*tag.TagAny) (out <-chan *tag.TagAny) {
	cha := make(chan *tag.TagAny)
	go func(out chan<- *tag.TagAny, inp ...[]*tag.TagAny) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanTagFuncNil returns a channel to receive all results of act until nil before close.
func ChanTagFuncNil(act func() *tag.TagAny) (out <-chan *tag.TagAny) {
	cha := make(chan *tag.TagAny)
	go func(out chan<- *tag.TagAny, act func() *tag.TagAny) {
		defer close(out)
		for {
			res := act() // Apply action
			if res == nil {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// ChanTagFuncNok returns a channel to receive all results of act until nok before close.
func ChanTagFuncNok(act func() (*tag.TagAny, bool)) (out <-chan *tag.TagAny) {
	cha := make(chan *tag.TagAny)
	go func(out chan<- *tag.TagAny, act func() (*tag.TagAny, bool)) {
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

// ChanTagFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanTagFuncErr(act func() (*tag.TagAny, error)) (out <-chan *tag.TagAny) {
	cha := make(chan *tag.TagAny)
	go func(out chan<- *tag.TagAny, act func() (*tag.TagAny, error)) {
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

// JoinTag sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinTag(out chan<- *tag.TagAny, inp ...*tag.TagAny) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *tag.TagAny, inp ...*tag.TagAny) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinTagSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinTagSlice(out chan<- *tag.TagAny, inp ...[]*tag.TagAny) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *tag.TagAny, inp ...[]*tag.TagAny) {
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

// JoinTagChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinTagChan(out chan<- *tag.TagAny, inp <-chan *tag.TagAny) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *tag.TagAny, inp <-chan *tag.TagAny) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneTag returns a channel to receive one signal before close after inp has been drained.
func DoneTag(inp <-chan *tag.TagAny) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *tag.TagAny) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneTagSlice returns a channel which will receive a slice
// of all the Tags received on inp channel before close.
// Unlike DoneTag, a full slice is sent once, not just an event.
func DoneTagSlice(inp <-chan *tag.TagAny) (done <-chan []*tag.TagAny) {
	cha := make(chan []*tag.TagAny)
	go func(inp <-chan *tag.TagAny, done chan<- []*tag.TagAny) {
		defer close(done)
		TagS := []*tag.TagAny{}
		for i := range inp {
			TagS = append(TagS, i)
		}
		done <- TagS
	}(inp, cha)
	return cha
}

// DoneTagFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneTagFunc(inp <-chan *tag.TagAny, act func(a *tag.TagAny)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *tag.TagAny) { return }
	}
	go func(done chan<- struct{}, inp <-chan *tag.TagAny, act func(a *tag.TagAny)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeTagBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeTagBuffer(inp <-chan *tag.TagAny, cap int) (out <-chan *tag.TagAny) {
	cha := make(chan *tag.TagAny, cap)
	go func(out chan<- *tag.TagAny, inp <-chan *tag.TagAny) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeTagFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeTagMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeTagFunc(inp <-chan *tag.TagAny, act func(a *tag.TagAny) *tag.TagAny) (out <-chan *tag.TagAny) {
	cha := make(chan *tag.TagAny)
	if act == nil {
		act = func(a *tag.TagAny) *tag.TagAny { return a }
	}
	go func(out chan<- *tag.TagAny, inp <-chan *tag.TagAny, act func(a *tag.TagAny) *tag.TagAny) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeTagFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeTagFork(inp <-chan *tag.TagAny) (out1, out2 <-chan *tag.TagAny) {
	cha1 := make(chan *tag.TagAny)
	cha2 := make(chan *tag.TagAny)
	go func(out1, out2 chan<- *tag.TagAny, inp <-chan *tag.TagAny) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// TagTube is the signature for a pipe function.
type TagTube func(inp <-chan *tag.TagAny, out <-chan *tag.TagAny)

// TagDaisy returns a channel to receive all inp after having passed thru tube.
func TagDaisy(inp <-chan *tag.TagAny, tube TagTube) (out <-chan *tag.TagAny) {
	cha := make(chan *tag.TagAny)
	go tube(inp, cha)
	return cha
}

// TagDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func TagDaisyChain(inp <-chan *tag.TagAny, tubes ...TagTube) (out <-chan *tag.TagAny) {
	cha := inp
	for i := range tubes {
		cha = TagDaisy(cha, tubes[i])
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
