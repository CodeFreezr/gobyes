// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsUnsigned

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UInt8Chan represents a
// bidirectional
// channel
type UInt8Chan interface {
	UInt8ROnlyChan // aka "<-chan" - receive only
	UInt8SOnlyChan // aka "chan<-" - send only
}

// UInt8ROnlyChan represents a
// receive-only
// channel
type UInt8ROnlyChan interface {
	RequestUInt8() (dat uint8)        // the receive function - aka "MyUInt8 := <-MyUInt8ROnlyChan"
	TryUInt8() (dat uint8, open bool) // the multi-valued comma-ok receive function - aka "MyUInt8, ok := <-MyUInt8ROnlyChan"
}

// UInt8SOnlyChan represents a
// send-only
// channel
type UInt8SOnlyChan interface {
	ProvideUInt8(dat uint8) // the send function - aka "MyKind <- some UInt8"
}

// DChUInt8 is a demand channel
type DChUInt8 struct {
	dat chan uint8
	req chan struct{}
}

// MakeDemandUInt8Chan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandUInt8Chan() *DChUInt8 {
	d := new(DChUInt8)
	d.dat = make(chan uint8)
	d.req = make(chan struct{})
	return d
}

// MakeDemandUInt8Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandUInt8Buff(cap int) *DChUInt8 {
	d := new(DChUInt8)
	d.dat = make(chan uint8, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideUInt8 is the send function - aka "MyKind <- some UInt8"
func (c *DChUInt8) ProvideUInt8(dat uint8) {
	<-c.req
	c.dat <- dat
}

// RequestUInt8 is the receive function - aka "some UInt8 <- MyKind"
func (c *DChUInt8) RequestUInt8() (dat uint8) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryUInt8 is the comma-ok multi-valued form of RequestUInt8 and
// reports whether a received value was sent before the UInt8 channel was closed.
func (c *DChUInt8) TryUInt8() (dat uint8, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
