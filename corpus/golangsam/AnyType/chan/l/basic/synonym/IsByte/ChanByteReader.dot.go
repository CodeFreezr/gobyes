// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ByteReaderChan represents a
// bidirectional
// channel
type ByteReaderChan interface {
	ByteReaderROnlyChan // aka "<-chan" - receive only
	ByteReaderSOnlyChan // aka "chan<-" - send only
}

// ByteReaderROnlyChan represents a
// receive-only
// channel
type ByteReaderROnlyChan interface {
	RequestByteReader() (dat io.ByteReader)        // the receive function - aka "MyByteReader := <-MyByteReaderROnlyChan"
	TryByteReader() (dat io.ByteReader, open bool) // the multi-valued comma-ok receive function - aka "MyByteReader, ok := <-MyByteReaderROnlyChan"
}

// ByteReaderSOnlyChan represents a
// send-only
// channel
type ByteReaderSOnlyChan interface {
	ProvideByteReader(dat io.ByteReader) // the send function - aka "MyKind <- some ByteReader"
}

// DChByteReader is a demand channel
type DChByteReader struct {
	dat chan io.ByteReader
	req chan struct{}
}

// MakeDemandByteReaderChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandByteReaderChan() *DChByteReader {
	d := new(DChByteReader)
	d.dat = make(chan io.ByteReader)
	d.req = make(chan struct{})
	return d
}

// MakeDemandByteReaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandByteReaderBuff(cap int) *DChByteReader {
	d := new(DChByteReader)
	d.dat = make(chan io.ByteReader, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideByteReader is the send function - aka "MyKind <- some ByteReader"
func (c *DChByteReader) ProvideByteReader(dat io.ByteReader) {
	<-c.req
	c.dat <- dat
}

// RequestByteReader is the receive function - aka "some ByteReader <- MyKind"
func (c *DChByteReader) RequestByteReader() (dat io.ByteReader) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryByteReader is the comma-ok multi-valued form of RequestByteReader and
// reports whether a received value was sent before the ByteReader channel was closed.
func (c *DChByteReader) TryByteReader() (dat io.ByteReader, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
