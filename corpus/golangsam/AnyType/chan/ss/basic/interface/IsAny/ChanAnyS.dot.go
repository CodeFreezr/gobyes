// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsAny

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeAnySChan returns a new open channel
// (simply a 'chan []interface{}' that is).
//
// Note: No 'AnyS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myAnySPipelineStartsHere := MakeAnySChan()
//	// ... lot's of code to design and build Your favourite "myAnySWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myAnySPipelineStartsHere <- drop
//	}
//	close(myAnySPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeAnySBuffer) the channel is unbuffered.
//
func MakeAnySChan() (out chan []interface{}) {
	return make(chan []interface{})
}

func sendAnyS(out chan<- []interface{}, inp ...[]interface{}) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanAnyS returns a channel to receive all inputs before close.
func ChanAnyS(inp ...[]interface{}) (out <-chan []interface{}) {
	cha := make(chan []interface{})
	go sendAnyS(cha, inp...)
	return cha
}

func sendAnySSlice(out chan<- []interface{}, inp ...[][]interface{}) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanAnySSlice returns a channel to receive all inputs before close.
func ChanAnySSlice(inp ...[][]interface{}) (out <-chan []interface{}) {
	cha := make(chan []interface{})
	go sendAnySSlice(cha, inp...)
	return cha
}

func chanAnySFuncNok(out chan<- []interface{}, act func() ([]interface{}, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanAnySFuncNok returns a channel to receive all results of act until nok before close.
func ChanAnySFuncNok(act func() ([]interface{}, bool)) (out <-chan []interface{}) {
	cha := make(chan []interface{})
	go chanAnySFuncNok(cha, act)
	return cha
}

func chanAnySFuncErr(out chan<- []interface{}, act func() ([]interface{}, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanAnySFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanAnySFuncErr(act func() ([]interface{}, error)) (out <-chan []interface{}) {
	cha := make(chan []interface{})
	go chanAnySFuncErr(cha, act)
	return cha
}

func joinAnyS(done chan<- struct{}, out chan<- []interface{}, inp ...[]interface{}) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinAnyS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinAnyS(out chan<- []interface{}, inp ...[]interface{}) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinAnyS(cha, out, inp...)
	return cha
}

func joinAnySSlice(done chan<- struct{}, out chan<- []interface{}, inp ...[][]interface{}) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinAnySSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinAnySSlice(out chan<- []interface{}, inp ...[][]interface{}) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinAnySSlice(cha, out, inp...)
	return cha
}

func joinAnySChan(done chan<- struct{}, out chan<- []interface{}, inp <-chan []interface{}) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinAnySChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinAnySChan(out chan<- []interface{}, inp <-chan []interface{}) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinAnySChan(cha, out, inp)
	return cha
}

func doitAnyS(done chan<- struct{}, inp <-chan []interface{}) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneAnyS returns a channel to receive one signal before close after inp has been drained.
func DoneAnyS(inp <-chan []interface{}) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitAnyS(cha, inp)
	return cha
}

func doitAnySSlice(done chan<- ([][]interface{}), inp <-chan []interface{}) {
	defer close(done)
	AnySS := [][]interface{}{}
	for i := range inp {
		AnySS = append(AnySS, i)
	}
	done <- AnySS
}

// DoneAnySSlice returns a channel which will receive a slice
// of all the AnySs received on inp channel before close.
// Unlike DoneAnyS, a full slice is sent once, not just an event.
func DoneAnySSlice(inp <-chan []interface{}) (done <-chan ([][]interface{})) {
	cha := make(chan ([][]interface{}))
	go doitAnySSlice(cha, inp)
	return cha
}

func doitAnySFunc(done chan<- struct{}, inp <-chan []interface{}, act func(a []interface{})) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneAnySFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneAnySFunc(inp <-chan []interface{}, act func(a []interface{})) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []interface{}) { return }
	}
	go doitAnySFunc(cha, inp, act)
	return cha
}

func pipeAnySBuffer(out chan<- []interface{}, inp <-chan []interface{}) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeAnySBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeAnySBuffer(inp <-chan []interface{}, cap int) (out <-chan []interface{}) {
	cha := make(chan []interface{}, cap)
	go pipeAnySBuffer(cha, inp)
	return cha
}

func pipeAnySFunc(out chan<- []interface{}, inp <-chan []interface{}, act func(a []interface{}) []interface{}) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeAnySFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeAnySMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeAnySFunc(inp <-chan []interface{}, act func(a []interface{}) []interface{}) (out <-chan []interface{}) {
	cha := make(chan []interface{})
	if act == nil {
		act = func(a []interface{}) []interface{} { return a }
	}
	go pipeAnySFunc(cha, inp, act)
	return cha
}

func pipeAnySFork(out1, out2 chan<- []interface{}, inp <-chan []interface{}) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeAnySFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeAnySFork(inp <-chan []interface{}) (out1, out2 <-chan []interface{}) {
	cha1 := make(chan []interface{})
	cha2 := make(chan []interface{})
	go pipeAnySFork(cha1, cha2, inp)
	return cha1, cha2
}

// AnySTube is the signature for a pipe function.
type AnySTube func(inp <-chan []interface{}, out <-chan []interface{})

// AnySDaisy returns a channel to receive all inp after having passed thru tube.
func AnySDaisy(inp <-chan []interface{}, tube AnySTube) (out <-chan []interface{}) {
	cha := make(chan []interface{})
	go tube(inp, cha)
	return cha
}

// AnySDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func AnySDaisyChain(inp <-chan []interface{}, tubes ...AnySTube) (out <-chan []interface{}) {
	cha := inp
	for i := range tubes {
		cha = AnySDaisy(cha, tubes[i])
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
