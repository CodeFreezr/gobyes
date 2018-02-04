// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pat

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakePatSChan returns a new open channel
// (simply a 'chan []struct{}' that is).
//
// Note: No 'PatS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPatSPipelineStartsHere := MakePatSChan()
//	// ... lot's of code to design and build Your favourite "myPatSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPatSPipelineStartsHere <- drop
//	}
//	close(myPatSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePatSBuffer) the channel is unbuffered.
//
func MakePatSChan() (out chan []struct{}) {
	return make(chan []struct{})
}

// ChanPatS returns a channel to receive all inputs before close.
func ChanPatS(inp ...[]struct{}) (out <-chan []struct{}) {
	cha := make(chan []struct{})
	go func(out chan<- []struct{}, inp ...[]struct{}) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanPatSSlice returns a channel to receive all inputs before close.
func ChanPatSSlice(inp ...[][]struct{}) (out <-chan []struct{}) {
	cha := make(chan []struct{})
	go func(out chan<- []struct{}, inp ...[][]struct{}) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanPatSFuncNok returns a channel to receive all results of act until nok before close.
func ChanPatSFuncNok(act func() ([]struct{}, bool)) (out <-chan []struct{}) {
	cha := make(chan []struct{})
	go func(out chan<- []struct{}, act func() ([]struct{}, bool)) {
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

// ChanPatSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanPatSFuncErr(act func() ([]struct{}, error)) (out <-chan []struct{}) {
	cha := make(chan []struct{})
	go func(out chan<- []struct{}, act func() ([]struct{}, error)) {
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

// JoinPatS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPatS(out chan<- []struct{}, inp ...[]struct{}) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []struct{}, inp ...[]struct{}) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinPatSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPatSSlice(out chan<- []struct{}, inp ...[][]struct{}) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []struct{}, inp ...[][]struct{}) {
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

// JoinPatSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPatSChan(out chan<- []struct{}, inp <-chan []struct{}) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []struct{}, inp <-chan []struct{}) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DonePatS returns a channel to receive one signal before close after inp has been drained.
func DonePatS(inp <-chan []struct{}) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan []struct{}) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DonePatSSlice returns a channel which will receive a slice
// of all the PatSs received on inp channel before close.
// Unlike DonePatS, a full slice is sent once, not just an event.
func DonePatSSlice(inp <-chan []struct{}) (done <-chan [][]struct{}) {
	cha := make(chan [][]struct{})
	go func(inp <-chan []struct{}, done chan<- [][]struct{}) {
		defer close(done)
		PatSS := [][]struct{}{}
		for i := range inp {
			PatSS = append(PatSS, i)
		}
		done <- PatSS
	}(inp, cha)
	return cha
}

// DonePatSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePatSFunc(inp <-chan []struct{}, act func(a []struct{})) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []struct{}) { return }
	}
	go func(done chan<- struct{}, inp <-chan []struct{}, act func(a []struct{})) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipePatSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePatSBuffer(inp <-chan []struct{}, cap int) (out <-chan []struct{}) {
	cha := make(chan []struct{}, cap)
	go func(out chan<- []struct{}, inp <-chan []struct{}) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipePatSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePatSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePatSFunc(inp <-chan []struct{}, act func(a []struct{}) []struct{}) (out <-chan []struct{}) {
	cha := make(chan []struct{})
	if act == nil {
		act = func(a []struct{}) []struct{} { return a }
	}
	go func(out chan<- []struct{}, inp <-chan []struct{}, act func(a []struct{}) []struct{}) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipePatSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePatSFork(inp <-chan []struct{}) (out1, out2 <-chan []struct{}) {
	cha1 := make(chan []struct{})
	cha2 := make(chan []struct{})
	go func(out1, out2 chan<- []struct{}, inp <-chan []struct{}) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// PatSTube is the signature for a pipe function.
type PatSTube func(inp <-chan []struct{}, out <-chan []struct{})

// PatSDaisy returns a channel to receive all inp after having passed thru tube.
func PatSDaisy(inp <-chan []struct{}, tube PatSTube) (out <-chan []struct{}) {
	cha := make(chan []struct{})
	go tube(inp, cha)
	return cha
}

// PatSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func PatSDaisyChain(inp <-chan []struct{}, tubes ...PatSTube) (out <-chan []struct{}) {
	cha := inp
	for i := range tubes {
		cha = PatSDaisy(cha, tubes[i])
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
