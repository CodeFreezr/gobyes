// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dotpath

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/GoLangsam/container/ccsafe/dotpath"
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
func MakeDotPathChan() (out chan *dotpath.DotPath) {
	return make(chan *dotpath.DotPath)
}

func sendDotPath(out chan<- *dotpath.DotPath, inp ...*dotpath.DotPath) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanDotPath returns a channel to receive all inputs before close.
func ChanDotPath(inp ...*dotpath.DotPath) (out <-chan *dotpath.DotPath) {
	cha := make(chan *dotpath.DotPath)
	go sendDotPath(cha, inp...)
	return cha
}

func sendDotPathSlice(out chan<- *dotpath.DotPath, inp ...[]*dotpath.DotPath) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanDotPathSlice returns a channel to receive all inputs before close.
func ChanDotPathSlice(inp ...[]*dotpath.DotPath) (out <-chan *dotpath.DotPath) {
	cha := make(chan *dotpath.DotPath)
	go sendDotPathSlice(cha, inp...)
	return cha
}

func chanDotPathFuncNil(out chan<- *dotpath.DotPath, act func() *dotpath.DotPath) {
	defer close(out)
	for {
		res := act() // Apply action
		if res == nil {
			return
		}
		out <- res
	}
}

// ChanDotPathFuncNil returns a channel to receive all results of act until nil before close.
func ChanDotPathFuncNil(act func() *dotpath.DotPath) (out <-chan *dotpath.DotPath) {
	cha := make(chan *dotpath.DotPath)
	go chanDotPathFuncNil(cha, act)
	return cha
}

func chanDotPathFuncNok(out chan<- *dotpath.DotPath, act func() (*dotpath.DotPath, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanDotPathFuncNok returns a channel to receive all results of act until nok before close.
func ChanDotPathFuncNok(act func() (*dotpath.DotPath, bool)) (out <-chan *dotpath.DotPath) {
	cha := make(chan *dotpath.DotPath)
	go chanDotPathFuncNok(cha, act)
	return cha
}

func chanDotPathFuncErr(out chan<- *dotpath.DotPath, act func() (*dotpath.DotPath, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanDotPathFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanDotPathFuncErr(act func() (*dotpath.DotPath, error)) (out <-chan *dotpath.DotPath) {
	cha := make(chan *dotpath.DotPath)
	go chanDotPathFuncErr(cha, act)
	return cha
}

func joinDotPath(done chan<- struct{}, out chan<- *dotpath.DotPath, inp ...*dotpath.DotPath) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinDotPath sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinDotPath(out chan<- *dotpath.DotPath, inp ...*dotpath.DotPath) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinDotPath(cha, out, inp...)
	return cha
}

func joinDotPathSlice(done chan<- struct{}, out chan<- *dotpath.DotPath, inp ...[]*dotpath.DotPath) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinDotPathSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinDotPathSlice(out chan<- *dotpath.DotPath, inp ...[]*dotpath.DotPath) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinDotPathSlice(cha, out, inp...)
	return cha
}

func joinDotPathChan(done chan<- struct{}, out chan<- *dotpath.DotPath, inp <-chan *dotpath.DotPath) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinDotPathChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinDotPathChan(out chan<- *dotpath.DotPath, inp <-chan *dotpath.DotPath) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinDotPathChan(cha, out, inp)
	return cha
}

func doitDotPath(done chan<- struct{}, inp <-chan *dotpath.DotPath) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneDotPath returns a channel to receive one signal before close after inp has been drained.
func DoneDotPath(inp <-chan *dotpath.DotPath) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitDotPath(cha, inp)
	return cha
}

func doitDotPathSlice(done chan<- ([]*dotpath.DotPath), inp <-chan *dotpath.DotPath) {
	defer close(done)
	DotPathS := []*dotpath.DotPath{}
	for i := range inp {
		DotPathS = append(DotPathS, i)
	}
	done <- DotPathS
}

// DoneDotPathSlice returns a channel which will receive a slice
// of all the DotPaths received on inp channel before close.
// Unlike DoneDotPath, a full slice is sent once, not just an event.
func DoneDotPathSlice(inp <-chan *dotpath.DotPath) (done <-chan ([]*dotpath.DotPath)) {
	cha := make(chan ([]*dotpath.DotPath))
	go doitDotPathSlice(cha, inp)
	return cha
}

func doitDotPathFunc(done chan<- struct{}, inp <-chan *dotpath.DotPath, act func(a *dotpath.DotPath)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneDotPathFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneDotPathFunc(inp <-chan *dotpath.DotPath, act func(a *dotpath.DotPath)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *dotpath.DotPath) { return }
	}
	go doitDotPathFunc(cha, inp, act)
	return cha
}

func pipeDotPathBuffer(out chan<- *dotpath.DotPath, inp <-chan *dotpath.DotPath) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeDotPathBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeDotPathBuffer(inp <-chan *dotpath.DotPath, cap int) (out <-chan *dotpath.DotPath) {
	cha := make(chan *dotpath.DotPath, cap)
	go pipeDotPathBuffer(cha, inp)
	return cha
}

func pipeDotPathFunc(out chan<- *dotpath.DotPath, inp <-chan *dotpath.DotPath, act func(a *dotpath.DotPath) *dotpath.DotPath) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeDotPathFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeDotPathMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeDotPathFunc(inp <-chan *dotpath.DotPath, act func(a *dotpath.DotPath) *dotpath.DotPath) (out <-chan *dotpath.DotPath) {
	cha := make(chan *dotpath.DotPath)
	if act == nil {
		act = func(a *dotpath.DotPath) *dotpath.DotPath { return a }
	}
	go pipeDotPathFunc(cha, inp, act)
	return cha
}

func pipeDotPathFork(out1, out2 chan<- *dotpath.DotPath, inp <-chan *dotpath.DotPath) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeDotPathFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeDotPathFork(inp <-chan *dotpath.DotPath) (out1, out2 <-chan *dotpath.DotPath) {
	cha1 := make(chan *dotpath.DotPath)
	cha2 := make(chan *dotpath.DotPath)
	go pipeDotPathFork(cha1, cha2, inp)
	return cha1, cha2
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
