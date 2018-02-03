// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Note: SendProxyPipeWriter imports "container/ring" for the expanding buffer.
import (
	"container/ring"
	"io"
)

// PipeWriterCAP is the capacity of the buffered proxy channel
const PipeWriterCAP = 10

// PipeWriterQUE is the allocated size of the circular queue
const PipeWriterQUE = 16

// SendProxyPipeWriter returns a channel to serve as a sending proxy to 'out'.
// Uses a goroutine to receive values from 'out' and store them
// in an expanding buffer, so that sending to 'out' never blocks.
//
// Note: the expanding buffer is implemented via "container/ring"
func SendProxyPipeWriter(out chan<- *io.PipeWriter) chan<- *io.PipeWriter {
	proxy := make(chan *io.PipeWriter, PipeWriterCAP)
	go func() {
		n := PipeWriterQUE // the allocated size of the circular queue
		first := ring.New(n)
		last := first
		var c chan<- *io.PipeWriter
		var e *io.PipeWriter
		for {
			c = out
			if first == last {
				// buffer empty: disable output
				c = nil
			} else {
				e = first.Value.(*io.PipeWriter)
			}
			select {
			case e = <-proxy:
				last.Value = e
				if last.Next() == first {
					// buffer full: expand it
					last.Link(ring.New(n))
					n *= 2
				}
				last = last.Next()
			case c <- e:
				first = first.Next()
			}
		}
	}()
	return proxy
}

/* usage as found in $GOROOT/test/chan/sieve2.go
func Sieve() {
	// ...
	primes := make(chan int, 10)
	primes <- 3
	// ...
	go func() {
		// In order to generate the nth prime we only need multiples of primes ≤ sqrt(nth prime).
		// Thus, the merging goroutine will receive from 'primes' much slower than this goroutine will send to it,
		// making the buffer accumulate and block this goroutine from sending, causing a deadlock.
		// The solution is to use a proxy goroutine to do automatic buffering.
		primes := sendproxy(primes)
		// ...

	}()
}
*/
