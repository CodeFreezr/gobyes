// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/tag"
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
func MakeTagChan() chan *tag.TagAny {
	return make(chan *tag.TagAny)
}

// ChanTag returns a channel to receive all inputs before close.
func ChanTag(inp ...*tag.TagAny) chan *tag.TagAny {
	out := make(chan *tag.TagAny)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanTagSlice returns a channel to receive all inputs before close.
func ChanTagSlice(inp ...[]*tag.TagAny) chan *tag.TagAny {
	out := make(chan *tag.TagAny)
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

// ChanTagFuncNil returns a channel to receive all results of act until nil before close.
func ChanTagFuncNil(act func() *tag.TagAny) <-chan *tag.TagAny {
	out := make(chan *tag.TagAny)
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

// ChanTagFuncNok returns a channel to receive all results of act until nok before close.
func ChanTagFuncNok(act func() (*tag.TagAny, bool)) <-chan *tag.TagAny {
	out := make(chan *tag.TagAny)
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

// ChanTagFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanTagFuncErr(act func() (*tag.TagAny, error)) <-chan *tag.TagAny {
	out := make(chan *tag.TagAny)
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

// JoinTag sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinTag(out chan<- *tag.TagAny, inp ...*tag.TagAny) chan struct{} {
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

// JoinTagSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinTagSlice(out chan<- *tag.TagAny, inp ...[]*tag.TagAny) chan struct{} {
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

// JoinTagChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinTagChan(out chan<- *tag.TagAny, inp <-chan *tag.TagAny) chan struct{} {
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

// DoneTag returns a channel to receive one signal before close after inp has been drained.
func DoneTag(inp <-chan *tag.TagAny) chan struct{} {
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

// DoneTagSlice returns a channel which will receive a slice
// of all the Tags received on inp channel before close.
// Unlike DoneTag, a full slice is sent once, not just an event.
func DoneTagSlice(inp <-chan *tag.TagAny) chan []*tag.TagAny {
	done := make(chan []*tag.TagAny)
	go func() {
		defer close(done)
		TagS := []*tag.TagAny{}
		for i := range inp {
			TagS = append(TagS, i)
		}
		done <- TagS
	}()
	return done
}

// DoneTagFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneTagFunc(inp <-chan *tag.TagAny, act func(a *tag.TagAny)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *tag.TagAny) { return }
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

// PipeTagBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeTagBuffer(inp <-chan *tag.TagAny, cap int) chan *tag.TagAny {
	out := make(chan *tag.TagAny, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeTagFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeTagMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeTagFunc(inp <-chan *tag.TagAny, act func(a *tag.TagAny) *tag.TagAny) chan *tag.TagAny {
	out := make(chan *tag.TagAny)
	if act == nil {
		act = func(a *tag.TagAny) *tag.TagAny { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeTagFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeTagFork(inp <-chan *tag.TagAny) (chan *tag.TagAny, chan *tag.TagAny) {
	out1 := make(chan *tag.TagAny)
	out2 := make(chan *tag.TagAny)
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
