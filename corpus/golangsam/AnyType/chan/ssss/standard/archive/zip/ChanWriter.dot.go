// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	zip "archive/zip"
)

// MakeWriterChan returns a new open channel
// (simply a 'chan *zip.Writer' that is).
//
// Note: No 'Writer-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myWriterPipelineStartsHere := MakeWriterChan()
//	// ... lot's of code to design and build Your favourite "myWriterWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myWriterPipelineStartsHere <- drop
//	}
//	close(myWriterPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeWriterBuffer) the channel is unbuffered.
//
func MakeWriterChan() chan *zip.Writer {
	return make(chan *zip.Writer)
}

// ChanWriter returns a channel to receive all inputs before close.
func ChanWriter(inp ...*zip.Writer) chan *zip.Writer {
	out := make(chan *zip.Writer)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanWriterSlice returns a channel to receive all inputs before close.
func ChanWriterSlice(inp ...[]*zip.Writer) chan *zip.Writer {
	out := make(chan *zip.Writer)
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

// ChanWriterFuncNil returns a channel to receive all results of act until nil before close.
func ChanWriterFuncNil(act func() *zip.Writer) <-chan *zip.Writer {
	out := make(chan *zip.Writer)
	go func() {
		defer close(out)
		for {
			res := act() // Apply action
			if res == nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanWriterFuncNok returns a channel to receive all results of act until nok before close.
func ChanWriterFuncNok(act func() (*zip.Writer, bool)) <-chan *zip.Writer {
	out := make(chan *zip.Writer)
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

// ChanWriterFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanWriterFuncErr(act func() (*zip.Writer, error)) <-chan *zip.Writer {
	out := make(chan *zip.Writer)
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

// JoinWriter sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriter(out chan<- *zip.Writer, inp ...*zip.Writer) chan struct{} {
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

// JoinWriterSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriterSlice(out chan<- *zip.Writer, inp ...[]*zip.Writer) chan struct{} {
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

// JoinWriterChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriterChan(out chan<- *zip.Writer, inp <-chan *zip.Writer) chan struct{} {
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

// DoneWriter returns a channel to receive one signal before close after inp has been drained.
func DoneWriter(inp <-chan *zip.Writer) chan struct{} {
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

// DoneWriterSlice returns a channel which will receive a slice
// of all the Writers received on inp channel before close.
// Unlike DoneWriter, a full slice is sent once, not just an event.
func DoneWriterSlice(inp <-chan *zip.Writer) chan []*zip.Writer {
	done := make(chan []*zip.Writer)
	go func() {
		defer close(done)
		WriterS := []*zip.Writer{}
		for i := range inp {
			WriterS = append(WriterS, i)
		}
		done <- WriterS
	}()
	return done
}

// DoneWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriterFunc(inp <-chan *zip.Writer, act func(a *zip.Writer)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *zip.Writer) { return }
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

// PipeWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriterBuffer(inp <-chan *zip.Writer, cap int) chan *zip.Writer {
	out := make(chan *zip.Writer, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriterFunc(inp <-chan *zip.Writer, act func(a *zip.Writer) *zip.Writer) chan *zip.Writer {
	out := make(chan *zip.Writer)
	if act == nil {
		act = func(a *zip.Writer) *zip.Writer { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriterFork(inp <-chan *zip.Writer) (chan *zip.Writer, chan *zip.Writer) {
	out1 := make(chan *zip.Writer)
	out2 := make(chan *zip.Writer)
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

// WriterTube is the signature for a pipe function.
type WriterTube func(inp <-chan *zip.Writer, out <-chan *zip.Writer)

// WriterDaisy returns a channel to receive all inp after having passed thru tube.
func WriterDaisy(inp <-chan *zip.Writer, tube WriterTube) (out <-chan *zip.Writer) {
	cha := make(chan *zip.Writer)
	go tube(inp, cha)
	return cha
}

// WriterDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func WriterDaisyChain(inp <-chan *zip.Writer, tubes ...WriterTube) (out <-chan *zip.Writer) {
	cha := inp
	for i := range tubes {
		cha = WriterDaisy(cha, tubes[i])
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
