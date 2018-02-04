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

// DChSeeker is a demand channel
type DChSeeker struct {
	dat chan io.Seeker
	req chan struct{}
}

// MakeDemandSeekerChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandSeekerChan() *DChSeeker {
	d := new(DChSeeker)
	d.dat = make(chan io.Seeker)
	d.req = make(chan struct{})
	return d
}

// MakeDemandSeekerBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandSeekerBuff(cap int) *DChSeeker {
	d := new(DChSeeker)
	d.dat = make(chan io.Seeker, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideSeeker is the send function - aka "MyKind <- some Seeker"
func (c *DChSeeker) ProvideSeeker(dat io.Seeker) {
	<-c.req
	c.dat <- dat
}

// RequestSeeker is the receive function - aka "some Seeker <- MyKind"
func (c *DChSeeker) RequestSeeker() (dat io.Seeker) {
	c.req <- struct{}{}
	return <-c.dat
}

// TrySeeker is the comma-ok multi-valued form of RequestSeeker and
// reports whether a received value was sent before the Seeker channel was closed.
func (c *DChSeeker) TrySeeker() (dat io.Seeker, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
