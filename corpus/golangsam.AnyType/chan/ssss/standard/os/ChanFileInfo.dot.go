// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

// MakeFileInfoChan returns a new open channel
// (simply a 'chan os.FileInfo' that is).
//
// Note: No 'FileInfo-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFileInfoPipelineStartsHere := MakeFileInfoChan()
//	// ... lot's of code to design and build Your favourite "myFileInfoWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFileInfoPipelineStartsHere <- drop
//	}
//	close(myFileInfoPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFileInfoBuffer) the channel is unbuffered.
//
func MakeFileInfoChan() chan os.FileInfo {
	return make(chan os.FileInfo)
}

// ChanFileInfo returns a channel to receive all inputs before close.
func ChanFileInfo(inp ...os.FileInfo) chan os.FileInfo {
	out := make(chan os.FileInfo)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanFileInfoSlice returns a channel to receive all inputs before close.
func ChanFileInfoSlice(inp ...[]os.FileInfo) chan os.FileInfo {
	out := make(chan os.FileInfo)
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

// ChanFileInfoFuncNok returns a channel to receive all results of act until nok before close.
func ChanFileInfoFuncNok(act func() (os.FileInfo, bool)) <-chan os.FileInfo {
	out := make(chan os.FileInfo)
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

// ChanFileInfoFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFileInfoFuncErr(act func() (os.FileInfo, error)) <-chan os.FileInfo {
	out := make(chan os.FileInfo)
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

// JoinFileInfo sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFileInfo(out chan<- os.FileInfo, inp ...os.FileInfo) chan struct{} {
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

// JoinFileInfoSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFileInfoSlice(out chan<- os.FileInfo, inp ...[]os.FileInfo) chan struct{} {
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

// JoinFileInfoChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFileInfoChan(out chan<- os.FileInfo, inp <-chan os.FileInfo) chan struct{} {
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

// DoneFileInfo returns a channel to receive one signal before close after inp has been drained.
func DoneFileInfo(inp <-chan os.FileInfo) chan struct{} {
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

// DoneFileInfoSlice returns a channel which will receive a slice
// of all the FileInfos received on inp channel before close.
// Unlike DoneFileInfo, a full slice is sent once, not just an event.
func DoneFileInfoSlice(inp <-chan os.FileInfo) chan []os.FileInfo {
	done := make(chan []os.FileInfo)
	go func() {
		defer close(done)
		FileInfoS := []os.FileInfo{}
		for i := range inp {
			FileInfoS = append(FileInfoS, i)
		}
		done <- FileInfoS
	}()
	return done
}

// DoneFileInfoFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFileInfoFunc(inp <-chan os.FileInfo, act func(a os.FileInfo)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a os.FileInfo) { return }
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

// PipeFileInfoBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFileInfoBuffer(inp <-chan os.FileInfo, cap int) chan os.FileInfo {
	out := make(chan os.FileInfo, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFileInfoFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFileInfoMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFileInfoFunc(inp <-chan os.FileInfo, act func(a os.FileInfo) os.FileInfo) chan os.FileInfo {
	out := make(chan os.FileInfo)
	if act == nil {
		act = func(a os.FileInfo) os.FileInfo { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFileInfoFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFileInfoFork(inp <-chan os.FileInfo) (chan os.FileInfo, chan os.FileInfo) {
	out1 := make(chan os.FileInfo)
	out2 := make(chan os.FileInfo)
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

// FileInfoTube is the signature for a pipe function.
type FileInfoTube func(inp <-chan os.FileInfo, out <-chan os.FileInfo)

// FileInfoDaisy returns a channel to receive all inp after having passed thru tube.
func FileInfoDaisy(inp <-chan os.FileInfo, tube FileInfoTube) (out <-chan os.FileInfo) {
	cha := make(chan os.FileInfo)
	go tube(inp, cha)
	return cha
}

// FileInfoDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func FileInfoDaisyChain(inp <-chan os.FileInfo, tubes ...FileInfoTube) (out <-chan os.FileInfo) {
	cha := inp
	for i := range tubes {
		cha = FileInfoDaisy(cha, tubes[i])
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
