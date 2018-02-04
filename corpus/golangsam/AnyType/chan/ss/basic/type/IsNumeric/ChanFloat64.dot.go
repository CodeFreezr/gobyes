// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

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
func MakeFloat64Chan() (out chan float64) {
	return make(chan float64)
}

func sendFloat64(out chan<- float64, inp ...float64) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanFloat64 returns a channel to receive all inputs before close.
func ChanFloat64(inp ...float64) (out <-chan float64) {
	cha := make(chan float64)
	go sendFloat64(cha, inp...)
	return cha
}

func sendFloat64Slice(out chan<- float64, inp ...[]float64) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanFloat64Slice returns a channel to receive all inputs before close.
func ChanFloat64Slice(inp ...[]float64) (out <-chan float64) {
	cha := make(chan float64)
	go sendFloat64Slice(cha, inp...)
	return cha
}

func chanFloat64FuncNok(out chan<- float64, act func() (float64, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanFloat64FuncNok returns a channel to receive all results of act until nok before close.
func ChanFloat64FuncNok(act func() (float64, bool)) (out <-chan float64) {
	cha := make(chan float64)
	go chanFloat64FuncNok(cha, act)
	return cha
}

func chanFloat64FuncErr(out chan<- float64, act func() (float64, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanFloat64FuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFloat64FuncErr(act func() (float64, error)) (out <-chan float64) {
	cha := make(chan float64)
	go chanFloat64FuncErr(cha, act)
	return cha
}

func joinFloat64(done chan<- struct{}, out chan<- float64, inp ...float64) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinFloat64 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFloat64(out chan<- float64, inp ...float64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFloat64(cha, out, inp...)
	return cha
}

func joinFloat64Slice(done chan<- struct{}, out chan<- float64, inp ...[]float64) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinFloat64Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFloat64Slice(out chan<- float64, inp ...[]float64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFloat64Slice(cha, out, inp...)
	return cha
}

func joinFloat64Chan(done chan<- struct{}, out chan<- float64, inp <-chan float64) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFloat64Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFloat64Chan(out chan<- float64, inp <-chan float64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFloat64Chan(cha, out, inp)
	return cha
}

func doitFloat64(done chan<- struct{}, inp <-chan float64) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneFloat64 returns a channel to receive one signal before close after inp has been drained.
func DoneFloat64(inp <-chan float64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitFloat64(cha, inp)
	return cha
}

func doitFloat64Slice(done chan<- ([]float64), inp <-chan float64) {
	defer close(done)
	Float64S := []float64{}
	for i := range inp {
		Float64S = append(Float64S, i)
	}
	done <- Float64S
}

// DoneFloat64Slice returns a channel which will receive a slice
// of all the Float64s received on inp channel before close.
// Unlike DoneFloat64, a full slice is sent once, not just an event.
func DoneFloat64Slice(inp <-chan float64) (done <-chan ([]float64)) {
	cha := make(chan ([]float64))
	go doitFloat64Slice(cha, inp)
	return cha
}

func doitFloat64Func(done chan<- struct{}, inp <-chan float64, act func(a float64)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFloat64Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFloat64Func(inp <-chan float64, act func(a float64)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a float64) { return }
	}
	go doitFloat64Func(cha, inp, act)
	return cha
}

func pipeFloat64Buffer(out chan<- float64, inp <-chan float64) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeFloat64Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFloat64Buffer(inp <-chan float64, cap int) (out <-chan float64) {
	cha := make(chan float64, cap)
	go pipeFloat64Buffer(cha, inp)
	return cha
}

func pipeFloat64Func(out chan<- float64, inp <-chan float64, act func(a float64) float64) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFloat64Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFloat64Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFloat64Func(inp <-chan float64, act func(a float64) float64) (out <-chan float64) {
	cha := make(chan float64)
	if act == nil {
		act = func(a float64) float64 { return a }
	}
	go pipeFloat64Func(cha, inp, act)
	return cha
}

func pipeFloat64Fork(out1, out2 chan<- float64, inp <-chan float64) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFloat64Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFloat64Fork(inp <-chan float64) (out1, out2 <-chan float64) {
	cha1 := make(chan float64)
	cha2 := make(chan float64)
	go pipeFloat64Fork(cha1, cha2, inp)
	return cha1, cha2
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
