// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/GoLangsam/container/ccsafe/dot"
)

// MakeDotChan returns a new open channel
// (simply a 'chan *dot.Dot' that is).
//
// Note: No 'Dot-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myDotPipelineStartsHere := MakeDotChan()
//	// ... lot's of code to design and build Your favourite "myDotWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myDotPipelineStartsHere <- drop
//	}
//	close(myDotPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeDotBuffer) the channel is unbuffered.
//
func MakeDotChan() chan *dot.Dot {
	return make(chan *dot.Dot)
}

// ChanDot returns a channel to receive all inputs before close.
func ChanDot(inp ...*dot.Dot) chan *dot.Dot {
	out := make(chan *dot.Dot)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanDotSlice returns a channel to receive all inputs before close.
func ChanDotSlice(inp ...[]*dot.Dot) chan *dot.Dot {
	out := make(chan *dot.Dot)
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

// ChanDotFuncNil returns a channel to receive all results of act until nil before close.
func ChanDotFuncNil(act func() *dot.Dot) <-chan *dot.Dot {
	out := make(chan *dot.Dot)
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

// ChanDotFuncNok returns a channel to receive all results of act until nok before close.
func ChanDotFuncNok(act func() (*dot.Dot, bool)) <-chan *dot.Dot {
	out := make(chan *dot.Dot)
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

// ChanDotFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanDotFuncErr(act func() (*dot.Dot, error)) <-chan *dot.Dot {
	out := make(chan *dot.Dot)
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

// JoinDot sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinDot(out chan<- *dot.Dot, inp ...*dot.Dot) chan struct{} {
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

// JoinDotSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinDotSlice(out chan<- *dot.Dot, inp ...[]*dot.Dot) chan struct{} {
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

// JoinDotChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinDotChan(out chan<- *dot.Dot, inp <-chan *dot.Dot) chan struct{} {
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

// DoneDot returns a channel to receive one signal before close after inp has been drained.
func DoneDot(inp <-chan *dot.Dot) chan struct{} {
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

// DoneDotSlice returns a channel which will receive a slice
// of all the Dots received on inp channel before close.
// Unlike DoneDot, a full slice is sent once, not just an event.
func DoneDotSlice(inp <-chan *dot.Dot) chan []*dot.Dot {
	done := make(chan []*dot.Dot)
	go func() {
		defer close(done)
		DotS := []*dot.Dot{}
		for i := range inp {
			DotS = append(DotS, i)
		}
		done <- DotS
	}()
	return done
}

// DoneDotFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneDotFunc(inp <-chan *dot.Dot, act func(a *dot.Dot)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *dot.Dot) { return }
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

// PipeDotBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeDotBuffer(inp <-chan *dot.Dot, cap int) chan *dot.Dot {
	out := make(chan *dot.Dot, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeDotFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeDotMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeDotFunc(inp <-chan *dot.Dot, act func(a *dot.Dot) *dot.Dot) chan *dot.Dot {
	out := make(chan *dot.Dot)
	if act == nil {
		act = func(a *dot.Dot) *dot.Dot { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeDotFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeDotFork(inp <-chan *dot.Dot) (chan *dot.Dot, chan *dot.Dot) {
	out1 := make(chan *dot.Dot)
	out2 := make(chan *dot.Dot)
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

// DotTube is the signature for a pipe function.
type DotTube func(inp <-chan *dot.Dot, out <-chan *dot.Dot)

// DotDaisy returns a channel to receive all inp after having passed thru tube.
func DotDaisy(inp <-chan *dot.Dot, tube DotTube) (out <-chan *dot.Dot) {
	cha := make(chan *dot.Dot)
	go tube(inp, cha)
	return cha
}

// DotDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func DotDaisyChain(inp <-chan *dot.Dot, tubes ...DotTube) (out <-chan *dot.Dot) {
	cha := inp
	for i := range tubes {
		cha = DotDaisy(cha, tubes[i])
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
