// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// WriterAtChan represents a
// bidirectional
// channel
type WriterAtChan interface {
	WriterAtROnlyChan // aka "<-chan" - receive only
	WriterAtSOnlyChan // aka "chan<-" - send only
}

// WriterAtROnlyChan represents a
// receive-only
// channel
type WriterAtROnlyChan interface {
	RequestWriterAt() (dat io.WriterAt)        // the receive function - aka "MyWriterAt := <-MyWriterAtROnlyChan"
	TryWriterAt() (dat io.WriterAt, open bool) // the multi-valued comma-ok receive function - aka "MyWriterAt, ok := <-MyWriterAtROnlyChan"
}

// WriterAtSOnlyChan represents a
// send-only
// channel
type WriterAtSOnlyChan interface {
	ProvideWriterAt(dat io.WriterAt) // the send function - aka "MyKind <- some WriterAt"
}

// DChWriterAt is a demand channel
type DChWriterAt struct {
	dat chan io.WriterAt
	req chan struct{}
}

// MakeDemandWriterAtChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandWriterAtChan() *DChWriterAt {
	d := new(DChWriterAt)
	d.dat = make(chan io.WriterAt)
	d.req = make(chan struct{})
	return d
}

// MakeDemandWriterAtBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandWriterAtBuff(cap int) *DChWriterAt {
	d := new(DChWriterAt)
	d.dat = make(chan io.WriterAt, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideWriterAt is the send function - aka "MyKind <- some WriterAt"
func (c *DChWriterAt) ProvideWriterAt(dat io.WriterAt) {
	<-c.req
	c.dat <- dat
}

// RequestWriterAt is the receive function - aka "some WriterAt <- MyKind"
func (c *DChWriterAt) RequestWriterAt() (dat io.WriterAt) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryWriterAt is the comma-ok multi-valued form of RequestWriterAt and
// reports whether a received value was sent before the WriterAt channel was closed.
func (c *DChWriterAt) TryWriterAt() (dat io.WriterAt, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChWriterAt is a supply channel
type SChWriterAt struct {
	dat chan io.WriterAt
	// req chan struct{}
}

// MakeSupplyWriterAtChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyWriterAtChan() *SChWriterAt {
	d := new(SChWriterAt)
	d.dat = make(chan io.WriterAt)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyWriterAtBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyWriterAtBuff(cap int) *SChWriterAt {
	d := new(SChWriterAt)
	d.dat = make(chan io.WriterAt, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideWriterAt is the send function - aka "MyKind <- some WriterAt"
func (c *SChWriterAt) ProvideWriterAt(dat io.WriterAt) {
	// .req
	c.dat <- dat
}

// RequestWriterAt is the receive function - aka "some WriterAt <- MyKind"
func (c *SChWriterAt) RequestWriterAt() (dat io.WriterAt) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryWriterAt is the comma-ok multi-valued form of RequestWriterAt and
// reports whether a received value was sent before the WriterAt channel was closed.
func (c *SChWriterAt) TryWriterAt() (dat io.WriterAt, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
