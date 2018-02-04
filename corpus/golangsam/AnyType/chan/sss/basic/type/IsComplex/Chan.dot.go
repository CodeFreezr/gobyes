// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsComplex

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeChan returns a new open channel
// (simply a 'chan complex128' that is).
//
// Note: No '-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPipelineStartsHere := MakeChan()
//	// ... lot's of code to design and build Your favourite "myWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPipelineStartsHere <- drop
//	}
//	close(myPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeBuffer) the channel is unbuffered.
//
func MakeChan() (out chan complex128) {
	return make(chan complex128)
}

// Chan returns a channel to receive all inputs before close.
func Chan(inp ...complex128) (out <-chan complex128) {
	cha := make(chan complex128)
	go func(out chan<- complex128, inp ...complex128) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanSlice returns a channel to receive all inputs before close.
func ChanSlice(inp ...[]complex128) (out <-chan complex128) {
	cha := make(chan complex128)
	go func(out chan<- complex128, inp ...[]complex128) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanFuncNok returns a channel to receive all results of act until nok before close.
func ChanFuncNok(act func() (complex128, bool)) (out <-chan complex128) {
	cha := make(chan complex128)
	go func(out chan<- complex128, act func() (complex128, bool)) {
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

// ChanFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFuncErr(act func() (complex128, error)) (out <-chan complex128) {
	cha := make(chan complex128)
	go func(out chan<- complex128, act func() (complex128, error)) {
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

// Join sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func Join(out chan<- complex128, inp ...complex128) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- complex128, inp ...complex128) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSlice(out chan<- complex128, inp ...[]complex128) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- complex128, inp ...[]complex128) {
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

// JoinChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinChan(out chan<- complex128, inp <-chan complex128) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- complex128, inp <-chan complex128) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// Done returns a channel to receive one signal before close after inp has been drained.
func Done(inp <-chan complex128) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan complex128) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneSlice returns a channel which will receive a slice
// of all the s received on inp channel before close.
// Unlike Done, a full slice is sent once, not just an event.
func DoneSlice(inp <-chan complex128) (done <-chan []complex128) {
	cha := make(chan []complex128)
	go func(inp <-chan complex128, done chan<- []complex128) {
		defer close(done)
		S := []complex128{}
		for i := range inp {
			S = append(S, i)
		}
		done <- S
	}(inp, cha)
	return cha
}

// DoneFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFunc(inp <-chan complex128, act func(a complex128)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a complex128) { return }
	}
	go func(done chan<- struct{}, inp <-chan complex128, act func(a complex128)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeBuffer(inp <-chan complex128, cap int) (out <-chan complex128) {
	cha := make(chan complex128, cap)
	go func(out chan<- complex128, inp <-chan complex128) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFunc(inp <-chan complex128, act func(a complex128) complex128) (out <-chan complex128) {
	cha := make(chan complex128)
	if act == nil {
		act = func(a complex128) complex128 { return a }
	}
	go func(out chan<- complex128, inp <-chan complex128, act func(a complex128) complex128) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFork(inp <-chan complex128) (out1, out2 <-chan complex128) {
	cha1 := make(chan complex128)
	cha2 := make(chan complex128)
	go func(out1, out2 chan<- complex128, inp <-chan complex128) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// Tube is the signature for a pipe function.
type Tube func(inp <-chan complex128, out <-chan complex128)

// Daisy returns a channel to receive all inp after having passed thru tube.
func Daisy(inp <-chan complex128, tube Tube) (out <-chan complex128) {
	cha := make(chan complex128)
	go tube(inp, cha)
	return cha
}

// DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func DaisyChain(inp <-chan complex128, tubes ...Tube) (out <-chan complex128) {
	cha := inp
	for i := range tubes {
		cha = Daisy(cha, tubes[i])
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
