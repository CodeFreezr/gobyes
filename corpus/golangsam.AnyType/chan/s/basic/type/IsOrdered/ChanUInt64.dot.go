// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

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

// SChUInt64 is a supply channel
type SChUInt64 struct {
	dat chan uint64
	// req chan struct{}
}

// MakeSupplyUInt64Chan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyUInt64Chan() *SChUInt64 {
	d := new(SChUInt64)
	d.dat = make(chan uint64)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyUInt64Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyUInt64Buff(cap int) *SChUInt64 {
	d := new(SChUInt64)
	d.dat = make(chan uint64, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideUInt64 is the send function - aka "MyKind <- some UInt64"
func (c *SChUInt64) ProvideUInt64(dat uint64) {
	// .req
	c.dat <- dat
}

// RequestUInt64 is the receive function - aka "some UInt64 <- MyKind"
func (c *SChUInt64) RequestUInt64() (dat uint64) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryUInt64 is the comma-ok multi-valued form of RequestUInt64 and
// reports whether a received value was sent before the UInt64 channel was closed.
func (c *SChUInt64) TryUInt64() (dat uint64, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
