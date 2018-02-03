// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeInt32Chan returns a new open channel
// (simply a 'chan int32' that is).
//
// Note: No 'Int32-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myInt32PipelineStartsHere := MakeInt32Chan()
//	// ... lot's of code to design and build Your favourite "myInt32WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myInt32PipelineStartsHere <- drop
//	}
//	close(myInt32PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeInt32Buffer) the channel is unbuffered.
//
func MakeInt32Chan() chan int32 {
	return make(chan int32)
}

// ChanInt32 returns a channel to receive all inputs before close.
func ChanInt32(inp ...int32) chan int32 {
	out := make(chan int32)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanInt32Slice returns a channel to receive all inputs before close.
func ChanInt32Slice(inp ...[]int32) chan int32 {
	out := make(chan int32)
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

// ChanInt32FuncNok returns a channel to receive all results of act until nok before close.
func ChanInt32FuncNok(act func() (int32, bool)) <-chan int32 {
	out := make(chan int32)
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

// ChanInt32FuncErr returns a channel to receive all results of act until err != nil before close.
func ChanInt32FuncErr(act func() (int32, error)) <-chan int32 {
	out := make(chan int32)
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

// JoinInt32 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt32(out chan<- int32, inp ...int32) chan struct{} {
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

// JoinInt32Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt32Slice(out chan<- int32, inp ...[]int32) chan struct{} {
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

// JoinInt32Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt32Chan(out chan<- int32, inp <-chan int32) chan struct{} {
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

// DoneInt32 returns a channel to receive one signal before close after inp has been drained.
func DoneInt32(inp <-chan int32) chan struct{} {
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

// DoneInt32Slice returns a channel which will receive a slice
// of all the Int32s received on inp channel before close.
// Unlike DoneInt32, a full slice is sent once, not just an event.
func DoneInt32Slice(inp <-chan int32) chan []int32 {
	done := make(chan []int32)
	go func() {
		defer close(done)
		Int32S := []int32{}
		for i := range inp {
			Int32S = append(Int32S, i)
		}
		done <- Int32S
	}()
	return done
}

// DoneInt32Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneInt32Func(inp <-chan int32, act func(a int32)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a int32) { return }
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

// PipeInt32Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeInt32Buffer(inp <-chan int32, cap int) chan int32 {
	out := make(chan int32, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeInt32Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeInt32Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeInt32Func(inp <-chan int32, act func(a int32) int32) chan int32 {
	out := make(chan int32)
	if act == nil {
		act = func(a int32) int32 { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeInt32Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeInt32Fork(inp <-chan int32) (chan int32, chan int32) {
	out1 := make(chan int32)
	out2 := make(chan int32)
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

// Int32Tube is the signature for a pipe function.
type Int32Tube func(inp <-chan int32, out <-chan int32)

// Int32Daisy returns a channel to receive all inp after having passed thru tube.
func Int32Daisy(inp <-chan int32, tube Int32Tube) (out <-chan int32) {
	cha := make(chan int32)
	go tube(inp, cha)
	return cha
}

// Int32DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func Int32DaisyChain(inp <-chan int32, tubes ...Int32Tube) (out <-chan int32) {
	cha := inp
	for i := range tubes {
		cha = Int32Daisy(cha, tubes[i])
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

// MergeInt32 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Int32channel is returned.
func MergeInt32(inps ...<-chan int32) (out <-chan int32) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan int32)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeInt322(inps[0], MergeInt32(inps[1:]...))
	}
}

// mergeInt322 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeInt322(i1, i2 <-chan int32) (out <-chan int32) {
	cha := make(chan int32)
	go func(out chan<- int32, i1, i2 <-chan int32) {
		defer close(out)
		var (
			clos1, clos2 bool  // we found the chan closed
			buff1, buff2 bool  // we've read 'from', but not sent (yet)
			ok           bool  // did we read sucessfully?
			from1, from2 int32 // what we've read
		)

		for !clos1 || !clos2 {

			if !clos1 && !buff1 {
				if from1, ok = <-i1; ok {
					buff1 = true
				} else {
					clos1 = true
				}
			}

			if !clos2 && !buff2 {
				if from2, ok = <-i2; ok {
					buff2 = true
				} else {
					clos2 = true
				}
			}

			if clos1 && !buff1 {
				from1 = from2
			}
			if clos2 && !buff2 {
				from2 = from1
			}

			if from1 < from2 {
				out <- from1
				buff1 = false
			} else if from2 < from1 {
				out <- from2
				buff2 = false
			} else {
				out <- from1 // == from2
				buff1 = false
				buff2 = false
			}
		}
	}(cha, i1, i2)
	return cha
}

// Note: merge2 is not my own. Just: I forgot where found it - please accept my apologies.
// I'd love to learn about it's origin/author, so I can give credit. Any hint is highly appreciated!
