// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ReadSeekerChan represents a
// bidirectional
// channel
type ReadSeekerChan interface {
	ReadSeekerROnlyChan // aka "<-chan" - receive only
	ReadSeekerSOnlyChan // aka "chan<-" - send only
}

// ReadSeekerROnlyChan represents a
// receive-only
// channel
type ReadSeekerROnlyChan interface {
	RequestReadSeeker() (dat io.ReadSeeker)        // the receive function - aka "MyReadSeeker := <-MyReadSeekerROnlyChan"
	TryReadSeeker() (dat io.ReadSeeker, open bool) // the multi-valued comma-ok receive function - aka "MyReadSeeker, ok := <-MyReadSeekerROnlyChan"
}

// ReadSeekerSOnlyChan represents a
// send-only
// channel
type ReadSeekerSOnlyChan interface {
	ProvideReadSeeker(dat io.ReadSeeker) // the send function - aka "MyKind <- some ReadSeeker"
}

// DChReadSeeker is a demand channel
type DChReadSeeker struct {
	dat chan io.ReadSeeker
	req chan struct{}
}

// MakeDemandReadSeekerChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandReadSeekerChan() *DChReadSeeker {
	d := new(DChReadSeeker)
	d.dat = make(chan io.ReadSeeker)
	d.req = make(chan struct{})
	return d
}

// MakeDemandReadSeekerBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandReadSeekerBuff(cap int) *DChReadSeeker {
	d := new(DChReadSeeker)
	d.dat = make(chan io.ReadSeeker, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideReadSeeker is the send function - aka "MyKind <- some ReadSeeker"
func (c *DChReadSeeker) ProvideReadSeeker(dat io.ReadSeeker) {
	<-c.req
	c.dat <- dat
}

// RequestReadSeeker is the receive function - aka "some ReadSeeker <- MyKind"
func (c *DChReadSeeker) RequestReadSeeker() (dat io.ReadSeeker) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryReadSeeker is the comma-ok multi-valued form of RequestReadSeeker and
// reports whether a received value was sent before the ReadSeeker channel was closed.
func (c *DChReadSeeker) TryReadSeeker() (dat io.ReadSeeker, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChReadSeeker is a supply channel
type SChReadSeeker struct {
	dat chan io.ReadSeeker
	// req chan struct{}
}

// MakeSupplyReadSeekerChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyReadSeekerChan() *SChReadSeeker {
	d := new(SChReadSeeker)
	d.dat = make(chan io.ReadSeeker)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyReadSeekerBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyReadSeekerBuff(cap int) *SChReadSeeker {
	d := new(SChReadSeeker)
	d.dat = make(chan io.ReadSeeker, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideReadSeeker is the send function - aka "MyKind <- some ReadSeeker"
func (c *SChReadSeeker) ProvideReadSeeker(dat io.ReadSeeker) {
	// .req
	c.dat <- dat
}

// RequestReadSeeker is the receive function - aka "some ReadSeeker <- MyKind"
func (c *SChReadSeeker) RequestReadSeeker() (dat io.ReadSeeker) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryReadSeeker is the comma-ok multi-valued form of RequestReadSeeker and
// reports whether a received value was sent before the ReadSeeker channel was closed.
func (c *SChReadSeeker) TryReadSeeker() (dat io.ReadSeeker, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
