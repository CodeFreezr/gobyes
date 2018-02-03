// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// WriterToChan represents a
// bidirectional
// channel
type WriterToChan interface {
	WriterToROnlyChan // aka "<-chan" - receive only
	WriterToSOnlyChan // aka "chan<-" - send only
}

// WriterToROnlyChan represents a
// receive-only
// channel
type WriterToROnlyChan interface {
	RequestWriterTo() (dat io.WriterTo)        // the receive function - aka "MyWriterTo := <-MyWriterToROnlyChan"
	TryWriterTo() (dat io.WriterTo, open bool) // the multi-valued comma-ok receive function - aka "MyWriterTo, ok := <-MyWriterToROnlyChan"
}

// WriterToSOnlyChan represents a
// send-only
// channel
type WriterToSOnlyChan interface {
	ProvideWriterTo(dat io.WriterTo) // the send function - aka "MyKind <- some WriterTo"
}

// SChWriterTo is a supply channel
type SChWriterTo struct {
	dat chan io.WriterTo
	// req chan struct{}
}

// MakeSupplyWriterToChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyWriterToChan() *SChWriterTo {
	d := new(SChWriterTo)
	d.dat = make(chan io.WriterTo)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyWriterToBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyWriterToBuff(cap int) *SChWriterTo {
	d := new(SChWriterTo)
	d.dat = make(chan io.WriterTo, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideWriterTo is the send function - aka "MyKind <- some WriterTo"
func (c *SChWriterTo) ProvideWriterTo(dat io.WriterTo) {
	// .req
	c.dat <- dat
}

// RequestWriterTo is the receive function - aka "some WriterTo <- MyKind"
func (c *SChWriterTo) RequestWriterTo() (dat io.WriterTo) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryWriterTo is the comma-ok multi-valued form of RequestWriterTo and
// reports whether a received value was sent before the WriterTo channel was closed.
func (c *SChWriterTo) TryWriterTo() (dat io.WriterTo, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
