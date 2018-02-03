// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeUInt64Chan returns a new open channel
// (simply a 'chan uint64' that is).
//
// Note: No 'UInt64-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myUInt64PipelineStartsHere := MakeUInt64Chan()
//	// ... lot's of code to design and build Your favourite "myUInt64WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myUInt64PipelineStartsHere <- drop
//	}
//	close(myUInt64PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeUInt64Buffer) the channel is unbuffered.
//
func MakeUInt64Chan() (out chan uint64) {
	return make(chan uint64)
}

func sendUInt64(out chan<- uint64, inp ...uint64) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanUInt64 returns a channel to receive all inputs before close.
func ChanUInt64(inp ...uint64) (out <-chan uint64) {
	cha := make(chan uint64)
	go sendUInt64(cha, inp...)
	return cha
}

func sendUInt64Slice(out chan<- uint64, inp ...[]uint64) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanUInt64Slice returns a channel to receive all inputs before close.
func ChanUInt64Slice(inp ...[]uint64) (out <-chan uint64) {
	cha := make(chan uint64)
	go sendUInt64Slice(cha, inp...)
	return cha
}

func chanUInt64FuncNok(out chan<- uint64, act func() (uint64, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanUInt64FuncNok returns a channel to receive all results of act until nok before close.
func ChanUInt64FuncNok(act func() (uint64, bool)) (out <-chan uint64) {
	cha := make(chan uint64)
	go chanUInt64FuncNok(cha, act)
	return cha
}

func chanUInt64FuncErr(out chan<- uint64, act func() (uint64, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanUInt64FuncErr returns a channel to receive all results of act until err != nil before close.
func ChanUInt64FuncErr(act func() (uint64, error)) (out <-chan uint64) {
	cha := make(chan uint64)
	go chanUInt64FuncErr(cha, act)
	return cha
}

func joinUInt64(done chan<- struct{}, out chan<- uint64, inp ...uint64) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinUInt64 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt64(out chan<- uint64, inp ...uint64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinUInt64(cha, out, inp...)
	return cha
}

func joinUInt64Slice(done chan<- struct{}, out chan<- uint64, inp ...[]uint64) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinUInt64Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt64Slice(out chan<- uint64, inp ...[]uint64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinUInt64Slice(cha, out, inp...)
	return cha
}

func joinUInt64Chan(done chan<- struct{}, out chan<- uint64, inp <-chan uint64) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinUInt64Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinUInt64Chan(out chan<- uint64, inp <-chan uint64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinUInt64Chan(cha, out, inp)
	return cha
}

func doitUInt64(done chan<- struct{}, inp <-chan uint64) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneUInt64 returns a channel to receive one signal before close after inp has been drained.
func DoneUInt64(inp <-chan uint64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitUInt64(cha, inp)
	return cha
}

func doitUInt64Slice(done chan<- ([]uint64), inp <-chan uint64) {
	defer close(done)
	UInt64S := []uint64{}
	for i := range inp {
		UInt64S = append(UInt64S, i)
	}
	done <- UInt64S
}

// DoneUInt64Slice returns a channel which will receive a slice
// of all the UInt64s received on inp channel before close.
// Unlike DoneUInt64, a full slice is sent once, not just an event.
func DoneUInt64Slice(inp <-chan uint64) (done <-chan ([]uint64)) {
	cha := make(chan ([]uint64))
	go doitUInt64Slice(cha, inp)
	return cha
}

func doitUInt64Func(done chan<- struct{}, inp <-chan uint64, act func(a uint64)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneUInt64Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneUInt64Func(inp <-chan uint64, act func(a uint64)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a uint64) { return }
	}
	go doitUInt64Func(cha, inp, act)
	return cha
}

func pipeUInt64Buffer(out chan<- uint64, inp <-chan uint64) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeUInt64Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeUInt64Buffer(inp <-chan uint64, cap int) (out <-chan uint64) {
	cha := make(chan uint64, cap)
	go pipeUInt64Buffer(cha, inp)
	return cha
}

func pipeUInt64Func(out chan<- uint64, inp <-chan uint64, act func(a uint64) uint64) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeUInt64Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeUInt64Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeUInt64Func(inp <-chan uint64, act func(a uint64) uint64) (out <-chan uint64) {
	cha := make(chan uint64)
	if act == nil {
		act = func(a uint64) uint64 { return a }
	}
	go pipeUInt64Func(cha, inp, act)
	return cha
}

func pipeUInt64Fork(out1, out2 chan<- uint64, inp <-chan uint64) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeUInt64Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeUInt64Fork(inp <-chan uint64) (out1, out2 <-chan uint64) {
	cha1 := make(chan uint64)
	cha2 := make(chan uint64)
	go pipeUInt64Fork(cha1, cha2, inp)
	return cha1, cha2
}

// UInt64Tube is the signature for a pipe function.
type UInt64Tube func(inp <-chan uint64, out <-chan uint64)

// UInt64Daisy returns a channel to receive all inp after having passed thru tube.
func UInt64Daisy(inp <-chan uint64, tube UInt64Tube) (out <-chan uint64) {
	cha := make(chan uint64)
	go tube(inp, cha)
	return cha
}

// UInt64DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func UInt64DaisyChain(inp <-chan uint64, tubes ...UInt64Tube) (out <-chan uint64) {
	cha := inp
	for i := range tubes {
		cha = UInt64Daisy(cha, tubes[i])
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

// MergeUInt64 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed UInt64channel is returned.
func MergeUInt64(inps ...<-chan uint64) (out <-chan uint64) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan uint64)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeUInt642(inps[0], MergeUInt64(inps[1:]...))
	}
}

// mergeUInt642 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeUInt642(i1, i2 <-chan uint64) (out <-chan uint64) {
	cha := make(chan uint64)
	go func(out chan<- uint64, i1, i2 <-chan uint64) {
		defer close(out)
		var (
			clos1, clos2 bool   // we found the chan closed
			buff1, buff2 bool   // we've read 'from', but not sent (yet)
			ok           bool   // did we read sucessfully?
			from1, from2 uint64 // what we've read
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
