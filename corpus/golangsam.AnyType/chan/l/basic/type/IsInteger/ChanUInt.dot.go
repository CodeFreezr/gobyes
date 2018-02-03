// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsInteger

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UIntChan represents a
// bidirectional
// channel
type UIntChan interface {
	UIntROnlyChan // aka "<-chan" - receive only
	UIntSOnlyChan // aka "chan<-" - send only
}

// UIntROnlyChan represents a
// receive-only
// channel
type UIntROnlyChan interface {
	RequestUInt() (dat uint)        // the receive function - aka "MyUInt := <-MyUIntROnlyChan"
	TryUInt() (dat uint, open bool) // the multi-valued comma-ok receive function - aka "MyUInt, ok := <-MyUIntROnlyChan"
}

// UIntSOnlyChan represents a
// send-only
// channel
type UIntSOnlyChan interface {
	ProvideUInt(dat uint) // the send function - aka "MyKind <- some UInt"
}

// DChUInt is a demand channel
type DChUInt struct {
	dat chan uint
	req chan struct{}
}

// MakeDemandUIntChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandUIntChan() *DChUInt {
	d := new(DChUInt)
	d.dat = make(chan uint)
	d.req = make(chan struct{})
	return d
}

// MakeDemandUIntBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandUIntBuff(cap int) *DChUInt {
	d := new(DChUInt)
	d.dat = make(chan uint, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideUInt is the send function - aka "MyKind <- some UInt"
func (c *DChUInt) ProvideUInt(dat uint) {
	<-c.req
	c.dat <- dat
}

// RequestUInt is the receive function - aka "some UInt <- MyKind"
func (c *DChUInt) RequestUInt() (dat uint) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryUInt is the comma-ok multi-valued form of RequestUInt and
// reports whether a received value was sent before the UInt channel was closed.
func (c *DChUInt) TryUInt() (dat uint, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
