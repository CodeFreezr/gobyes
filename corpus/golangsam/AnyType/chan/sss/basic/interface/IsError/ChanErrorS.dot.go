// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsError

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeErrorSChan returns a new open channel
// (simply a 'chan []error' that is).
//
// Note: No 'ErrorS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myErrorSPipelineStartsHere := MakeErrorSChan()
//	// ... lot's of code to design and build Your favourite "myErrorSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myErrorSPipelineStartsHere <- drop
//	}
//	close(myErrorSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeErrorSBuffer) the channel is unbuffered.
//
func MakeErrorSChan() (out chan []error) {
	return make(chan []error)
}

// ChanErrorS returns a channel to receive all inputs before close.
func ChanErrorS(inp ...[]error) (out <-chan []error) {
	cha := make(chan []error)
	go func(out chan<- []error, inp ...[]error) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanErrorSSlice returns a channel to receive all inputs before close.
func ChanErrorSSlice(inp ...[][]error) (out <-chan []error) {
	cha := make(chan []error)
	go func(out chan<- []error, inp ...[][]error) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanErrorSFuncNok returns a channel to receive all results of act until nok before close.
func ChanErrorSFuncNok(act func() ([]error, bool)) (out <-chan []error) {
	cha := make(chan []error)
	go func(out chan<- []error, act func() ([]error, bool)) {
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

// ChanErrorSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanErrorSFuncErr(act func() ([]error, error)) (out <-chan []error) {
	cha := make(chan []error)
	go func(out chan<- []error, act func() ([]error, error)) {
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

// JoinErrorS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinErrorS(out chan<- []error, inp ...[]error) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []error, inp ...[]error) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinErrorSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinErrorSSlice(out chan<- []error, inp ...[][]error) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []error, inp ...[][]error) {
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

// JoinErrorSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinErrorSChan(out chan<- []error, inp <-chan []error) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []error, inp <-chan []error) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneErrorS returns a channel to receive one signal before close after inp has been drained.
func DoneErrorS(inp <-chan []error) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan []error) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneErrorSSlice returns a channel which will receive a slice
// of all the ErrorSs received on inp channel before close.
// Unlike DoneErrorS, a full slice is sent once, not just an event.
func DoneErrorSSlice(inp <-chan []error) (done <-chan [][]error) {
	cha := make(chan [][]error)
	go func(inp <-chan []error, done chan<- [][]error) {
		defer close(done)
		ErrorSS := [][]error{}
		for i := range inp {
			ErrorSS = append(ErrorSS, i)
		}
		done <- ErrorSS
	}(inp, cha)
	return cha
}

// DoneErrorSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneErrorSFunc(inp <-chan []error, act func(a []error)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []error) { return }
	}
	go func(done chan<- struct{}, inp <-chan []error, act func(a []error)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeErrorSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeErrorSBuffer(inp <-chan []error, cap int) (out <-chan []error) {
	cha := make(chan []error, cap)
	go func(out chan<- []error, inp <-chan []error) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeErrorSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeErrorSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeErrorSFunc(inp <-chan []error, act func(a []error) []error) (out <-chan []error) {
	cha := make(chan []error)
	if act == nil {
		act = func(a []error) []error { return a }
	}
	go func(out chan<- []error, inp <-chan []error, act func(a []error) []error) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeErrorSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeErrorSFork(inp <-chan []error) (out1, out2 <-chan []error) {
	cha1 := make(chan []error)
	cha2 := make(chan []error)
	go func(out1, out2 chan<- []error, inp <-chan []error) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// ErrorSTube is the signature for a pipe function.
type ErrorSTube func(inp <-chan []error, out <-chan []error)

// ErrorSDaisy returns a channel to receive all inp after having passed thru tube.
func ErrorSDaisy(inp <-chan []error, tube ErrorSTube) (out <-chan []error) {
	cha := make(chan []error)
	go tube(inp, cha)
	return cha
}

// ErrorSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ErrorSDaisyChain(inp <-chan []error, tubes ...ErrorSTube) (out <-chan []error) {
	cha := inp
	for i := range tubes {
		cha = ErrorSDaisy(cha, tubes[i])
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
