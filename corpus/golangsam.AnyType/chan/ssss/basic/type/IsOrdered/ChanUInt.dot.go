// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeUIntChan returns a new open channel
// (simply a 'chan uint' that is).
//
// Note: No 'UInt-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myUIntPipelineStartsHere := MakeUIntChan()
//	// ... lot's of code to design and build Your favourite "myUIntWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myUIntPipelineStartsHere <- drop
//	}
//	close(myUIntPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeUIntBuffer) the channel is unbuffered.
//
func MakeUIntChan() chan uint {
	return make(chan uint)
}

// ChanUInt returns a channel to receive all inputs before close.
func ChanUInt(inp ...uint) chan uint {
	out := make(chan uint)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanUIntSlice returns a channel to receive all inputs before close.
func ChanUIntSlice(inp ...[]uint) chan uint {
	out := make(chan uint)
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

// ChanUIntFuncNok returns a channel to receive all results of act until nok before close.
func ChanUIntFuncNok(act func() (uint, bool)) <-chan uint {
	out := make(chan uint)
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

// ChanUIntFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanUIntFuncErr(act func() (uint, error)) <-chan uint {
	out := make(chan uint)
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

// JoinUInt sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt(out chan<- uint, inp ...uint) chan struct{} {
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

// JoinUIntSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUIntSlice(out chan<- uint, inp ...[]uint) chan struct{} {
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

// JoinUIntChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUIntChan(out chan<- uint, inp <-chan uint) chan struct{} {
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

// DoneUInt returns a channel to receive one signal before close after inp has been drained.
func DoneUInt(inp <-chan uint) chan struct{} {
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

// DoneUIntSlice returns a channel which will receive a slice
// of all the UInts received on inp channel before close.
// Unlike DoneUInt, a full slice is sent once, not just an event.
func DoneUIntSlice(inp <-chan uint) chan []uint {
	done := make(chan []uint)
	go func() {
		defer close(done)
		UIntS := []uint{}
		for i := range inp {
			UIntS = append(UIntS, i)
		}
		done <- UIntS
	}()
	return done
}

// DoneUIntFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneUIntFunc(inp <-chan uint, act func(a uint)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a uint) { return }
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

// PipeUIntBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeUIntBuffer(inp <-chan uint, cap int) chan uint {
	out := make(chan uint, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeUIntFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeUIntMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeUIntFunc(inp <-chan uint, act func(a uint) uint) chan uint {
	out := make(chan uint)
	if act == nil {
		act = func(a uint) uint { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeUIntFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeUIntFork(inp <-chan uint) (chan uint, chan uint) {
	out1 := make(chan uint)
	out2 := make(chan uint)
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

// UIntTube is the signature for a pipe function.
type UIntTube func(inp <-chan uint, out <-chan uint)

// UIntDaisy returns a channel to receive all inp after having passed thru tube.
func UIntDaisy(inp <-chan uint, tube UIntTube) (out <-chan uint) {
	cha := make(chan uint)
	go tube(inp, cha)
	return cha
}

// UIntDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func UIntDaisyChain(inp <-chan uint, tubes ...UIntTube) (out <-chan uint) {
	cha := inp
	for i := range tubes {
		cha = UIntDaisy(cha, tubes[i])
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

// MergeUInt returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed UIntchannel is returned.
func MergeUInt(inps ...<-chan uint) (out <-chan uint) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan uint)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeUInt2(inps[0], MergeUInt(inps[1:]...))
	}
}

// mergeUInt2 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeUInt2(i1, i2 <-chan uint) (out <-chan uint) {
	cha := make(chan uint)
	go func(out chan<- uint, i1, i2 <-chan uint) {
		defer close(out)
		var (
			clos1, clos2 bool // we found the chan closed
			buff1, buff2 bool // we've read 'from', but not sent (yet)
			ok           bool // did we read sucessfully?
			from1, from2 uint // what we've read
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
