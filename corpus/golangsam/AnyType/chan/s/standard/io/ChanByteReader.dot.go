// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

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

// SChByteReader is a supply channel
type SChByteReader struct {
	dat chan io.ByteReader
	// req chan struct{}
}

// MakeSupplyByteReaderChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyByteReaderChan() *SChByteReader {
	d := new(SChByteReader)
	d.dat = make(chan io.ByteReader)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyByteReaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyByteReaderBuff(cap int) *SChByteReader {
	d := new(SChByteReader)
	d.dat = make(chan io.ByteReader, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideByteReader is the send function - aka "MyKind <- some ByteReader"
func (c *SChByteReader) ProvideByteReader(dat io.ByteReader) {
	// .req
	c.dat <- dat
}

// RequestByteReader is the receive function - aka "some ByteReader <- MyKind"
func (c *SChByteReader) RequestByteReader() (dat io.ByteReader) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryByteReader is the comma-ok multi-valued form of RequestByteReader and
// reports whether a received value was sent before the ByteReader channel was closed.
func (c *SChByteReader) TryByteReader() (dat io.ByteReader, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
