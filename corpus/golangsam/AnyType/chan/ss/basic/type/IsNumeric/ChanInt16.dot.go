// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeInt16Chan returns a new open channel
// (simply a 'chan int16' that is).
//
// Note: No 'Int16-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myInt16PipelineStartsHere := MakeInt16Chan()
//	// ... lot's of code to design and build Your favourite "myInt16WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myInt16PipelineStartsHere <- drop
//	}
//	close(myInt16PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeInt16Buffer) the channel is unbuffered.
//
func MakeInt16Chan() (out chan int16) {
	return make(chan int16)
}

func sendInt16(out chan<- int16, inp ...int16) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanInt16 returns a channel to receive all inputs before close.
func ChanInt16(inp ...int16) (out <-chan int16) {
	cha := make(chan int16)
	go sendInt16(cha, inp...)
	return cha
}

func sendInt16Slice(out chan<- int16, inp ...[]int16) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanInt16Slice returns a channel to receive all inputs before close.
func ChanInt16Slice(inp ...[]int16) (out <-chan int16) {
	cha := make(chan int16)
	go sendInt16Slice(cha, inp...)
	return cha
}

func chanInt16FuncNok(out chan<- int16, act func() (int16, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanInt16FuncNok returns a channel to receive all results of act until nok before close.
func ChanInt16FuncNok(act func() (int16, bool)) (out <-chan int16) {
	cha := make(chan int16)
	go chanInt16FuncNok(cha, act)
	return cha
}

func chanInt16FuncErr(out chan<- int16, act func() (int16, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanInt16FuncErr returns a channel to receive all results of act until err != nil before close.
func ChanInt16FuncErr(act func() (int16, error)) (out <-chan int16) {
	cha := make(chan int16)
	go chanInt16FuncErr(cha, act)
	return cha
}

func joinInt16(done chan<- struct{}, out chan<- int16, inp ...int16) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinInt16 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt16(out chan<- int16, inp ...int16) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt16(cha, out, inp...)
	return cha
}

func joinInt16Slice(done chan<- struct{}, out chan<- int16, inp ...[]int16) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinInt16Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt16Slice(out chan<- int16, inp ...[]int16) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt16Slice(cha, out, inp...)
	return cha
}

func joinInt16Chan(done chan<- struct{}, out chan<- int16, inp <-chan int16) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinInt16Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt16Chan(out chan<- int16, inp <-chan int16) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt16Chan(cha, out, inp)
	return cha
}

func doitInt16(done chan<- struct{}, inp <-chan int16) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneInt16 returns a channel to receive one signal before close after inp has been drained.
func DoneInt16(inp <-chan int16) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitInt16(cha, inp)
	return cha
}

func doitInt16Slice(done chan<- ([]int16), inp <-chan int16) {
	defer close(done)
	Int16S := []int16{}
	for i := range inp {
		Int16S = append(Int16S, i)
	}
	done <- Int16S
}

// DoneInt16Slice returns a channel which will receive a slice
// of all the Int16s received on inp channel before close.
// Unlike DoneInt16, a full slice is sent once, not just an event.
func DoneInt16Slice(inp <-chan int16) (done <-chan ([]int16)) {
	cha := make(chan ([]int16))
	go doitInt16Slice(cha, inp)
	return cha
}

func doitInt16Func(done chan<- struct{}, inp <-chan int16, act func(a int16)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneInt16Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneInt16Func(inp <-chan int16, act func(a int16)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a int16) { return }
	}
	go doitInt16Func(cha, inp, act)
	return cha
}

func pipeInt16Buffer(out chan<- int16, inp <-chan int16) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeInt16Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeInt16Buffer(inp <-chan int16, cap int) (out <-chan int16) {
	cha := make(chan int16, cap)
	go pipeInt16Buffer(cha, inp)
	return cha
}

func pipeInt16Func(out chan<- int16, inp <-chan int16, act func(a int16) int16) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeInt16Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeInt16Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeInt16Func(inp <-chan int16, act func(a int16) int16) (out <-chan int16) {
	cha := make(chan int16)
	if act == nil {
		act = func(a int16) int16 { return a }
	}
	go pipeInt16Func(cha, inp, act)
	return cha
}

func pipeInt16Fork(out1, out2 chan<- int16, inp <-chan int16) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeInt16Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeInt16Fork(inp <-chan int16) (out1, out2 <-chan int16) {
	cha1 := make(chan int16)
	cha2 := make(chan int16)
	go pipeInt16Fork(cha1, cha2, inp)
	return cha1, cha2
}

// Int16Tube is the signature for a pipe function.
type Int16Tube func(inp <-chan int16, out <-chan int16)

// Int16Daisy returns a channel to receive all inp after having passed thru tube.
func Int16Daisy(inp <-chan int16, tube Int16Tube) (out <-chan int16) {
	cha := make(chan int16)
	go tube(inp, cha)
	return cha
}

// Int16DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func Int16DaisyChain(inp <-chan int16, tubes ...Int16Tube) (out <-chan int16) {
	cha := inp
	for i := range tubes {
		cha = Int16Daisy(cha, tubes[i])
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
