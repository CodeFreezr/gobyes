// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UInt32Chan represents a
// bidirectional
// channel
type UInt32Chan interface {
	UInt32ROnlyChan // aka "<-chan" - receive only
	UInt32SOnlyChan // aka "chan<-" - send only
}

// UInt32ROnlyChan represents a
// receive-only
// channel
type UInt32ROnlyChan interface {
	RequestUInt32() (dat uint32)        // the receive function - aka "MyUInt32 := <-MyUInt32ROnlyChan"
	TryUInt32() (dat uint32, open bool) // the multi-valued comma-ok receive function - aka "MyUInt32, ok := <-MyUInt32ROnlyChan"
}

// UInt32SOnlyChan represents a
// send-only
// channel
type UInt32SOnlyChan interface {
	ProvideUInt32(dat uint32) // the send function - aka "MyKind <- some UInt32"
}

// DChUInt32 is a demand channel
type DChUInt32 struct {
	dat chan uint32
	req chan struct{}
}

// MakeDemandUInt32Chan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandUInt32Chan() *DChUInt32 {
	d := new(DChUInt32)
	d.dat = make(chan uint32)
	d.req = make(chan struct{})
	return d
}

// MakeDemandUInt32Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandUInt32Buff(cap int) *DChUInt32 {
	d := new(DChUInt32)
	d.dat = make(chan uint32, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideUInt32 is the send function - aka "MyKind <- some UInt32"
func (c *DChUInt32) ProvideUInt32(dat uint32) {
	<-c.req
	c.dat <- dat
}

// RequestUInt32 is the receive function - aka "some UInt32 <- MyKind"
func (c *DChUInt32) RequestUInt32() (dat uint32) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryUInt32 is the comma-ok multi-valued form of RequestUInt32 and
// reports whether a received value was sent before the UInt32 channel was closed.
func (c *DChUInt32) TryUInt32() (dat uint32, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChUInt32 is a supply channel
type SChUInt32 struct {
	dat chan uint32
	// req chan struct{}
}

// MakeSupplyUInt32Chan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyUInt32Chan() *SChUInt32 {
	d := new(SChUInt32)
	d.dat = make(chan uint32)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyUInt32Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyUInt32Buff(cap int) *SChUInt32 {
	d := new(SChUInt32)
	d.dat = make(chan uint32, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideUInt32 is the send function - aka "MyKind <- some UInt32"
func (c *SChUInt32) ProvideUInt32(dat uint32) {
	// .req
	c.dat <- dat
}

// RequestUInt32 is the receive function - aka "some UInt32 <- MyKind"
func (c *SChUInt32) RequestUInt32() (dat uint32) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryUInt32 is the comma-ok multi-valued form of RequestUInt32 and
// reports whether a received value was sent before the UInt32 channel was closed.
func (c *SChUInt32) TryUInt32() (dat uint32, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
