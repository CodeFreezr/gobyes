// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// PipeWriterChan represents a
// bidirectional
// channel
type PipeWriterChan interface {
	PipeWriterROnlyChan // aka "<-chan" - receive only
	PipeWriterSOnlyChan // aka "chan<-" - send only
}

// PipeWriterROnlyChan represents a
// receive-only
// channel
type PipeWriterROnlyChan interface {
	RequestPipeWriter() (dat *io.PipeWriter)        // the receive function - aka "MyPipeWriter := <-MyPipeWriterROnlyChan"
	TryPipeWriter() (dat *io.PipeWriter, open bool) // the multi-valued comma-ok receive function - aka "MyPipeWriter, ok := <-MyPipeWriterROnlyChan"
}

// PipeWriterSOnlyChan represents a
// send-only
// channel
type PipeWriterSOnlyChan interface {
	ProvidePipeWriter(dat *io.PipeWriter) // the send function - aka "MyKind <- some PipeWriter"
}

// SChPipeWriter is a supply channel
type SChPipeWriter struct {
	dat chan *io.PipeWriter
	// req chan struct{}
}

// MakeSupplyPipeWriterChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyPipeWriterChan() *SChPipeWriter {
	d := new(SChPipeWriter)
	d.dat = make(chan *io.PipeWriter)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyPipeWriterBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyPipeWriterBuff(cap int) *SChPipeWriter {
	d := new(SChPipeWriter)
	d.dat = make(chan *io.PipeWriter, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvidePipeWriter is the send function - aka "MyKind <- some PipeWriter"
func (c *SChPipeWriter) ProvidePipeWriter(dat *io.PipeWriter) {
	// .req
	c.dat <- dat
}

// RequestPipeWriter is the receive function - aka "some PipeWriter <- MyKind"
func (c *SChPipeWriter) RequestPipeWriter() (dat *io.PipeWriter) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryPipeWriter is the comma-ok multi-valued form of RequestPipeWriter and
// reports whether a received value was sent before the PipeWriter channel was closed.
func (c *SChPipeWriter) TryPipeWriter() (dat *io.PipeWriter, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
