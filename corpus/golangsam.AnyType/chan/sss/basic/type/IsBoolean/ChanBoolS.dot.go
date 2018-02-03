// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsBoolean

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeBoolSChan returns a new open channel
// (simply a 'chan []bool' that is).
//
// Note: No 'BoolS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myBoolSPipelineStartsHere := MakeBoolSChan()
//	// ... lot's of code to design and build Your favourite "myBoolSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myBoolSPipelineStartsHere <- drop
//	}
//	close(myBoolSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeBoolSBuffer) the channel is unbuffered.
//
func MakeBoolSChan() (out chan []bool) {
	return make(chan []bool)
}

// ChanBoolS returns a channel to receive all inputs before close.
func ChanBoolS(inp ...[]bool) (out <-chan []bool) {
	cha := make(chan []bool)
	go func(out chan<- []bool, inp ...[]bool) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanBoolSSlice returns a channel to receive all inputs before close.
func ChanBoolSSlice(inp ...[][]bool) (out <-chan []bool) {
	cha := make(chan []bool)
	go func(out chan<- []bool, inp ...[][]bool) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanBoolSFuncNok returns a channel to receive all results of act until nok before close.
func ChanBoolSFuncNok(act func() ([]bool, bool)) (out <-chan []bool) {
	cha := make(chan []bool)
	go func(out chan<- []bool, act func() ([]bool, bool)) {
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

// ChanBoolSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanBoolSFuncErr(act func() ([]bool, error)) (out <-chan []bool) {
	cha := make(chan []bool)
	go func(out chan<- []bool, act func() ([]bool, error)) {
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

// JoinBoolS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinBoolS(out chan<- []bool, inp ...[]bool) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []bool, inp ...[]bool) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinBoolSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinBoolSSlice(out chan<- []bool, inp ...[][]bool) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []bool, inp ...[][]bool) {
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

// JoinBoolSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinBoolSChan(out chan<- []bool, inp <-chan []bool) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []bool, inp <-chan []bool) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneBoolS returns a channel to receive one signal before close after inp has been drained.
func DoneBoolS(inp <-chan []bool) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan []bool) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneBoolSSlice returns a channel which will receive a slice
// of all the BoolSs received on inp channel before close.
// Unlike DoneBoolS, a full slice is sent once, not just an event.
func DoneBoolSSlice(inp <-chan []bool) (done <-chan [][]bool) {
	cha := make(chan [][]bool)
	go func(inp <-chan []bool, done chan<- [][]bool) {
		defer close(done)
		BoolSS := [][]bool{}
		for i := range inp {
			BoolSS = append(BoolSS, i)
		}
		done <- BoolSS
	}(inp, cha)
	return cha
}

// DoneBoolSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneBoolSFunc(inp <-chan []bool, act func(a []bool)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []bool) { return }
	}
	go func(done chan<- struct{}, inp <-chan []bool, act func(a []bool)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeBoolSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeBoolSBuffer(inp <-chan []bool, cap int) (out <-chan []bool) {
	cha := make(chan []bool, cap)
	go func(out chan<- []bool, inp <-chan []bool) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeBoolSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeBoolSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeBoolSFunc(inp <-chan []bool, act func(a []bool) []bool) (out <-chan []bool) {
	cha := make(chan []bool)
	if act == nil {
		act = func(a []bool) []bool { return a }
	}
	go func(out chan<- []bool, inp <-chan []bool, act func(a []bool) []bool) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeBoolSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeBoolSFork(inp <-chan []bool) (out1, out2 <-chan []bool) {
	cha1 := make(chan []bool)
	cha2 := make(chan []bool)
	go func(out1, out2 chan<- []bool, inp <-chan []bool) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// BoolSTube is the signature for a pipe function.
type BoolSTube func(inp <-chan []bool, out <-chan []bool)

// BoolSDaisy returns a channel to receive all inp after having passed thru tube.
func BoolSDaisy(inp <-chan []bool, tube BoolSTube) (out <-chan []bool) {
	cha := make(chan []bool)
	go tube(inp, cha)
	return cha
}

// BoolSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func BoolSDaisyChain(inp <-chan []bool, tubes ...BoolSTube) (out <-chan []bool) {
	cha := inp
	for i := range tubes {
		cha = BoolSDaisy(cha, tubes[i])
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
