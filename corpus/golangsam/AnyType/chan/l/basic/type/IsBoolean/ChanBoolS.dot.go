// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsBoolean

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// BoolSChan represents a
// bidirectional
// channel
type BoolSChan interface {
	BoolSROnlyChan // aka "<-chan" - receive only
	BoolSSOnlyChan // aka "chan<-" - send only
}

// BoolSROnlyChan represents a
// receive-only
// channel
type BoolSROnlyChan interface {
	RequestBoolS() (dat []bool)        // the receive function - aka "MyBoolS := <-MyBoolSROnlyChan"
	TryBoolS() (dat []bool, open bool) // the multi-valued comma-ok receive function - aka "MyBoolS, ok := <-MyBoolSROnlyChan"
}

// BoolSSOnlyChan represents a
// send-only
// channel
type BoolSSOnlyChan interface {
	ProvideBoolS(dat []bool) // the send function - aka "MyKind <- some BoolS"
}

// DChBoolS is a demand channel
type DChBoolS struct {
	dat chan []bool
	req chan struct{}
}

// MakeDemandBoolSChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandBoolSChan() *DChBoolS {
	d := new(DChBoolS)
	d.dat = make(chan []bool)
	d.req = make(chan struct{})
	return d
}

// MakeDemandBoolSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandBoolSBuff(cap int) *DChBoolS {
	d := new(DChBoolS)
	d.dat = make(chan []bool, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideBoolS is the send function - aka "MyKind <- some BoolS"
func (c *DChBoolS) ProvideBoolS(dat []bool) {
	<-c.req
	c.dat <- dat
}

// RequestBoolS is the receive function - aka "some BoolS <- MyKind"
func (c *DChBoolS) RequestBoolS() (dat []bool) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryBoolS is the comma-ok multi-valued form of RequestBoolS and
// reports whether a received value was sent before the BoolS channel was closed.
func (c *DChBoolS) TryBoolS() (dat []bool, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
