// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsInteger

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Int16Chan represents a
// bidirectional
// channel
type Int16Chan interface {
	Int16ROnlyChan // aka "<-chan" - receive only
	Int16SOnlyChan // aka "chan<-" - send only
}

// Int16ROnlyChan represents a
// receive-only
// channel
type Int16ROnlyChan interface {
	RequestInt16() (dat int16)        // the receive function - aka "MyInt16 := <-MyInt16ROnlyChan"
	TryInt16() (dat int16, open bool) // the multi-valued comma-ok receive function - aka "MyInt16, ok := <-MyInt16ROnlyChan"
}

// Int16SOnlyChan represents a
// send-only
// channel
type Int16SOnlyChan interface {
	ProvideInt16(dat int16) // the send function - aka "MyKind <- some Int16"
}

// DChInt16 is a demand channel
type DChInt16 struct {
	dat chan int16
	req chan struct{}
}

// MakeDemandInt16Chan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandInt16Chan() *DChInt16 {
	d := new(DChInt16)
	d.dat = make(chan int16)
	d.req = make(chan struct{})
	return d
}

// MakeDemandInt16Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandInt16Buff(cap int) *DChInt16 {
	d := new(DChInt16)
	d.dat = make(chan int16, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideInt16 is the send function - aka "MyKind <- some Int16"
func (c *DChInt16) ProvideInt16(dat int16) {
	<-c.req
	c.dat <- dat
}

// RequestInt16 is the receive function - aka "some Int16 <- MyKind"
func (c *DChInt16) RequestInt16() (dat int16) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryInt16 is the comma-ok multi-valued form of RequestInt16 and
// reports whether a received value was sent before the Int16 channel was closed.
func (c *DChInt16) TryInt16() (dat int16, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
