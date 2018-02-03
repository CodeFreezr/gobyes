// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Int64Chan represents a
// bidirectional
// channel
type Int64Chan interface {
	Int64ROnlyChan // aka "<-chan" - receive only
	Int64SOnlyChan // aka "chan<-" - send only
}

// Int64ROnlyChan represents a
// receive-only
// channel
type Int64ROnlyChan interface {
	RequestInt64() (dat int64)        // the receive function - aka "MyInt64 := <-MyInt64ROnlyChan"
	TryInt64() (dat int64, open bool) // the multi-valued comma-ok receive function - aka "MyInt64, ok := <-MyInt64ROnlyChan"
}

// Int64SOnlyChan represents a
// send-only
// channel
type Int64SOnlyChan interface {
	ProvideInt64(dat int64) // the send function - aka "MyKind <- some Int64"
}

// DChInt64 is a demand channel
type DChInt64 struct {
	dat chan int64
	req chan struct{}
}

// MakeDemandInt64Chan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandInt64Chan() *DChInt64 {
	d := new(DChInt64)
	d.dat = make(chan int64)
	d.req = make(chan struct{})
	return d
}

// MakeDemandInt64Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandInt64Buff(cap int) *DChInt64 {
	d := new(DChInt64)
	d.dat = make(chan int64, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideInt64 is the send function - aka "MyKind <- some Int64"
func (c *DChInt64) ProvideInt64(dat int64) {
	<-c.req
	c.dat <- dat
}

// RequestInt64 is the receive function - aka "some Int64 <- MyKind"
func (c *DChInt64) RequestInt64() (dat int64) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryInt64 is the comma-ok multi-valued form of RequestInt64 and
// reports whether a received value was sent before the Int64 channel was closed.
func (c *DChInt64) TryInt64() (dat int64, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChInt64 is a supply channel
type SChInt64 struct {
	dat chan int64
	// req chan struct{}
}

// MakeSupplyInt64Chan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyInt64Chan() *SChInt64 {
	d := new(SChInt64)
	d.dat = make(chan int64)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyInt64Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyInt64Buff(cap int) *SChInt64 {
	d := new(SChInt64)
	d.dat = make(chan int64, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideInt64 is the send function - aka "MyKind <- some Int64"
func (c *SChInt64) ProvideInt64(dat int64) {
	// .req
	c.dat <- dat
}

// RequestInt64 is the receive function - aka "some Int64 <- MyKind"
func (c *SChInt64) RequestInt64() (dat int64) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryInt64 is the comma-ok multi-valued form of RequestInt64 and
// reports whether a received value was sent before the Int64 channel was closed.
func (c *SChInt64) TryInt64() (dat int64, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
