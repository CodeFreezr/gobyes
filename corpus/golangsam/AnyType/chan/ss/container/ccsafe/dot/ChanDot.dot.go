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
func MakeDotChan() (out chan *dot.Dot) {
	return make(chan *dot.Dot)
}

func sendDot(out chan<- *dot.Dot, inp ...*dot.Dot) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanDot returns a channel to receive all inputs before close.
func ChanDot(inp ...*dot.Dot) (out <-chan *dot.Dot) {
	cha := make(chan *dot.Dot)
	go sendDot(cha, inp...)
	return cha
}

func sendDotSlice(out chan<- *dot.Dot, inp ...[]*dot.Dot) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanDotSlice returns a channel to receive all inputs before close.
func ChanDotSlice(inp ...[]*dot.Dot) (out <-chan *dot.Dot) {
	cha := make(chan *dot.Dot)
	go sendDotSlice(cha, inp...)
	return cha
}

func chanDotFuncNil(out chan<- *dot.Dot, act func() *dot.Dot) {
	defer close(out)
	for {
		res := act() // Apply action
		if res == nil {
			return
		}
		out <- res
	}
}

// ChanDotFuncNil returns a channel to receive all results of act until nil before close.
func ChanDotFuncNil(act func() *dot.Dot) (out <-chan *dot.Dot) {
	cha := make(chan *dot.Dot)
	go chanDotFuncNil(cha, act)
	return cha
}

func chanDotFuncNok(out chan<- *dot.Dot, act func() (*dot.Dot, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanDotFuncNok returns a channel to receive all results of act until nok before close.
func ChanDotFuncNok(act func() (*dot.Dot, bool)) (out <-chan *dot.Dot) {
	cha := make(chan *dot.Dot)
	go chanDotFuncNok(cha, act)
	return cha
}

func chanDotFuncErr(out chan<- *dot.Dot, act func() (*dot.Dot, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanDotFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanDotFuncErr(act func() (*dot.Dot, error)) (out <-chan *dot.Dot) {
	cha := make(chan *dot.Dot)
	go chanDotFuncErr(cha, act)
	return cha
}

func joinDot(done chan<- struct{}, out chan<- *dot.Dot, inp ...*dot.Dot) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinDot sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinDot(out chan<- *dot.Dot, inp ...*dot.Dot) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinDot(cha, out, inp...)
	return cha
}

func joinDotSlice(done chan<- struct{}, out chan<- *dot.Dot, inp ...[]*dot.Dot) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinDotSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinDotSlice(out chan<- *dot.Dot, inp ...[]*dot.Dot) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinDotSlice(cha, out, inp...)
	return cha
}

func joinDotChan(done chan<- struct{}, out chan<- *dot.Dot, inp <-chan *dot.Dot) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinDotChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinDotChan(out chan<- *dot.Dot, inp <-chan *dot.Dot) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinDotChan(cha, out, inp)
	return cha
}

func doitDot(done chan<- struct{}, inp <-chan *dot.Dot) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneDot returns a channel to receive one signal before close after inp has been drained.
func DoneDot(inp <-chan *dot.Dot) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitDot(cha, inp)
	return cha
}

func doitDotSlice(done chan<- ([]*dot.Dot), inp <-chan *dot.Dot) {
	defer close(done)
	DotS := []*dot.Dot{}
	for i := range inp {
		DotS = append(DotS, i)
	}
	done <- DotS
}

// DoneDotSlice returns a channel which will receive a slice
// of all the Dots received on inp channel before close.
// Unlike DoneDot, a full slice is sent once, not just an event.
func DoneDotSlice(inp <-chan *dot.Dot) (done <-chan ([]*dot.Dot)) {
	cha := make(chan ([]*dot.Dot))
	go doitDotSlice(cha, inp)
	return cha
}

func doitDotFunc(done chan<- struct{}, inp <-chan *dot.Dot, act func(a *dot.Dot)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneDotFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneDotFunc(inp <-chan *dot.Dot, act func(a *dot.Dot)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *dot.Dot) { return }
	}
	go doitDotFunc(cha, inp, act)
	return cha
}

func pipeDotBuffer(out chan<- *dot.Dot, inp <-chan *dot.Dot) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeDotBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeDotBuffer(inp <-chan *dot.Dot, cap int) (out <-chan *dot.Dot) {
	cha := make(chan *dot.Dot, cap)
	go pipeDotBuffer(cha, inp)
	return cha
}

func pipeDotFunc(out chan<- *dot.Dot, inp <-chan *dot.Dot, act func(a *dot.Dot) *dot.Dot) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeDotFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeDotMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeDotFunc(inp <-chan *dot.Dot, act func(a *dot.Dot) *dot.Dot) (out <-chan *dot.Dot) {
	cha := make(chan *dot.Dot)
	if act == nil {
		act = func(a *dot.Dot) *dot.Dot { return a }
	}
	go pipeDotFunc(cha, inp, act)
	return cha
}

func pipeDotFork(out1, out2 chan<- *dot.Dot, inp <-chan *dot.Dot) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeDotFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeDotFork(inp <-chan *dot.Dot) (out1, out2 <-chan *dot.Dot) {
	cha1 := make(chan *dot.Dot)
	cha2 := make(chan *dot.Dot)
	go pipeDotFork(cha1, cha2, inp)
	return cha1, cha2
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
