// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeInt8Chan returns a new open channel
// (simply a 'chan int8' that is).
//
// Note: No 'Int8-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myInt8PipelineStartsHere := MakeInt8Chan()
//	// ... lot's of code to design and build Your favourite "myInt8WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myInt8PipelineStartsHere <- drop
//	}
//	close(myInt8PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeInt8Buffer) the channel is unbuffered.
//
func MakeInt8Chan() chan int8 {
	return make(chan int8)
}

// ChanInt8 returns a channel to receive all inputs before close.
func ChanInt8(inp ...int8) chan int8 {
	out := make(chan int8)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanInt8Slice returns a channel to receive all inputs before close.
func ChanInt8Slice(inp ...[]int8) chan int8 {
	out := make(chan int8)
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

// ChanInt8FuncNok returns a channel to receive all results of act until nok before close.
func ChanInt8FuncNok(act func() (int8, bool)) <-chan int8 {
	out := make(chan int8)
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

// ChanInt8FuncErr returns a channel to receive all results of act until err != nil before close.
func ChanInt8FuncErr(act func() (int8, error)) <-chan int8 {
	out := make(chan int8)
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

// JoinInt8 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt8(out chan<- int8, inp ...int8) chan struct{} {
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

// JoinInt8Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt8Slice(out chan<- int8, inp ...[]int8) chan struct{} {
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

// JoinInt8Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt8Chan(out chan<- int8, inp <-chan int8) chan struct{} {
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

// DoneInt8 returns a channel to receive one signal before close after inp has been drained.
func DoneInt8(inp <-chan int8) chan struct{} {
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

// DoneInt8Slice returns a channel which will receive a slice
// of all the Int8s received on inp channel before close.
// Unlike DoneInt8, a full slice is sent once, not just an event.
func DoneInt8Slice(inp <-chan int8) chan []int8 {
	done := make(chan []int8)
	go func() {
		defer close(done)
		Int8S := []int8{}
		for i := range inp {
			Int8S = append(Int8S, i)
		}
		done <- Int8S
	}()
	return done
}

// DoneInt8Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneInt8Func(inp <-chan int8, act func(a int8)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a int8) { return }
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

// PipeInt8Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeInt8Buffer(inp <-chan int8, cap int) chan int8 {
	out := make(chan int8, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeInt8Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeInt8Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeInt8Func(inp <-chan int8, act func(a int8) int8) chan int8 {
	out := make(chan int8)
	if act == nil {
		act = func(a int8) int8 { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeInt8Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeInt8Fork(inp <-chan int8) (chan int8, chan int8) {
	out1 := make(chan int8)
	out2 := make(chan int8)
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

// Int8Tube is the signature for a pipe function.
type Int8Tube func(inp <-chan int8, out <-chan int8)

// Int8Daisy returns a channel to receive all inp after having passed thru tube.
func Int8Daisy(inp <-chan int8, tube Int8Tube) (out <-chan int8) {
	cha := make(chan int8)
	go tube(inp, cha)
	return cha
}

// Int8DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func Int8DaisyChain(inp <-chan int8, tubes ...Int8Tube) (out <-chan int8) {
	cha := inp
	for i := range tubes {
		cha = Int8Daisy(cha, tubes[i])
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
