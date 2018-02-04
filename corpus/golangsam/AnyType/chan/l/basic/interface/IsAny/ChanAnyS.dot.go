// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsAny

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// AnySChan represents a
// bidirectional
// channel
type AnySChan interface {
	AnySROnlyChan // aka "<-chan" - receive only
	AnySSOnlyChan // aka "chan<-" - send only
}

// AnySROnlyChan represents a
// receive-only
// channel
type AnySROnlyChan interface {
	RequestAnyS() (dat []interface{})        // the receive function - aka "MyAnyS := <-MyAnySROnlyChan"
	TryAnyS() (dat []interface{}, open bool) // the multi-valued comma-ok receive function - aka "MyAnyS, ok := <-MyAnySROnlyChan"
}

// AnySSOnlyChan represents a
// send-only
// channel
type AnySSOnlyChan interface {
	ProvideAnyS(dat []interface{}) // the send function - aka "MyKind <- some AnyS"
}

// DChAnyS is a demand channel
type DChAnyS struct {
	dat chan []interface{}
	req chan struct{}
}

// MakeDemandAnySChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandAnySChan() *DChAnyS {
	d := new(DChAnyS)
	d.dat = make(chan []interface{})
	d.req = make(chan struct{})
	return d
}

// MakeDemandAnySBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandAnySBuff(cap int) *DChAnyS {
	d := new(DChAnyS)
	d.dat = make(chan []interface{}, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideAnyS is the send function - aka "MyKind <- some AnyS"
func (c *DChAnyS) ProvideAnyS(dat []interface{}) {
	<-c.req
	c.dat <- dat
}

// RequestAnyS is the receive function - aka "some AnyS <- MyKind"
func (c *DChAnyS) RequestAnyS() (dat []interface{}) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryAnyS is the comma-ok multi-valued form of RequestAnyS and
// reports whether a received value was sent before the AnyS channel was closed.
func (c *DChAnyS) TryAnyS() (dat []interface{}, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
