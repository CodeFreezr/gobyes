// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ReaderFromChan represents a
// bidirectional
// channel
type ReaderFromChan interface {
	ReaderFromROnlyChan // aka "<-chan" - receive only
	ReaderFromSOnlyChan // aka "chan<-" - send only
}

// ReaderFromROnlyChan represents a
// receive-only
// channel
type ReaderFromROnlyChan interface {
	RequestReaderFrom() (dat io.ReaderFrom)        // the receive function - aka "MyReaderFrom := <-MyReaderFromROnlyChan"
	TryReaderFrom() (dat io.ReaderFrom, open bool) // the multi-valued comma-ok receive function - aka "MyReaderFrom, ok := <-MyReaderFromROnlyChan"
}

// ReaderFromSOnlyChan represents a
// send-only
// channel
type ReaderFromSOnlyChan interface {
	ProvideReaderFrom(dat io.ReaderFrom) // the send function - aka "MyKind <- some ReaderFrom"
}

// SChReaderFrom is a supply channel
type SChReaderFrom struct {
	dat chan io.ReaderFrom
	// req chan struct{}
}

// MakeSupplyReaderFromChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyReaderFromChan() *SChReaderFrom {
	d := new(SChReaderFrom)
	d.dat = make(chan io.ReaderFrom)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyReaderFromBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyReaderFromBuff(cap int) *SChReaderFrom {
	d := new(SChReaderFrom)
	d.dat = make(chan io.ReaderFrom, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideReaderFrom is the send function - aka "MyKind <- some ReaderFrom"
func (c *SChReaderFrom) ProvideReaderFrom(dat io.ReaderFrom) {
	// .req
	c.dat <- dat
}

// RequestReaderFrom is the receive function - aka "some ReaderFrom <- MyKind"
func (c *SChReaderFrom) RequestReaderFrom() (dat io.ReaderFrom) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryReaderFrom is the comma-ok multi-valued form of RequestReaderFrom and
// reports whether a received value was sent before the ReaderFrom channel was closed.
func (c *SChReaderFrom) TryReaderFrom() (dat io.ReaderFrom, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
