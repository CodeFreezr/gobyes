// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeStringSChan returns a new open channel
// (simply a 'chan []string' that is).
//
// Note: No 'StringS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myStringSPipelineStartsHere := MakeStringSChan()
//	// ... lot's of code to design and build Your favourite "myStringSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myStringSPipelineStartsHere <- drop
//	}
//	close(myStringSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeStringSBuffer) the channel is unbuffered.
//
func MakeStringSChan() (out chan []string) {
	return make(chan []string)
}

// ChanStringS returns a channel to receive all inputs before close.
func ChanStringS(inp ...[]string) (out <-chan []string) {
	cha := make(chan []string)
	go func(out chan<- []string, inp ...[]string) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanStringSSlice returns a channel to receive all inputs before close.
func ChanStringSSlice(inp ...[][]string) (out <-chan []string) {
	cha := make(chan []string)
	go func(out chan<- []string, inp ...[][]string) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanStringSFuncNok returns a channel to receive all results of act until nok before close.
func ChanStringSFuncNok(act func() ([]string, bool)) (out <-chan []string) {
	cha := make(chan []string)
	go func(out chan<- []string, act func() ([]string, bool)) {
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

// ChanStringSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanStringSFuncErr(act func() ([]string, error)) (out <-chan []string) {
	cha := make(chan []string)
	go func(out chan<- []string, act func() ([]string, error)) {
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

// JoinStringS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinStringS(out chan<- []string, inp ...[]string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []string, inp ...[]string) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinStringSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinStringSSlice(out chan<- []string, inp ...[][]string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []string, inp ...[][]string) {
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

// JoinStringSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinStringSChan(out chan<- []string, inp <-chan []string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []string, inp <-chan []string) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneStringS returns a channel to receive one signal before close after inp has been drained.
func DoneStringS(inp <-chan []string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan []string) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneStringSSlice returns a channel which will receive a slice
// of all the StringSs received on inp channel before close.
// Unlike DoneStringS, a full slice is sent once, not just an event.
func DoneStringSSlice(inp <-chan []string) (done <-chan [][]string) {
	cha := make(chan [][]string)
	go func(inp <-chan []string, done chan<- [][]string) {
		defer close(done)
		StringSS := [][]string{}
		for i := range inp {
			StringSS = append(StringSS, i)
		}
		done <- StringSS
	}(inp, cha)
	return cha
}

// DoneStringSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneStringSFunc(inp <-chan []string, act func(a []string)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []string) { return }
	}
	go func(done chan<- struct{}, inp <-chan []string, act func(a []string)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeStringSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeStringSBuffer(inp <-chan []string, cap int) (out <-chan []string) {
	cha := make(chan []string, cap)
	go func(out chan<- []string, inp <-chan []string) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeStringSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeStringSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeStringSFunc(inp <-chan []string, act func(a []string) []string) (out <-chan []string) {
	cha := make(chan []string)
	if act == nil {
		act = func(a []string) []string { return a }
	}
	go func(out chan<- []string, inp <-chan []string, act func(a []string) []string) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeStringSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeStringSFork(inp <-chan []string) (out1, out2 <-chan []string) {
	cha1 := make(chan []string)
	cha2 := make(chan []string)
	go func(out1, out2 chan<- []string, inp <-chan []string) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// StringSTube is the signature for a pipe function.
type StringSTube func(inp <-chan []string, out <-chan []string)

// StringSDaisy returns a channel to receive all inp after having passed thru tube.
func StringSDaisy(inp <-chan []string, tube StringSTube) (out <-chan []string) {
	cha := make(chan []string)
	go tube(inp, cha)
	return cha
}

// StringSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func StringSDaisyChain(inp <-chan []string, tubes ...StringSTube) (out <-chan []string) {
	cha := inp
	for i := range tubes {
		cha = StringSDaisy(cha, tubes[i])
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
