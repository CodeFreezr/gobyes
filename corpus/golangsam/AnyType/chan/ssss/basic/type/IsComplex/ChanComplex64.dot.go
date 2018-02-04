// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsComplex

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeComplex64Chan returns a new open channel
// (simply a 'chan complex64' that is).
//
// Note: No 'Complex64-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myComplex64PipelineStartsHere := MakeComplex64Chan()
//	// ... lot's of code to design and build Your favourite "myComplex64WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myComplex64PipelineStartsHere <- drop
//	}
//	close(myComplex64PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeComplex64Buffer) the channel is unbuffered.
//
func MakeComplex64Chan() chan complex64 {
	return make(chan complex64)
}

// ChanComplex64 returns a channel to receive all inputs before close.
func ChanComplex64(inp ...complex64) chan complex64 {
	out := make(chan complex64)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanComplex64Slice returns a channel to receive all inputs before close.
func ChanComplex64Slice(inp ...[]complex64) chan complex64 {
	out := make(chan complex64)
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

// ChanComplex64FuncNok returns a channel to receive all results of act until nok before close.
func ChanComplex64FuncNok(act func() (complex64, bool)) <-chan complex64 {
	out := make(chan complex64)
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

// ChanComplex64FuncErr returns a channel to receive all results of act until err != nil before close.
func ChanComplex64FuncErr(act func() (complex64, error)) <-chan complex64 {
	out := make(chan complex64)
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

// JoinComplex64 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinComplex64(out chan<- complex64, inp ...complex64) chan struct{} {
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

// JoinComplex64Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinComplex64Slice(out chan<- complex64, inp ...[]complex64) chan struct{} {
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

// JoinComplex64Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinComplex64Chan(out chan<- complex64, inp <-chan complex64) chan struct{} {
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

// DoneComplex64 returns a channel to receive one signal before close after inp has been drained.
func DoneComplex64(inp <-chan complex64) chan struct{} {
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

// DoneComplex64Slice returns a channel which will receive a slice
// of all the Complex64s received on inp channel before close.
// Unlike DoneComplex64, a full slice is sent once, not just an event.
func DoneComplex64Slice(inp <-chan complex64) chan []complex64 {
	done := make(chan []complex64)
	go func() {
		defer close(done)
		Complex64S := []complex64{}
		for i := range inp {
			Complex64S = append(Complex64S, i)
		}
		done <- Complex64S
	}()
	return done
}

// DoneComplex64Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneComplex64Func(inp <-chan complex64, act func(a complex64)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a complex64) { return }
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

// PipeComplex64Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeComplex64Buffer(inp <-chan complex64, cap int) chan complex64 {
	out := make(chan complex64, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeComplex64Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeComplex64Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeComplex64Func(inp <-chan complex64, act func(a complex64) complex64) chan complex64 {
	out := make(chan complex64)
	if act == nil {
		act = func(a complex64) complex64 { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeComplex64Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeComplex64Fork(inp <-chan complex64) (chan complex64, chan complex64) {
	out1 := make(chan complex64)
	out2 := make(chan complex64)
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

// Complex64Tube is the signature for a pipe function.
type Complex64Tube func(inp <-chan complex64, out <-chan complex64)

// Complex64Daisy returns a channel to receive all inp after having passed thru tube.
func Complex64Daisy(inp <-chan complex64, tube Complex64Tube) (out <-chan complex64) {
	cha := make(chan complex64)
	go tube(inp, cha)
	return cha
}

// Complex64DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func Complex64DaisyChain(inp <-chan complex64, tubes ...Complex64Tube) (out <-chan complex64) {
	cha := inp
	for i := range tubes {
		cha = Complex64Daisy(cha, tubes[i])
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
