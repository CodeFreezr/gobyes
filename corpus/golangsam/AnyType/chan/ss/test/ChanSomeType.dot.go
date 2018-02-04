// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeSomeTypeChan returns a new open channel
// (simply a 'chan SomeType' that is).
//
// Note: No 'SomeType-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var mySomeTypePipelineStartsHere := MakeSomeTypeChan()
//	// ... lot's of code to design and build Your favourite "mySomeTypeWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		mySomeTypePipelineStartsHere <- drop
//	}
//	close(mySomeTypePipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeSomeTypeBuffer) the channel is unbuffered.
//
func MakeSomeTypeChan() (out chan SomeType) {
	return make(chan SomeType)
}

func sendSomeType(out chan<- SomeType, inp ...SomeType) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanSomeType returns a channel to receive all inputs before close.
func ChanSomeType(inp ...SomeType) (out <-chan SomeType) {
	cha := make(chan SomeType)
	go sendSomeType(cha, inp...)
	return cha
}

func sendSomeTypeSlice(out chan<- SomeType, inp ...[]SomeType) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanSomeTypeSlice returns a channel to receive all inputs before close.
func ChanSomeTypeSlice(inp ...[]SomeType) (out <-chan SomeType) {
	cha := make(chan SomeType)
	go sendSomeTypeSlice(cha, inp...)
	return cha
}

func chanSomeTypeFuncNok(out chan<- SomeType, act func() (SomeType, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanSomeTypeFuncNok returns a channel to receive all results of act until nok before close.
func ChanSomeTypeFuncNok(act func() (SomeType, bool)) (out <-chan SomeType) {
	cha := make(chan SomeType)
	go chanSomeTypeFuncNok(cha, act)
	return cha
}

func chanSomeTypeFuncErr(out chan<- SomeType, act func() (SomeType, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanSomeTypeFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanSomeTypeFuncErr(act func() (SomeType, error)) (out <-chan SomeType) {
	cha := make(chan SomeType)
	go chanSomeTypeFuncErr(cha, act)
	return cha
}

func joinSomeType(done chan<- struct{}, out chan<- SomeType, inp ...SomeType) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinSomeType sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSomeType(out chan<- SomeType, inp ...SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSomeType(cha, out, inp...)
	return cha
}

func joinSomeTypeSlice(done chan<- struct{}, out chan<- SomeType, inp ...[]SomeType) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinSomeTypeSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSomeTypeSlice(out chan<- SomeType, inp ...[]SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSomeTypeSlice(cha, out, inp...)
	return cha
}

func joinSomeTypeChan(done chan<- struct{}, out chan<- SomeType, inp <-chan SomeType) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinSomeTypeChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSomeTypeChan(out chan<- SomeType, inp <-chan SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSomeTypeChan(cha, out, inp)
	return cha
}

func doitSomeType(done chan<- struct{}, inp <-chan SomeType) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneSomeType returns a channel to receive one signal before close after inp has been drained.
func DoneSomeType(inp <-chan SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitSomeType(cha, inp)
	return cha
}

func doitSomeTypeSlice(done chan<- ([]SomeType), inp <-chan SomeType) {
	defer close(done)
	SomeTypeS := []SomeType{}
	for i := range inp {
		SomeTypeS = append(SomeTypeS, i)
	}
	done <- SomeTypeS
}

// DoneSomeTypeSlice returns a channel which will receive a slice
// of all the SomeTypes received on inp channel before close.
// Unlike DoneSomeType, a full slice is sent once, not just an event.
func DoneSomeTypeSlice(inp <-chan SomeType) (done <-chan ([]SomeType)) {
	cha := make(chan ([]SomeType))
	go doitSomeTypeSlice(cha, inp)
	return cha
}

func doitSomeTypeFunc(done chan<- struct{}, inp <-chan SomeType, act func(a SomeType)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneSomeTypeFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSomeTypeFunc(inp <-chan SomeType, act func(a SomeType)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a SomeType) { return }
	}
	go doitSomeTypeFunc(cha, inp, act)
	return cha
}

func pipeSomeTypeBuffer(out chan<- SomeType, inp <-chan SomeType) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeSomeTypeBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSomeTypeBuffer(inp <-chan SomeType, cap int) (out <-chan SomeType) {
	cha := make(chan SomeType, cap)
	go pipeSomeTypeBuffer(cha, inp)
	return cha
}

func pipeSomeTypeFunc(out chan<- SomeType, inp <-chan SomeType, act func(a SomeType) SomeType) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeSomeTypeFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSomeTypeMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSomeTypeFunc(inp <-chan SomeType, act func(a SomeType) SomeType) (out <-chan SomeType) {
	cha := make(chan SomeType)
	if act == nil {
		act = func(a SomeType) SomeType { return a }
	}
	go pipeSomeTypeFunc(cha, inp, act)
	return cha
}

func pipeSomeTypeFork(out1, out2 chan<- SomeType, inp <-chan SomeType) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeSomeTypeFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSomeTypeFork(inp <-chan SomeType) (out1, out2 <-chan SomeType) {
	cha1 := make(chan SomeType)
	cha2 := make(chan SomeType)
	go pipeSomeTypeFork(cha1, cha2, inp)
	return cha1, cha2
}

// SomeTypeTube is the signature for a pipe function.
type SomeTypeTube func(inp <-chan SomeType, out <-chan SomeType)

// SomeTypeDaisy returns a channel to receive all inp after having passed thru tube.
func SomeTypeDaisy(inp <-chan SomeType, tube SomeTypeTube) (out <-chan SomeType) {
	cha := make(chan SomeType)
	go tube(inp, cha)
	return cha
}

// SomeTypeDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func SomeTypeDaisyChain(inp <-chan SomeType, tubes ...SomeTypeTube) (out <-chan SomeType) {
	cha := inp
	for i := range tubes {
		cha = SomeTypeDaisy(cha, tubes[i])
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
