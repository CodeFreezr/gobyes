// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bytes"
)

// BufferChan represents a
// bidirectional
// channel
type BufferChan interface {
	BufferROnlyChan // aka "<-chan" - receive only
	BufferSOnlyChan // aka "chan<-" - send only
}

// BufferROnlyChan represents a
// receive-only
// channel
type BufferROnlyChan interface {
	RequestBuffer() (dat bytes.Buffer)        // the receive function - aka "MyBuffer := <-MyBufferROnlyChan"
	TryBuffer() (dat bytes.Buffer, open bool) // the multi-valued comma-ok receive function - aka "MyBuffer, ok := <-MyBufferROnlyChan"
}

// BufferSOnlyChan represents a
// send-only
// channel
type BufferSOnlyChan interface {
	ProvideBuffer(dat bytes.Buffer) // the send function - aka "MyKind <- some Buffer"
}

// DChBuffer is a demand channel
type DChBuffer struct {
	dat chan bytes.Buffer
	req chan struct{}
}

// MakeDemandBufferChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandBufferChan() *DChBuffer {
	d := new(DChBuffer)
	d.dat = make(chan bytes.Buffer)
	d.req = make(chan struct{})
	return d
}

// MakeDemandBufferBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandBufferBuff(cap int) *DChBuffer {
	d := new(DChBuffer)
	d.dat = make(chan bytes.Buffer, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideBuffer is the send function - aka "MyKind <- some Buffer"
func (c *DChBuffer) ProvideBuffer(dat bytes.Buffer) {
	<-c.req
	c.dat <- dat
}

// RequestBuffer is the receive function - aka "some Buffer <- MyKind"
func (c *DChBuffer) RequestBuffer() (dat bytes.Buffer) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryBuffer is the comma-ok multi-valued form of RequestBuffer and
// reports whether a received value was sent before the Buffer channel was closed.
func (c *DChBuffer) TryBuffer() (dat bytes.Buffer, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChBuffer is a supply channel
type SChBuffer struct {
	dat chan bytes.Buffer
	// req chan struct{}
}

// MakeSupplyBufferChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyBufferChan() *SChBuffer {
	d := new(SChBuffer)
	d.dat = make(chan bytes.Buffer)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyBufferBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyBufferBuff(cap int) *SChBuffer {
	d := new(SChBuffer)
	d.dat = make(chan bytes.Buffer, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideBuffer is the send function - aka "MyKind <- some Buffer"
func (c *SChBuffer) ProvideBuffer(dat bytes.Buffer) {
	// .req
	c.dat <- dat
}

// RequestBuffer is the receive function - aka "some Buffer <- MyKind"
func (c *SChBuffer) RequestBuffer() (dat bytes.Buffer) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryBuffer is the comma-ok multi-valued form of RequestBuffer and
// reports whether a received value was sent before the Buffer channel was closed.
func (c *SChBuffer) TryBuffer() (dat bytes.Buffer, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
