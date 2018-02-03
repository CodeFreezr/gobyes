// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeComplex128Chan returns a new open channel
// (simply a 'chan complex128' that is).
//
// Note: No 'Complex128-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myComplex128PipelineStartsHere := MakeComplex128Chan()
//	// ... lot's of code to design and build Your favourite "myComplex128WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myComplex128PipelineStartsHere <- drop
//	}
//	close(myComplex128PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeComplex128Buffer) the channel is unbuffered.
//
func MakeComplex128Chan() chan complex128 {
	return make(chan complex128)
}

// ChanComplex128 returns a channel to receive all inputs before close.
func ChanComplex128(inp ...complex128) chan complex128 {
	out := make(chan complex128)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanComplex128Slice returns a channel to receive all inputs before close.
func ChanComplex128Slice(inp ...[]complex128) chan complex128 {
	out := make(chan complex128)
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

// ChanComplex128FuncNok returns a channel to receive all results of act until nok before close.
func ChanComplex128FuncNok(act func() (complex128, bool)) <-chan complex128 {
	out := make(chan complex128)
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

// ChanComplex128FuncErr returns a channel to receive all results of act until err != nil before close.
func ChanComplex128FuncErr(act func() (complex128, error)) <-chan complex128 {
	out := make(chan complex128)
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

// JoinComplex128 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinComplex128(out chan<- complex128, inp ...complex128) chan struct{} {
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

// JoinComplex128Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinComplex128Slice(out chan<- complex128, inp ...[]complex128) chan struct{} {
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

// JoinComplex128Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinComplex128Chan(out chan<- complex128, inp <-chan complex128) chan struct{} {
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

// DoneComplex128 returns a channel to receive one signal before close after inp has been drained.
func DoneComplex128(inp <-chan complex128) chan struct{} {
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

// DoneComplex128Slice returns a channel which will receive a slice
// of all the Complex128s received on inp channel before close.
// Unlike DoneComplex128, a full slice is sent once, not just an event.
func DoneComplex128Slice(inp <-chan complex128) chan []complex128 {
	done := make(chan []complex128)
	go func() {
		defer close(done)
		Complex128S := []complex128{}
		for i := range inp {
			Complex128S = append(Complex128S, i)
		}
		done <- Complex128S
	}()
	return done
}

// DoneComplex128Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneComplex128Func(inp <-chan complex128, act func(a complex128)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a complex128) { return }
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

// PipeComplex128Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeComplex128Buffer(inp <-chan complex128, cap int) chan complex128 {
	out := make(chan complex128, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeComplex128Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeComplex128Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeComplex128Func(inp <-chan complex128, act func(a complex128) complex128) chan complex128 {
	out := make(chan complex128)
	if act == nil {
		act = func(a complex128) complex128 { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeComplex128Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeComplex128Fork(inp <-chan complex128) (chan complex128, chan complex128) {
	out1 := make(chan complex128)
	out2 := make(chan complex128)
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

// Complex128Tube is the signature for a pipe function.
type Complex128Tube func(inp <-chan complex128, out <-chan complex128)

// Complex128Daisy returns a channel to receive all inp after having passed thru tube.
func Complex128Daisy(inp <-chan complex128, tube Complex128Tube) (out <-chan complex128) {
	cha := make(chan complex128)
	go tube(inp, cha)
	return cha
}

// Complex128DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func Complex128DaisyChain(inp <-chan complex128, tubes ...Complex128Tube) (out <-chan complex128) {
	cha := inp
	for i := range tubes {
		cha = Complex128Daisy(cha, tubes[i])
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
