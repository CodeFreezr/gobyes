// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Int32Chan represents a
// bidirectional
// channel
type Int32Chan interface {
	Int32ROnlyChan // aka "<-chan" - receive only
	Int32SOnlyChan // aka "chan<-" - send only
}

// Int32ROnlyChan represents a
// receive-only
// channel
type Int32ROnlyChan interface {
	RequestInt32() (dat int32)        // the receive function - aka "MyInt32 := <-MyInt32ROnlyChan"
	TryInt32() (dat int32, open bool) // the multi-valued comma-ok receive function - aka "MyInt32, ok := <-MyInt32ROnlyChan"
}

// Int32SOnlyChan represents a
// send-only
// channel
type Int32SOnlyChan interface {
	ProvideInt32(dat int32) // the send function - aka "MyKind <- some Int32"
}

// DChInt32 is a demand channel
type DChInt32 struct {
	dat chan int32
	req chan struct{}
}

// MakeDemandInt32Chan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandInt32Chan() *DChInt32 {
	d := new(DChInt32)
	d.dat = make(chan int32)
	d.req = make(chan struct{})
	return d
}

// MakeDemandInt32Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandInt32Buff(cap int) *DChInt32 {
	d := new(DChInt32)
	d.dat = make(chan int32, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideInt32 is the send function - aka "MyKind <- some Int32"
func (c *DChInt32) ProvideInt32(dat int32) {
	<-c.req
	c.dat <- dat
}

// RequestInt32 is the receive function - aka "some Int32 <- MyKind"
func (c *DChInt32) RequestInt32() (dat int32) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryInt32 is the comma-ok multi-valued form of RequestInt32 and
// reports whether a received value was sent before the Int32 channel was closed.
func (c *DChInt32) TryInt32() (dat int32, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
