// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// PointerSChan represents a
// bidirectional
// channel
type PointerSChan interface {
	PointerSROnlyChan // aka "<-chan" - receive only
	PointerSSOnlyChan // aka "chan<-" - send only
}

// PointerSROnlyChan represents a
// receive-only
// channel
type PointerSROnlyChan interface {
	RequestPointerS() (dat []*SomeType)        // the receive function - aka "MyPointerS := <-MyPointerSROnlyChan"
	TryPointerS() (dat []*SomeType, open bool) // the multi-valued comma-ok receive function - aka "MyPointerS, ok := <-MyPointerSROnlyChan"
}

// PointerSSOnlyChan represents a
// send-only
// channel
type PointerSSOnlyChan interface {
	ProvidePointerS(dat []*SomeType) // the send function - aka "MyKind <- some PointerS"
}

// DChPointerS is a demand channel
type DChPointerS struct {
	dat chan []*SomeType
	req chan struct{}
}

// MakeDemandPointerSChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandPointerSChan() *DChPointerS {
	d := new(DChPointerS)
	d.dat = make(chan []*SomeType)
	d.req = make(chan struct{})
	return d
}

// MakeDemandPointerSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandPointerSBuff(cap int) *DChPointerS {
	d := new(DChPointerS)
	d.dat = make(chan []*SomeType, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvidePointerS is the send function - aka "MyKind <- some PointerS"
func (c *DChPointerS) ProvidePointerS(dat []*SomeType) {
	<-c.req
	c.dat <- dat
}

// RequestPointerS is the receive function - aka "some PointerS <- MyKind"
func (c *DChPointerS) RequestPointerS() (dat []*SomeType) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryPointerS is the comma-ok multi-valued form of RequestPointerS and
// reports whether a received value was sent before the PointerS channel was closed.
func (c *DChPointerS) TryPointerS() (dat []*SomeType, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
