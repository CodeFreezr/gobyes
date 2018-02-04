// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ReaderAtChan represents a
// bidirectional
// channel
type ReaderAtChan interface {
	ReaderAtROnlyChan // aka "<-chan" - receive only
	ReaderAtSOnlyChan // aka "chan<-" - send only
}

// ReaderAtROnlyChan represents a
// receive-only
// channel
type ReaderAtROnlyChan interface {
	RequestReaderAt() (dat io.ReaderAt)        // the receive function - aka "MyReaderAt := <-MyReaderAtROnlyChan"
	TryReaderAt() (dat io.ReaderAt, open bool) // the multi-valued comma-ok receive function - aka "MyReaderAt, ok := <-MyReaderAtROnlyChan"
}

// ReaderAtSOnlyChan represents a
// send-only
// channel
type ReaderAtSOnlyChan interface {
	ProvideReaderAt(dat io.ReaderAt) // the send function - aka "MyKind <- some ReaderAt"
}

// SChReaderAt is a supply channel
type SChReaderAt struct {
	dat chan io.ReaderAt
	// req chan struct{}
}

// MakeSupplyReaderAtChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyReaderAtChan() *SChReaderAt {
	d := new(SChReaderAt)
	d.dat = make(chan io.ReaderAt)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyReaderAtBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyReaderAtBuff(cap int) *SChReaderAt {
	d := new(SChReaderAt)
	d.dat = make(chan io.ReaderAt, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideReaderAt is the send function - aka "MyKind <- some ReaderAt"
func (c *SChReaderAt) ProvideReaderAt(dat io.ReaderAt) {
	// .req
	c.dat <- dat
}

// RequestReaderAt is the receive function - aka "some ReaderAt <- MyKind"
func (c *SChReaderAt) RequestReaderAt() (dat io.ReaderAt) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryReaderAt is the comma-ok multi-valued form of RequestReaderAt and
// reports whether a received value was sent before the ReaderAt channel was closed.
func (c *SChReaderAt) TryReaderAt() (dat io.ReaderAt, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
