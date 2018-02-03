// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeUInt16Chan returns a new open channel
// (simply a 'chan uint16' that is).
//
// Note: No 'UInt16-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myUInt16PipelineStartsHere := MakeUInt16Chan()
//	// ... lot's of code to design and build Your favourite "myUInt16WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myUInt16PipelineStartsHere <- drop
//	}
//	close(myUInt16PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeUInt16Buffer) the channel is unbuffered.
//
func MakeUInt16Chan() chan uint16 {
	return make(chan uint16)
}

// ChanUInt16 returns a channel to receive all inputs before close.
func ChanUInt16(inp ...uint16) chan uint16 {
	out := make(chan uint16)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanUInt16Slice returns a channel to receive all inputs before close.
func ChanUInt16Slice(inp ...[]uint16) chan uint16 {
	out := make(chan uint16)
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

// ChanUInt16FuncNok returns a channel to receive all results of act until nok before close.
func ChanUInt16FuncNok(act func() (uint16, bool)) <-chan uint16 {
	out := make(chan uint16)
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

// ChanUInt16FuncErr returns a channel to receive all results of act until err != nil before close.
func ChanUInt16FuncErr(act func() (uint16, error)) <-chan uint16 {
	out := make(chan uint16)
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

// JoinUInt16 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt16(out chan<- uint16, inp ...uint16) chan struct{} {
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

// JoinUInt16Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt16Slice(out chan<- uint16, inp ...[]uint16) chan struct{} {
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

// JoinUInt16Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt16Chan(out chan<- uint16, inp <-chan uint16) chan struct{} {
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

// DoneUInt16 returns a channel to receive one signal before close after inp has been drained.
func DoneUInt16(inp <-chan uint16) chan struct{} {
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

// DoneUInt16Slice returns a channel which will receive a slice
// of all the UInt16s received on inp channel before close.
// Unlike DoneUInt16, a full slice is sent once, not just an event.
func DoneUInt16Slice(inp <-chan uint16) chan []uint16 {
	done := make(chan []uint16)
	go func() {
		defer close(done)
		UInt16S := []uint16{}
		for i := range inp {
			UInt16S = append(UInt16S, i)
		}
		done <- UInt16S
	}()
	return done
}

// DoneUInt16Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneUInt16Func(inp <-chan uint16, act func(a uint16)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a uint16) { return }
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

// PipeUInt16Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeUInt16Buffer(inp <-chan uint16, cap int) chan uint16 {
	out := make(chan uint16, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeUInt16Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeUInt16Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeUInt16Func(inp <-chan uint16, act func(a uint16) uint16) chan uint16 {
	out := make(chan uint16)
	if act == nil {
		act = func(a uint16) uint16 { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeUInt16Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeUInt16Fork(inp <-chan uint16) (chan uint16, chan uint16) {
	out1 := make(chan uint16)
	out2 := make(chan uint16)
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

// UInt16Tube is the signature for a pipe function.
type UInt16Tube func(inp <-chan uint16, out <-chan uint16)

// UInt16Daisy returns a channel to receive all inp after having passed thru tube.
func UInt16Daisy(inp <-chan uint16, tube UInt16Tube) (out <-chan uint16) {
	cha := make(chan uint16)
	go tube(inp, cha)
	return cha
}

// UInt16DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func UInt16DaisyChain(inp <-chan uint16, tubes ...UInt16Tube) (out <-chan uint16) {
	cha := inp
	for i := range tubes {
		cha = UInt16Daisy(cha, tubes[i])
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
