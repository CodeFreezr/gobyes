// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dotpath

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/dotpath"
)

// MakeDotPathChan returns a new open channel
// (simply a 'chan *dotpath.DotPath' that is).
//
// Note: No 'DotPath-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myDotPathPipelineStartsHere := MakeDotPathChan()
//	// ... lot's of code to design and build Your favourite "myDotPathWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myDotPathPipelineStartsHere <- drop
//	}
//	close(myDotPathPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeDotPathBuffer) the channel is unbuffered.
//
func MakeDotPathChan() chan *dotpath.DotPath {
	return make(chan *dotpath.DotPath)
}

// ChanDotPath returns a channel to receive all inputs before close.
func ChanDotPath(inp ...*dotpath.DotPath) chan *dotpath.DotPath {
	out := make(chan *dotpath.DotPath)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanDotPathSlice returns a channel to receive all inputs before close.
func ChanDotPathSlice(inp ...[]*dotpath.DotPath) chan *dotpath.DotPath {
	out := make(chan *dotpath.DotPath)
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

// ChanDotPathFuncNil returns a channel to receive all results of act until nil before close.
func ChanDotPathFuncNil(act func() *dotpath.DotPath) <-chan *dotpath.DotPath {
	out := make(chan *dotpath.DotPath)
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

// ChanDotPathFuncNok returns a channel to receive all results of act until nok before close.
func ChanDotPathFuncNok(act func() (*dotpath.DotPath, bool)) <-chan *dotpath.DotPath {
	out := make(chan *dotpath.DotPath)
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

// ChanDotPathFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanDotPathFuncErr(act func() (*dotpath.DotPath, error)) <-chan *dotpath.DotPath {
	out := make(chan *dotpath.DotPath)
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

// JoinDotPath sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinDotPath(out chan<- *dotpath.DotPath, inp ...*dotpath.DotPath) chan struct{} {
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

// JoinDotPathSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinDotPathSlice(out chan<- *dotpath.DotPath, inp ...[]*dotpath.DotPath) chan struct{} {
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

// JoinDotPathChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinDotPathChan(out chan<- *dotpath.DotPath, inp <-chan *dotpath.DotPath) chan struct{} {
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

// DoneDotPath returns a channel to receive one signal before close after inp has been drained.
func DoneDotPath(inp <-chan *dotpath.DotPath) chan struct{} {
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

// DoneDotPathSlice returns a channel which will receive a slice
// of all the DotPaths received on inp channel before close.
// Unlike DoneDotPath, a full slice is sent once, not just an event.
func DoneDotPathSlice(inp <-chan *dotpath.DotPath) chan []*dotpath.DotPath {
	done := make(chan []*dotpath.DotPath)
	go func() {
		defer close(done)
		DotPathS := []*dotpath.DotPath{}
		for i := range inp {
			DotPathS = append(DotPathS, i)
		}
		done <- DotPathS
	}()
	return done
}

// DoneDotPathFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneDotPathFunc(inp <-chan *dotpath.DotPath, act func(a *dotpath.DotPath)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *dotpath.DotPath) { return }
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

// PipeDotPathBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeDotPathBuffer(inp <-chan *dotpath.DotPath, cap int) chan *dotpath.DotPath {
	out := make(chan *dotpath.DotPath, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeDotPathFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeDotPathMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeDotPathFunc(inp <-chan *dotpath.DotPath, act func(a *dotpath.DotPath) *dotpath.DotPath) chan *dotpath.DotPath {
	out := make(chan *dotpath.DotPath)
	if act == nil {
		act = func(a *dotpath.DotPath) *dotpath.DotPath { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeDotPathFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeDotPathFork(inp <-chan *dotpath.DotPath) (chan *dotpath.DotPath, chan *dotpath.DotPath) {
	out1 := make(chan *dotpath.DotPath)
	out2 := make(chan *dotpath.DotPath)
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

// DotPathTube is the signature for a pipe function.
type DotPathTube func(inp <-chan *dotpath.DotPath, out <-chan *dotpath.DotPath)

// DotPathDaisy returns a channel to receive all inp after having passed thru tube.
func DotPathDaisy(inp <-chan *dotpath.DotPath, tube DotPathTube) (out <-chan *dotpath.DotPath) {
	cha := make(chan *dotpath.DotPath)
	go tube(inp, cha)
	return cha
}

// DotPathDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func DotPathDaisyChain(inp <-chan *dotpath.DotPath, tubes ...DotPathTube) (out <-chan *dotpath.DotPath) {
	cha := inp
	for i := range tubes {
		cha = DotPathDaisy(cha, tubes[i])
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
