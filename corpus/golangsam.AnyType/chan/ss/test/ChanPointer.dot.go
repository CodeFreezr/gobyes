// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakePointerChan returns a new open channel
// (simply a 'chan *SomeType' that is).
//
// Note: No 'Pointer-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPointerPipelineStartsHere := MakePointerChan()
//	// ... lot's of code to design and build Your favourite "myPointerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPointerPipelineStartsHere <- drop
//	}
//	close(myPointerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePointerBuffer) the channel is unbuffered.
//
func MakePointerChan() (out chan *SomeType) {
	return make(chan *SomeType)
}

func sendPointer(out chan<- *SomeType, inp ...*SomeType) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanPointer returns a channel to receive all inputs before close.
func ChanPointer(inp ...*SomeType) (out <-chan *SomeType) {
	cha := make(chan *SomeType)
	go sendPointer(cha, inp...)
	return cha
}

func sendPointerSlice(out chan<- *SomeType, inp ...[]*SomeType) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanPointerSlice returns a channel to receive all inputs before close.
func ChanPointerSlice(inp ...[]*SomeType) (out <-chan *SomeType) {
	cha := make(chan *SomeType)
	go sendPointerSlice(cha, inp...)
	return cha
}

func chanPointerFuncNok(out chan<- *SomeType, act func() (*SomeType, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanPointerFuncNok returns a channel to receive all results of act until nok before close.
func ChanPointerFuncNok(act func() (*SomeType, bool)) (out <-chan *SomeType) {
	cha := make(chan *SomeType)
	go chanPointerFuncNok(cha, act)
	return cha
}

func chanPointerFuncErr(out chan<- *SomeType, act func() (*SomeType, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanPointerFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanPointerFuncErr(act func() (*SomeType, error)) (out <-chan *SomeType) {
	cha := make(chan *SomeType)
	go chanPointerFuncErr(cha, act)
	return cha
}

func joinPointer(done chan<- struct{}, out chan<- *SomeType, inp ...*SomeType) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinPointer sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPointer(out chan<- *SomeType, inp ...*SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPointer(cha, out, inp...)
	return cha
}

func joinPointerSlice(done chan<- struct{}, out chan<- *SomeType, inp ...[]*SomeType) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinPointerSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPointerSlice(out chan<- *SomeType, inp ...[]*SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPointerSlice(cha, out, inp...)
	return cha
}

func joinPointerChan(done chan<- struct{}, out chan<- *SomeType, inp <-chan *SomeType) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinPointerChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPointerChan(out chan<- *SomeType, inp <-chan *SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPointerChan(cha, out, inp)
	return cha
}

func doitPointer(done chan<- struct{}, inp <-chan *SomeType) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DonePointer returns a channel to receive one signal before close after inp has been drained.
func DonePointer(inp <-chan *SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitPointer(cha, inp)
	return cha
}

func doitPointerSlice(done chan<- ([]*SomeType), inp <-chan *SomeType) {
	defer close(done)
	PointerS := []*SomeType{}
	for i := range inp {
		PointerS = append(PointerS, i)
	}
	done <- PointerS
}

// DonePointerSlice returns a channel which will receive a slice
// of all the Pointers received on inp channel before close.
// Unlike DonePointer, a full slice is sent once, not just an event.
func DonePointerSlice(inp <-chan *SomeType) (done <-chan ([]*SomeType)) {
	cha := make(chan ([]*SomeType))
	go doitPointerSlice(cha, inp)
	return cha
}

func doitPointerFunc(done chan<- struct{}, inp <-chan *SomeType, act func(a *SomeType)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DonePointerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePointerFunc(inp <-chan *SomeType, act func(a *SomeType)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *SomeType) { return }
	}
	go doitPointerFunc(cha, inp, act)
	return cha
}

func pipePointerBuffer(out chan<- *SomeType, inp <-chan *SomeType) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipePointerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePointerBuffer(inp <-chan *SomeType, cap int) (out <-chan *SomeType) {
	cha := make(chan *SomeType, cap)
	go pipePointerBuffer(cha, inp)
	return cha
}

func pipePointerFunc(out chan<- *SomeType, inp <-chan *SomeType, act func(a *SomeType) *SomeType) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipePointerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePointerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePointerFunc(inp <-chan *SomeType, act func(a *SomeType) *SomeType) (out <-chan *SomeType) {
	cha := make(chan *SomeType)
	if act == nil {
		act = func(a *SomeType) *SomeType { return a }
	}
	go pipePointerFunc(cha, inp, act)
	return cha
}

func pipePointerFork(out1, out2 chan<- *SomeType, inp <-chan *SomeType) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipePointerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePointerFork(inp <-chan *SomeType) (out1, out2 <-chan *SomeType) {
	cha1 := make(chan *SomeType)
	cha2 := make(chan *SomeType)
	go pipePointerFork(cha1, cha2, inp)
	return cha1, cha2
}

// PointerTube is the signature for a pipe function.
type PointerTube func(inp <-chan *SomeType, out <-chan *SomeType)

// PointerDaisy returns a channel to receive all inp after having passed thru tube.
func PointerDaisy(inp <-chan *SomeType, tube PointerTube) (out <-chan *SomeType) {
	cha := make(chan *SomeType)
	go tube(inp, cha)
	return cha
}

// PointerDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func PointerDaisyChain(inp <-chan *SomeType, tubes ...PointerTube) (out <-chan *SomeType) {
	cha := inp
	for i := range tubes {
		cha = PointerDaisy(cha, tubes[i])
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
