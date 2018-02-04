// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeIntChan returns a new open channel
// (simply a 'chan int' that is).
//
// Note: No 'Int-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myIntPipelineStartsHere := MakeIntChan()
//	// ... lot's of code to design and build Your favourite "myIntWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myIntPipelineStartsHere <- drop
//	}
//	close(myIntPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeIntBuffer) the channel is unbuffered.
//
func MakeIntChan() chan int {
	return make(chan int)
}

// ChanInt returns a channel to receive all inputs before close.
func ChanInt(inp ...int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanIntSlice returns a channel to receive all inputs before close.
func ChanIntSlice(inp ...[]int) chan int {
	out := make(chan int)
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

// ChanIntFuncNok returns a channel to receive all results of act until nok before close.
func ChanIntFuncNok(act func() (int, bool)) <-chan int {
	out := make(chan int)
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

// ChanIntFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanIntFuncErr(act func() (int, error)) <-chan int {
	out := make(chan int)
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

// JoinInt sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt(out chan<- int, inp ...int) chan struct{} {
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

// JoinIntSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinIntSlice(out chan<- int, inp ...[]int) chan struct{} {
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

// JoinIntChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinIntChan(out chan<- int, inp <-chan int) chan struct{} {
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

// DoneInt returns a channel to receive one signal before close after inp has been drained.
func DoneInt(inp <-chan int) chan struct{} {
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

// DoneIntSlice returns a channel which will receive a slice
// of all the Ints received on inp channel before close.
// Unlike DoneInt, a full slice is sent once, not just an event.
func DoneIntSlice(inp <-chan int) chan []int {
	done := make(chan []int)
	go func() {
		defer close(done)
		IntS := []int{}
		for i := range inp {
			IntS = append(IntS, i)
		}
		done <- IntS
	}()
	return done
}

// DoneIntFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneIntFunc(inp <-chan int, act func(a int)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a int) { return }
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

// PipeIntBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeIntBuffer(inp <-chan int, cap int) chan int {
	out := make(chan int, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeIntFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeIntMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeIntFunc(inp <-chan int, act func(a int) int) chan int {
	out := make(chan int)
	if act == nil {
		act = func(a int) int { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeIntFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeIntFork(inp <-chan int) (chan int, chan int) {
	out1 := make(chan int)
	out2 := make(chan int)
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

// IntTube is the signature for a pipe function.
type IntTube func(inp <-chan int, out <-chan int)

// IntDaisy returns a channel to receive all inp after having passed thru tube.
func IntDaisy(inp <-chan int, tube IntTube) (out <-chan int) {
	cha := make(chan int)
	go tube(inp, cha)
	return cha
}

// IntDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func IntDaisyChain(inp <-chan int, tubes ...IntTube) (out <-chan int) {
	cha := inp
	for i := range tubes {
		cha = IntDaisy(cha, tubes[i])
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
