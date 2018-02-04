// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsRune

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// RuneReaderChan represents a
// bidirectional
// channel
type RuneReaderChan interface {
	RuneReaderROnlyChan // aka "<-chan" - receive only
	RuneReaderSOnlyChan // aka "chan<-" - send only
}

// RuneReaderROnlyChan represents a
// receive-only
// channel
type RuneReaderROnlyChan interface {
	RequestRuneReader() (dat io.RuneReader)        // the receive function - aka "MyRuneReader := <-MyRuneReaderROnlyChan"
	TryRuneReader() (dat io.RuneReader, open bool) // the multi-valued comma-ok receive function - aka "MyRuneReader, ok := <-MyRuneReaderROnlyChan"
}

// RuneReaderSOnlyChan represents a
// send-only
// channel
type RuneReaderSOnlyChan interface {
	ProvideRuneReader(dat io.RuneReader) // the send function - aka "MyKind <- some RuneReader"
}

// SChRuneReader is a supply channel
type SChRuneReader struct {
	dat chan io.RuneReader
	// req chan struct{}
}

// MakeSupplyRuneReaderChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyRuneReaderChan() *SChRuneReader {
	d := new(SChRuneReader)
	d.dat = make(chan io.RuneReader)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyRuneReaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyRuneReaderBuff(cap int) *SChRuneReader {
	d := new(SChRuneReader)
	d.dat = make(chan io.RuneReader, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideRuneReader is the send function - aka "MyKind <- some RuneReader"
func (c *SChRuneReader) ProvideRuneReader(dat io.RuneReader) {
	// .req
	c.dat <- dat
}

// RequestRuneReader is the receive function - aka "some RuneReader <- MyKind"
func (c *SChRuneReader) RequestRuneReader() (dat io.RuneReader) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryRuneReader is the comma-ok multi-valued form of RequestRuneReader and
// reports whether a received value was sent before the RuneReader channel was closed.
func (c *SChRuneReader) TryRuneReader() (dat io.RuneReader, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
