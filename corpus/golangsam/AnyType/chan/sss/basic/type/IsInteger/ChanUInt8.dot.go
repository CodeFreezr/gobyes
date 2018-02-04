// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsInteger

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeUInt8Chan returns a new open channel
// (simply a 'chan uint8' that is).
//
// Note: No 'UInt8-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myUInt8PipelineStartsHere := MakeUInt8Chan()
//	// ... lot's of code to design and build Your favourite "myUInt8WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myUInt8PipelineStartsHere <- drop
//	}
//	close(myUInt8PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeUInt8Buffer) the channel is unbuffered.
//
func MakeUInt8Chan() (out chan uint8) {
	return make(chan uint8)
}

// ChanUInt8 returns a channel to receive all inputs before close.
func ChanUInt8(inp ...uint8) (out <-chan uint8) {
	cha := make(chan uint8)
	go func(out chan<- uint8, inp ...uint8) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanUInt8Slice returns a channel to receive all inputs before close.
func ChanUInt8Slice(inp ...[]uint8) (out <-chan uint8) {
	cha := make(chan uint8)
	go func(out chan<- uint8, inp ...[]uint8) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanUInt8FuncNok returns a channel to receive all results of act until nok before close.
func ChanUInt8FuncNok(act func() (uint8, bool)) (out <-chan uint8) {
	cha := make(chan uint8)
	go func(out chan<- uint8, act func() (uint8, bool)) {
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

// ChanUInt8FuncErr returns a channel to receive all results of act until err != nil before close.
func ChanUInt8FuncErr(act func() (uint8, error)) (out <-chan uint8) {
	cha := make(chan uint8)
	go func(out chan<- uint8, act func() (uint8, error)) {
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

// JoinUInt8 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt8(out chan<- uint8, inp ...uint8) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- uint8, inp ...uint8) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinUInt8Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt8Slice(out chan<- uint8, inp ...[]uint8) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- uint8, inp ...[]uint8) {
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

// JoinUInt8Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt8Chan(out chan<- uint8, inp <-chan uint8) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- uint8, inp <-chan uint8) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneUInt8 returns a channel to receive one signal before close after inp has been drained.
func DoneUInt8(inp <-chan uint8) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan uint8) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneUInt8Slice returns a channel which will receive a slice
// of all the UInt8s received on inp channel before close.
// Unlike DoneUInt8, a full slice is sent once, not just an event.
func DoneUInt8Slice(inp <-chan uint8) (done <-chan []uint8) {
	cha := make(chan []uint8)
	go func(inp <-chan uint8, done chan<- []uint8) {
		defer close(done)
		UInt8S := []uint8{}
		for i := range inp {
			UInt8S = append(UInt8S, i)
		}
		done <- UInt8S
	}(inp, cha)
	return cha
}

// DoneUInt8Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneUInt8Func(inp <-chan uint8, act func(a uint8)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a uint8) { return }
	}
	go func(done chan<- struct{}, inp <-chan uint8, act func(a uint8)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeUInt8Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeUInt8Buffer(inp <-chan uint8, cap int) (out <-chan uint8) {
	cha := make(chan uint8, cap)
	go func(out chan<- uint8, inp <-chan uint8) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeUInt8Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeUInt8Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeUInt8Func(inp <-chan uint8, act func(a uint8) uint8) (out <-chan uint8) {
	cha := make(chan uint8)
	if act == nil {
		act = func(a uint8) uint8 { return a }
	}
	go func(out chan<- uint8, inp <-chan uint8, act func(a uint8) uint8) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeUInt8Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeUInt8Fork(inp <-chan uint8) (out1, out2 <-chan uint8) {
	cha1 := make(chan uint8)
	cha2 := make(chan uint8)
	go func(out1, out2 chan<- uint8, inp <-chan uint8) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// UInt8Tube is the signature for a pipe function.
type UInt8Tube func(inp <-chan uint8, out <-chan uint8)

// UInt8Daisy returns a channel to receive all inp after having passed thru tube.
func UInt8Daisy(inp <-chan uint8, tube UInt8Tube) (out <-chan uint8) {
	cha := make(chan uint8)
	go tube(inp, cha)
	return cha
}

// UInt8DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func UInt8DaisyChain(inp <-chan uint8, tubes ...UInt8Tube) (out <-chan uint8) {
	cha := inp
	for i := range tubes {
		cha = UInt8Daisy(cha, tubes[i])
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

// MergeUInt8 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed UInt8channel is returned.
func MergeUInt8(inps ...<-chan uint8) (out <-chan uint8) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan uint8)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeUInt82(inps[0], MergeUInt8(inps[1:]...))
	}
}

// mergeUInt82 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeUInt82(i1, i2 <-chan uint8) (out <-chan uint8) {
	cha := make(chan uint8)
	go func(out chan<- uint8, i1, i2 <-chan uint8) {
		defer close(out)
		var (
			clos1, clos2 bool  // we found the chan closed
			buff1, buff2 bool  // we've read 'from', but not sent (yet)
			ok           bool  // did we read sucessfully?
			from1, from2 uint8 // what we've read
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
