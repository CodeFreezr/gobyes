// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsInteger

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
func MakeInt8Chan() (out chan int8) {
	return make(chan int8)
}

// ChanInt8 returns a channel to receive all inputs before close.
func ChanInt8(inp ...int8) (out <-chan int8) {
	cha := make(chan int8)
	go func(out chan<- int8, inp ...int8) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanInt8Slice returns a channel to receive all inputs before close.
func ChanInt8Slice(inp ...[]int8) (out <-chan int8) {
	cha := make(chan int8)
	go func(out chan<- int8, inp ...[]int8) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanInt8FuncNok returns a channel to receive all results of act until nok before close.
func ChanInt8FuncNok(act func() (int8, bool)) (out <-chan int8) {
	cha := make(chan int8)
	go func(out chan<- int8, act func() (int8, bool)) {
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

// ChanInt8FuncErr returns a channel to receive all results of act until err != nil before close.
func ChanInt8FuncErr(act func() (int8, error)) (out <-chan int8) {
	cha := make(chan int8)
	go func(out chan<- int8, act func() (int8, error)) {
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

// JoinInt8 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt8(out chan<- int8, inp ...int8) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- int8, inp ...int8) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinInt8Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt8Slice(out chan<- int8, inp ...[]int8) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- int8, inp ...[]int8) {
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

// JoinInt8Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt8Chan(out chan<- int8, inp <-chan int8) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- int8, inp <-chan int8) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneInt8 returns a channel to receive one signal before close after inp has been drained.
func DoneInt8(inp <-chan int8) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan int8) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneInt8Slice returns a channel which will receive a slice
// of all the Int8s received on inp channel before close.
// Unlike DoneInt8, a full slice is sent once, not just an event.
func DoneInt8Slice(inp <-chan int8) (done <-chan []int8) {
	cha := make(chan []int8)
	go func(inp <-chan int8, done chan<- []int8) {
		defer close(done)
		Int8S := []int8{}
		for i := range inp {
			Int8S = append(Int8S, i)
		}
		done <- Int8S
	}(inp, cha)
	return cha
}

// DoneInt8Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneInt8Func(inp <-chan int8, act func(a int8)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a int8) { return }
	}
	go func(done chan<- struct{}, inp <-chan int8, act func(a int8)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeInt8Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeInt8Buffer(inp <-chan int8, cap int) (out <-chan int8) {
	cha := make(chan int8, cap)
	go func(out chan<- int8, inp <-chan int8) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeInt8Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeInt8Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeInt8Func(inp <-chan int8, act func(a int8) int8) (out <-chan int8) {
	cha := make(chan int8)
	if act == nil {
		act = func(a int8) int8 { return a }
	}
	go func(out chan<- int8, inp <-chan int8, act func(a int8) int8) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeInt8Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeInt8Fork(inp <-chan int8) (out1, out2 <-chan int8) {
	cha1 := make(chan int8)
	cha2 := make(chan int8)
	go func(out1, out2 chan<- int8, inp <-chan int8) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
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

// MergeInt8 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Int8channel is returned.
func MergeInt8(inps ...<-chan int8) (out <-chan int8) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan int8)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeInt82(inps[0], MergeInt8(inps[1:]...))
	}
}

// mergeInt82 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeInt82(i1, i2 <-chan int8) (out <-chan int8) {
	cha := make(chan int8)
	go func(out chan<- int8, i1, i2 <-chan int8) {
		defer close(out)
		var (
			clos1, clos2 bool // we found the chan closed
			buff1, buff2 bool // we've read 'from', but not sent (yet)
			ok           bool // did we read sucessfully?
			from1, from2 int8 // what we've read
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
