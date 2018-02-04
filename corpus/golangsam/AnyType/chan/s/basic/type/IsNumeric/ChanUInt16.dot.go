// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UInt16Chan represents a
// bidirectional
// channel
type UInt16Chan interface {
	UInt16ROnlyChan // aka "<-chan" - receive only
	UInt16SOnlyChan // aka "chan<-" - send only
}

// UInt16ROnlyChan represents a
// receive-only
// channel
type UInt16ROnlyChan interface {
	RequestUInt16() (dat uint16)        // the receive function - aka "MyUInt16 := <-MyUInt16ROnlyChan"
	TryUInt16() (dat uint16, open bool) // the multi-valued comma-ok receive function - aka "MyUInt16, ok := <-MyUInt16ROnlyChan"
}

// UInt16SOnlyChan represents a
// send-only
// channel
type UInt16SOnlyChan interface {
	ProvideUInt16(dat uint16) // the send function - aka "MyKind <- some UInt16"
}

// SChUInt16 is a supply channel
type SChUInt16 struct {
	dat chan uint16
	// req chan struct{}
}

// MakeSupplyUInt16Chan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyUInt16Chan() *SChUInt16 {
	d := new(SChUInt16)
	d.dat = make(chan uint16)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyUInt16Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyUInt16Buff(cap int) *SChUInt16 {
	d := new(SChUInt16)
	d.dat = make(chan uint16, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideUInt16 is the send function - aka "MyKind <- some UInt16"
func (c *SChUInt16) ProvideUInt16(dat uint16) {
	// .req
	c.dat <- dat
}

// RequestUInt16 is the receive function - aka "some UInt16 <- MyKind"
func (c *SChUInt16) RequestUInt16() (dat uint16) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryUInt16 is the comma-ok multi-valued form of RequestUInt16 and
// reports whether a received value was sent before the UInt16 channel was closed.
func (c *SChUInt16) TryUInt16() (dat uint16, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
