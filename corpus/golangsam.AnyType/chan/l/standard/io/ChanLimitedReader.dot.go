// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// LimitedReaderChan represents a
// bidirectional
// channel
type LimitedReaderChan interface {
	LimitedReaderROnlyChan // aka "<-chan" - receive only
	LimitedReaderSOnlyChan // aka "chan<-" - send only
}

// LimitedReaderROnlyChan represents a
// receive-only
// channel
type LimitedReaderROnlyChan interface {
	RequestLimitedReader() (dat *io.LimitedReader)        // the receive function - aka "MyLimitedReader := <-MyLimitedReaderROnlyChan"
	TryLimitedReader() (dat *io.LimitedReader, open bool) // the multi-valued comma-ok receive function - aka "MyLimitedReader, ok := <-MyLimitedReaderROnlyChan"
}

// LimitedReaderSOnlyChan represents a
// send-only
// channel
type LimitedReaderSOnlyChan interface {
	ProvideLimitedReader(dat *io.LimitedReader) // the send function - aka "MyKind <- some LimitedReader"
}

// DChLimitedReader is a demand channel
type DChLimitedReader struct {
	dat chan *io.LimitedReader
	req chan struct{}
}

// MakeDemandLimitedReaderChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandLimitedReaderChan() *DChLimitedReader {
	d := new(DChLimitedReader)
	d.dat = make(chan *io.LimitedReader)
	d.req = make(chan struct{})
	return d
}

// MakeDemandLimitedReaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandLimitedReaderBuff(cap int) *DChLimitedReader {
	d := new(DChLimitedReader)
	d.dat = make(chan *io.LimitedReader, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideLimitedReader is the send function - aka "MyKind <- some LimitedReader"
func (c *DChLimitedReader) ProvideLimitedReader(dat *io.LimitedReader) {
	<-c.req
	c.dat <- dat
}

// RequestLimitedReader is the receive function - aka "some LimitedReader <- MyKind"
func (c *DChLimitedReader) RequestLimitedReader() (dat *io.LimitedReader) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryLimitedReader is the comma-ok multi-valued form of RequestLimitedReader and
// reports whether a received value was sent before the LimitedReader channel was closed.
func (c *DChLimitedReader) TryLimitedReader() (dat *io.LimitedReader, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
