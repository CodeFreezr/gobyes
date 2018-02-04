// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// SeekerChan represents a
// bidirectional
// channel
type SeekerChan interface {
	SeekerROnlyChan // aka "<-chan" - receive only
	SeekerSOnlyChan // aka "chan<-" - send only
}

// SeekerROnlyChan represents a
// receive-only
// channel
type SeekerROnlyChan interface {
	RequestSeeker() (dat io.Seeker)        // the receive function - aka "MySeeker := <-MySeekerROnlyChan"
	TrySeeker() (dat io.Seeker, open bool) // the multi-valued comma-ok receive function - aka "MySeeker, ok := <-MySeekerROnlyChan"
}

// SeekerSOnlyChan represents a
// send-only
// channel
type SeekerSOnlyChan interface {
	ProvideSeeker(dat io.Seeker) // the send function - aka "MyKind <- some Seeker"
}

// SChSeeker is a supply channel
type SChSeeker struct {
	dat chan io.Seeker
	// req chan struct{}
}

// MakeSupplySeekerChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplySeekerChan() *SChSeeker {
	d := new(SChSeeker)
	d.dat = make(chan io.Seeker)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplySeekerBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplySeekerBuff(cap int) *SChSeeker {
	d := new(SChSeeker)
	d.dat = make(chan io.Seeker, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideSeeker is the send function - aka "MyKind <- some Seeker"
func (c *SChSeeker) ProvideSeeker(dat io.Seeker) {
	// .req
	c.dat <- dat
}

// RequestSeeker is the receive function - aka "some Seeker <- MyKind"
func (c *SChSeeker) RequestSeeker() (dat io.Seeker) {
	// eq <- struct{}{}
	return <-c.dat
}

// TrySeeker is the comma-ok multi-valued form of RequestSeeker and
// reports whether a received value was sent before the Seeker channel was closed.
func (c *SChSeeker) TrySeeker() (dat io.Seeker, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
