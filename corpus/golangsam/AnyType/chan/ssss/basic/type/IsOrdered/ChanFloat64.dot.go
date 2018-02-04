// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeFloat64Chan returns a new open channel
// (simply a 'chan float64' that is).
//
// Note: No 'Float64-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFloat64PipelineStartsHere := MakeFloat64Chan()
//	// ... lot's of code to design and build Your favourite "myFloat64WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFloat64PipelineStartsHere <- drop
//	}
//	close(myFloat64PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFloat64Buffer) the channel is unbuffered.
//
func MakeFloat64Chan() chan float64 {
	return make(chan float64)
}

// ChanFloat64 returns a channel to receive all inputs before close.
func ChanFloat64(inp ...float64) chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanFloat64Slice returns a channel to receive all inputs before close.
func ChanFloat64Slice(inp ...[]float64) chan float64 {
	out := make(chan float64)
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

// ChanFloat64FuncNok returns a channel to receive all results of act until nok before close.
func ChanFloat64FuncNok(act func() (float64, bool)) <-chan float64 {
	out := make(chan float64)
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

// ChanFloat64FuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFloat64FuncErr(act func() (float64, error)) <-chan float64 {
	out := make(chan float64)
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

// JoinFloat64 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFloat64(out chan<- float64, inp ...float64) chan struct{} {
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

// JoinFloat64Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFloat64Slice(out chan<- float64, inp ...[]float64) chan struct{} {
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

// JoinFloat64Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFloat64Chan(out chan<- float64, inp <-chan float64) chan struct{} {
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

// DoneFloat64 returns a channel to receive one signal before close after inp has been drained.
func DoneFloat64(inp <-chan float64) chan struct{} {
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

// DoneFloat64Slice returns a channel which will receive a slice
// of all the Float64s received on inp channel before close.
// Unlike DoneFloat64, a full slice is sent once, not just an event.
func DoneFloat64Slice(inp <-chan float64) chan []float64 {
	done := make(chan []float64)
	go func() {
		defer close(done)
		Float64S := []float64{}
		for i := range inp {
			Float64S = append(Float64S, i)
		}
		done <- Float64S
	}()
	return done
}

// DoneFloat64Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFloat64Func(inp <-chan float64, act func(a float64)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a float64) { return }
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

// PipeFloat64Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFloat64Buffer(inp <-chan float64, cap int) chan float64 {
	out := make(chan float64, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFloat64Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFloat64Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFloat64Func(inp <-chan float64, act func(a float64) float64) chan float64 {
	out := make(chan float64)
	if act == nil {
		act = func(a float64) float64 { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFloat64Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFloat64Fork(inp <-chan float64) (chan float64, chan float64) {
	out1 := make(chan float64)
	out2 := make(chan float64)
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

// Float64Tube is the signature for a pipe function.
type Float64Tube func(inp <-chan float64, out <-chan float64)

// Float64Daisy returns a channel to receive all inp after having passed thru tube.
func Float64Daisy(inp <-chan float64, tube Float64Tube) (out <-chan float64) {
	cha := make(chan float64)
	go tube(inp, cha)
	return cha
}

// Float64DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func Float64DaisyChain(inp <-chan float64, tubes ...Float64Tube) (out <-chan float64) {
	cha := inp
	for i := range tubes {
		cha = Float64Daisy(cha, tubes[i])
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

// MergeFloat64 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Float64channel is returned.
func MergeFloat64(inps ...<-chan float64) (out <-chan float64) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan float64)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeFloat642(inps[0], MergeFloat64(inps[1:]...))
	}
}

// mergeFloat642 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeFloat642(i1, i2 <-chan float64) (out <-chan float64) {
	cha := make(chan float64)
	go func(out chan<- float64, i1, i2 <-chan float64) {
		defer close(out)
		var (
			clos1, clos2 bool    // we found the chan closed
			buff1, buff2 bool    // we've read 'from', but not sent (yet)
			ok           bool    // did we read sucessfully?
			from1, from2 float64 // what we've read
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
