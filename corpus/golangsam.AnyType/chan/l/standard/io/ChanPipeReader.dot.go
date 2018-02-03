// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// PipeReaderChan represents a
// bidirectional
// channel
type PipeReaderChan interface {
	PipeReaderROnlyChan // aka "<-chan" - receive only
	PipeReaderSOnlyChan // aka "chan<-" - send only
}

// PipeReaderROnlyChan represents a
// receive-only
// channel
type PipeReaderROnlyChan interface {
	RequestPipeReader() (dat *io.PipeReader)        // the receive function - aka "MyPipeReader := <-MyPipeReaderROnlyChan"
	TryPipeReader() (dat *io.PipeReader, open bool) // the multi-valued comma-ok receive function - aka "MyPipeReader, ok := <-MyPipeReaderROnlyChan"
}

// PipeReaderSOnlyChan represents a
// send-only
// channel
type PipeReaderSOnlyChan interface {
	ProvidePipeReader(dat *io.PipeReader) // the send function - aka "MyKind <- some PipeReader"
}

// DChPipeReader is a demand channel
type DChPipeReader struct {
	dat chan *io.PipeReader
	req chan struct{}
}

// MakeDemandPipeReaderChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandPipeReaderChan() *DChPipeReader {
	d := new(DChPipeReader)
	d.dat = make(chan *io.PipeReader)
	d.req = make(chan struct{})
	return d
}

// MakeDemandPipeReaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandPipeReaderBuff(cap int) *DChPipeReader {
	d := new(DChPipeReader)
	d.dat = make(chan *io.PipeReader, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvidePipeReader is the send function - aka "MyKind <- some PipeReader"
func (c *DChPipeReader) ProvidePipeReader(dat *io.PipeReader) {
	<-c.req
	c.dat <- dat
}

// RequestPipeReader is the receive function - aka "some PipeReader <- MyKind"
func (c *DChPipeReader) RequestPipeReader() (dat *io.PipeReader) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryPipeReader is the comma-ok multi-valued form of RequestPipeReader and
// reports whether a received value was sent before the PipeReader channel was closed.
func (c *DChPipeReader) TryPipeReader() (dat *io.PipeReader, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
