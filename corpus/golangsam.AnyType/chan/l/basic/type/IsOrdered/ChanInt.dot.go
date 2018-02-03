// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// IntChan represents a
// bidirectional
// channel
type IntChan interface {
	IntROnlyChan // aka "<-chan" - receive only
	IntSOnlyChan // aka "chan<-" - send only
}

// IntROnlyChan represents a
// receive-only
// channel
type IntROnlyChan interface {
	RequestInt() (dat int)        // the receive function - aka "MyInt := <-MyIntROnlyChan"
	TryInt() (dat int, open bool) // the multi-valued comma-ok receive function - aka "MyInt, ok := <-MyIntROnlyChan"
}

// IntSOnlyChan represents a
// send-only
// channel
type IntSOnlyChan interface {
	ProvideInt(dat int) // the send function - aka "MyKind <- some Int"
}

// DChInt is a demand channel
type DChInt struct {
	dat chan int
	req chan struct{}
}

// MakeDemandIntChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandIntChan() *DChInt {
	d := new(DChInt)
	d.dat = make(chan int)
	d.req = make(chan struct{})
	return d
}

// MakeDemandIntBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandIntBuff(cap int) *DChInt {
	d := new(DChInt)
	d.dat = make(chan int, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideInt is the send function - aka "MyKind <- some Int"
func (c *DChInt) ProvideInt(dat int) {
	<-c.req
	c.dat <- dat
}

// RequestInt is the receive function - aka "some Int <- MyKind"
func (c *DChInt) RequestInt() (dat int) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryInt is the comma-ok multi-valued form of RequestInt and
// reports whether a received value was sent before the Int channel was closed.
func (c *DChInt) TryInt() (dat int, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
