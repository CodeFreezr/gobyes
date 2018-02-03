// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeStringChan returns a new open channel
// (simply a 'chan string' that is).
//
// Note: No 'String-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myStringPipelineStartsHere := MakeStringChan()
//	// ... lot's of code to design and build Your favourite "myStringWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myStringPipelineStartsHere <- drop
//	}
//	close(myStringPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeStringBuffer) the channel is unbuffered.
//
func MakeStringChan() (out chan string) {
	return make(chan string)
}

// ChanString returns a channel to receive all inputs before close.
func ChanString(inp ...string) (out <-chan string) {
	cha := make(chan string)
	go func(out chan<- string, inp ...string) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanStringSlice returns a channel to receive all inputs before close.
func ChanStringSlice(inp ...[]string) (out <-chan string) {
	cha := make(chan string)
	go func(out chan<- string, inp ...[]string) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanStringFuncNok returns a channel to receive all results of act until nok before close.
func ChanStringFuncNok(act func() (string, bool)) (out <-chan string) {
	cha := make(chan string)
	go func(out chan<- string, act func() (string, bool)) {
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

// ChanStringFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanStringFuncErr(act func() (string, error)) (out <-chan string) {
	cha := make(chan string)
	go func(out chan<- string, act func() (string, error)) {
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

// JoinString sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinString(out chan<- string, inp ...string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- string, inp ...string) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinStringSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinStringSlice(out chan<- string, inp ...[]string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- string, inp ...[]string) {
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

// JoinStringChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinStringChan(out chan<- string, inp <-chan string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- string, inp <-chan string) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneString returns a channel to receive one signal before close after inp has been drained.
func DoneString(inp <-chan string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan string) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneStringSlice returns a channel which will receive a slice
// of all the Strings received on inp channel before close.
// Unlike DoneString, a full slice is sent once, not just an event.
func DoneStringSlice(inp <-chan string) (done <-chan []string) {
	cha := make(chan []string)
	go func(inp <-chan string, done chan<- []string) {
		defer close(done)
		StringS := []string{}
		for i := range inp {
			StringS = append(StringS, i)
		}
		done <- StringS
	}(inp, cha)
	return cha
}

// DoneStringFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneStringFunc(inp <-chan string, act func(a string)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a string) { return }
	}
	go func(done chan<- struct{}, inp <-chan string, act func(a string)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeStringBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeStringBuffer(inp <-chan string, cap int) (out <-chan string) {
	cha := make(chan string, cap)
	go func(out chan<- string, inp <-chan string) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeStringFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeStringMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeStringFunc(inp <-chan string, act func(a string) string) (out <-chan string) {
	cha := make(chan string)
	if act == nil {
		act = func(a string) string { return a }
	}
	go func(out chan<- string, inp <-chan string, act func(a string) string) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeStringFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeStringFork(inp <-chan string) (out1, out2 <-chan string) {
	cha1 := make(chan string)
	cha2 := make(chan string)
	go func(out1, out2 chan<- string, inp <-chan string) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// StringTube is the signature for a pipe function.
type StringTube func(inp <-chan string, out <-chan string)

// StringDaisy returns a channel to receive all inp after having passed thru tube.
func StringDaisy(inp <-chan string, tube StringTube) (out <-chan string) {
	cha := make(chan string)
	go tube(inp, cha)
	return cha
}

// StringDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func StringDaisyChain(inp <-chan string, tubes ...StringTube) (out <-chan string) {
	cha := inp
	for i := range tubes {
		cha = StringDaisy(cha, tubes[i])
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

// MergeString returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Stringchannel is returned.
func MergeString(inps ...<-chan string) (out <-chan string) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan string)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeString2(inps[0], MergeString(inps[1:]...))
	}
}

// mergeString2 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeString2(i1, i2 <-chan string) (out <-chan string) {
	cha := make(chan string)
	go func(out chan<- string, i1, i2 <-chan string) {
		defer close(out)
		var (
			clos1, clos2 bool   // we found the chan closed
			buff1, buff2 bool   // we've read 'from', but not sent (yet)
			ok           bool   // did we read sucessfully?
			from1, from2 string // what we've read
		)

		for !clos1 || !clos2 {

			if !clos1 && !buff1 {
				if from1, ok = <-i1; ok {
					buff1 = true
				} else {
					clos1 = true
				}
			}

			if !clos2 && !buff2 {
				if from2, ok = <-i2; ok {
					buff2 = true
				} else {
					clos2 = true
				}
			}

			if clos1 && !buff1 {
				from1 = from2
			}
			if clos2 && !buff2 {
				from2 = from1
			}

			if from1 < from2 {
				out <- from1
				buff1 = false
			} else if from2 < from1 {
				out <- from2
				buff2 = false
			} else {
				out <- from1 // == from2
				buff1 = false
				buff2 = false
			}
		}
	}(cha, i1, i2)
	return cha
}

// Note: merge2 is not my own. Just: I forgot where found it - please accept my apologies.
// I'd love to learn about it's origin/author, so I can give credit. Any hint is highly appreciated!
