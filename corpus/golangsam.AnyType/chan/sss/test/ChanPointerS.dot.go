// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakePointerSChan returns a new open channel
// (simply a 'chan []*SomeType' that is).
//
// Note: No 'PointerS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPointerSPipelineStartsHere := MakePointerSChan()
//	// ... lot's of code to design and build Your favourite "myPointerSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPointerSPipelineStartsHere <- drop
//	}
//	close(myPointerSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePointerSBuffer) the channel is unbuffered.
//
func MakePointerSChan() (out chan []*SomeType) {
	return make(chan []*SomeType)
}

// ChanPointerS returns a channel to receive all inputs before close.
func ChanPointerS(inp ...[]*SomeType) (out <-chan []*SomeType) {
	cha := make(chan []*SomeType)
	go func(out chan<- []*SomeType, inp ...[]*SomeType) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanPointerSSlice returns a channel to receive all inputs before close.
func ChanPointerSSlice(inp ...[][]*SomeType) (out <-chan []*SomeType) {
	cha := make(chan []*SomeType)
	go func(out chan<- []*SomeType, inp ...[][]*SomeType) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanPointerSFuncNok returns a channel to receive all results of act until nok before close.
func ChanPointerSFuncNok(act func() ([]*SomeType, bool)) (out <-chan []*SomeType) {
	cha := make(chan []*SomeType)
	go func(out chan<- []*SomeType, act func() ([]*SomeType, bool)) {
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

// ChanPointerSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanPointerSFuncErr(act func() ([]*SomeType, error)) (out <-chan []*SomeType) {
	cha := make(chan []*SomeType)
	go func(out chan<- []*SomeType, act func() ([]*SomeType, error)) {
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

// JoinPointerS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPointerS(out chan<- []*SomeType, inp ...[]*SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []*SomeType, inp ...[]*SomeType) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinPointerSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPointerSSlice(out chan<- []*SomeType, inp ...[][]*SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []*SomeType, inp ...[][]*SomeType) {
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

// JoinPointerSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPointerSChan(out chan<- []*SomeType, inp <-chan []*SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []*SomeType, inp <-chan []*SomeType) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DonePointerS returns a channel to receive one signal before close after inp has been drained.
func DonePointerS(inp <-chan []*SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan []*SomeType) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DonePointerSSlice returns a channel which will receive a slice
// of all the PointerSs received on inp channel before close.
// Unlike DonePointerS, a full slice is sent once, not just an event.
func DonePointerSSlice(inp <-chan []*SomeType) (done <-chan [][]*SomeType) {
	cha := make(chan [][]*SomeType)
	go func(inp <-chan []*SomeType, done chan<- [][]*SomeType) {
		defer close(done)
		PointerSS := [][]*SomeType{}
		for i := range inp {
			PointerSS = append(PointerSS, i)
		}
		done <- PointerSS
	}(inp, cha)
	return cha
}

// DonePointerSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePointerSFunc(inp <-chan []*SomeType, act func(a []*SomeType)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []*SomeType) { return }
	}
	go func(done chan<- struct{}, inp <-chan []*SomeType, act func(a []*SomeType)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipePointerSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePointerSBuffer(inp <-chan []*SomeType, cap int) (out <-chan []*SomeType) {
	cha := make(chan []*SomeType, cap)
	go func(out chan<- []*SomeType, inp <-chan []*SomeType) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipePointerSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePointerSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePointerSFunc(inp <-chan []*SomeType, act func(a []*SomeType) []*SomeType) (out <-chan []*SomeType) {
	cha := make(chan []*SomeType)
	if act == nil {
		act = func(a []*SomeType) []*SomeType { return a }
	}
	go func(out chan<- []*SomeType, inp <-chan []*SomeType, act func(a []*SomeType) []*SomeType) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipePointerSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePointerSFork(inp <-chan []*SomeType) (out1, out2 <-chan []*SomeType) {
	cha1 := make(chan []*SomeType)
	cha2 := make(chan []*SomeType)
	go func(out1, out2 chan<- []*SomeType, inp <-chan []*SomeType) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// PointerSTube is the signature for a pipe function.
type PointerSTube func(inp <-chan []*SomeType, out <-chan []*SomeType)

// PointerSDaisy returns a channel to receive all inp after having passed thru tube.
func PointerSDaisy(inp <-chan []*SomeType, tube PointerSTube) (out <-chan []*SomeType) {
	cha := make(chan []*SomeType)
	go tube(inp, cha)
	return cha
}

// PointerSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func PointerSDaisyChain(inp <-chan []*SomeType, tubes ...PointerSTube) (out <-chan []*SomeType) {
	cha := inp
	for i := range tubes {
		cha = PointerSDaisy(cha, tubes[i])
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
