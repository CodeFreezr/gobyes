// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeFloat32Chan returns a new open channel
// (simply a 'chan float32' that is).
//
// Note: No 'Float32-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFloat32PipelineStartsHere := MakeFloat32Chan()
//	// ... lot's of code to design and build Your favourite "myFloat32WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFloat32PipelineStartsHere <- drop
//	}
//	close(myFloat32PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFloat32Buffer) the channel is unbuffered.
//
func MakeFloat32Chan() chan float32 {
	return make(chan float32)
}

// ChanFloat32 returns a channel to receive all inputs before close.
func ChanFloat32(inp ...float32) chan float32 {
	out := make(chan float32)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanFloat32Slice returns a channel to receive all inputs before close.
func ChanFloat32Slice(inp ...[]float32) chan float32 {
	out := make(chan float32)
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

// ChanFloat32FuncNok returns a channel to receive all results of act until nok before close.
func ChanFloat32FuncNok(act func() (float32, bool)) <-chan float32 {
	out := make(chan float32)
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

// ChanFloat32FuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFloat32FuncErr(act func() (float32, error)) <-chan float32 {
	out := make(chan float32)
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

// JoinFloat32 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFloat32(out chan<- float32, inp ...float32) chan struct{} {
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

// JoinFloat32Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFloat32Slice(out chan<- float32, inp ...[]float32) chan struct{} {
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

// JoinFloat32Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFloat32Chan(out chan<- float32, inp <-chan float32) chan struct{} {
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

// DoneFloat32 returns a channel to receive one signal before close after inp has been drained.
func DoneFloat32(inp <-chan float32) chan struct{} {
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

// DoneFloat32Slice returns a channel which will receive a slice
// of all the Float32s received on inp channel before close.
// Unlike DoneFloat32, a full slice is sent once, not just an event.
func DoneFloat32Slice(inp <-chan float32) chan []float32 {
	done := make(chan []float32)
	go func() {
		defer close(done)
		Float32S := []float32{}
		for i := range inp {
			Float32S = append(Float32S, i)
		}
		done <- Float32S
	}()
	return done
}

// DoneFloat32Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFloat32Func(inp <-chan float32, act func(a float32)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a float32) { return }
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

// PipeFloat32Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFloat32Buffer(inp <-chan float32, cap int) chan float32 {
	out := make(chan float32, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFloat32Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFloat32Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFloat32Func(inp <-chan float32, act func(a float32) float32) chan float32 {
	out := make(chan float32)
	if act == nil {
		act = func(a float32) float32 { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFloat32Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFloat32Fork(inp <-chan float32) (chan float32, chan float32) {
	out1 := make(chan float32)
	out2 := make(chan float32)
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

// Float32Tube is the signature for a pipe function.
type Float32Tube func(inp <-chan float32, out <-chan float32)

// Float32Daisy returns a channel to receive all inp after having passed thru tube.
func Float32Daisy(inp <-chan float32, tube Float32Tube) (out <-chan float32) {
	cha := make(chan float32)
	go tube(inp, cha)
	return cha
}

// Float32DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func Float32DaisyChain(inp <-chan float32, tubes ...Float32Tube) (out <-chan float32) {
	cha := inp
	for i := range tubes {
		cha = Float32Daisy(cha, tubes[i])
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
