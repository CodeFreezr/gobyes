// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsUnsigned

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UInt64Chan represents a
// bidirectional
// channel
type UInt64Chan interface {
	UInt64ROnlyChan // aka "<-chan" - receive only
	UInt64SOnlyChan // aka "chan<-" - send only
}

// UInt64ROnlyChan represents a
// receive-only
// channel
type UInt64ROnlyChan interface {
	RequestUInt64() (dat uint64)        // the receive function - aka "MyUInt64 := <-MyUInt64ROnlyChan"
	TryUInt64() (dat uint64, open bool) // the multi-valued comma-ok receive function - aka "MyUInt64, ok := <-MyUInt64ROnlyChan"
}

// UInt64SOnlyChan represents a
// send-only
// channel
type UInt64SOnlyChan interface {
	ProvideUInt64(dat uint64) // the send function - aka "MyKind <- some UInt64"
}

// DChUInt64 is a demand channel
type DChUInt64 struct {
	dat chan uint64
	req chan struct{}
}

// MakeDemandUInt64Chan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandUInt64Chan() *DChUInt64 {
	d := new(DChUInt64)
	d.dat = make(chan uint64)
	d.req = make(chan struct{})
	return d
}

// MakeDemandUInt64Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandUInt64Buff(cap int) *DChUInt64 {
	d := new(DChUInt64)
	d.dat = make(chan uint64, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideUInt64 is the send function - aka "MyKind <- some UInt64"
func (c *DChUInt64) ProvideUInt64(dat uint64) {
	<-c.req
	c.dat <- dat
}

// RequestUInt64 is the receive function - aka "some UInt64 <- MyKind"
func (c *DChUInt64) RequestUInt64() (dat uint64) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryUInt64 is the comma-ok multi-valued form of RequestUInt64 and
// reports whether a received value was sent before the UInt64 channel was closed.
func (c *DChUInt64) TryUInt64() (dat uint64, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
