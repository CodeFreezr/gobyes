// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Int8Chan represents a
// bidirectional
// channel
type Int8Chan interface {
	Int8ROnlyChan // aka "<-chan" - receive only
	Int8SOnlyChan // aka "chan<-" - send only
}

// Int8ROnlyChan represents a
// receive-only
// channel
type Int8ROnlyChan interface {
	RequestInt8() (dat int8)        // the receive function - aka "MyInt8 := <-MyInt8ROnlyChan"
	TryInt8() (dat int8, open bool) // the multi-valued comma-ok receive function - aka "MyInt8, ok := <-MyInt8ROnlyChan"
}

// Int8SOnlyChan represents a
// send-only
// channel
type Int8SOnlyChan interface {
	ProvideInt8(dat int8) // the send function - aka "MyKind <- some Int8"
}

// DChInt8 is a demand channel
type DChInt8 struct {
	dat chan int8
	req chan struct{}
}

// MakeDemandInt8Chan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandInt8Chan() *DChInt8 {
	d := new(DChInt8)
	d.dat = make(chan int8)
	d.req = make(chan struct{})
	return d
}

// MakeDemandInt8Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandInt8Buff(cap int) *DChInt8 {
	d := new(DChInt8)
	d.dat = make(chan int8, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideInt8 is the send function - aka "MyKind <- some Int8"
func (c *DChInt8) ProvideInt8(dat int8) {
	<-c.req
	c.dat <- dat
}

// RequestInt8 is the receive function - aka "some Int8 <- MyKind"
func (c *DChInt8) RequestInt8() (dat int8) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryInt8 is the comma-ok multi-valued form of RequestInt8 and
// reports whether a received value was sent before the Int8 channel was closed.
func (c *DChInt8) TryInt8() (dat int8, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChInt8 is a supply channel
type SChInt8 struct {
	dat chan int8
	// req chan struct{}
}

// MakeSupplyInt8Chan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyInt8Chan() *SChInt8 {
	d := new(SChInt8)
	d.dat = make(chan int8)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyInt8Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyInt8Buff(cap int) *SChInt8 {
	d := new(SChInt8)
	d.dat = make(chan int8, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideInt8 is the send function - aka "MyKind <- some Int8"
func (c *SChInt8) ProvideInt8(dat int8) {
	// .req
	c.dat <- dat
}

// RequestInt8 is the receive function - aka "some Int8 <- MyKind"
func (c *SChInt8) RequestInt8() (dat int8) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryInt8 is the comma-ok multi-valued form of RequestInt8 and
// reports whether a received value was sent before the Int8 channel was closed.
func (c *SChInt8) TryInt8() (dat int8, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
