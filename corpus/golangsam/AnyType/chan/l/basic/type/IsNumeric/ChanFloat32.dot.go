// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Float32Chan represents a
// bidirectional
// channel
type Float32Chan interface {
	Float32ROnlyChan // aka "<-chan" - receive only
	Float32SOnlyChan // aka "chan<-" - send only
}

// Float32ROnlyChan represents a
// receive-only
// channel
type Float32ROnlyChan interface {
	RequestFloat32() (dat float32)        // the receive function - aka "MyFloat32 := <-MyFloat32ROnlyChan"
	TryFloat32() (dat float32, open bool) // the multi-valued comma-ok receive function - aka "MyFloat32, ok := <-MyFloat32ROnlyChan"
}

// Float32SOnlyChan represents a
// send-only
// channel
type Float32SOnlyChan interface {
	ProvideFloat32(dat float32) // the send function - aka "MyKind <- some Float32"
}

// DChFloat32 is a demand channel
type DChFloat32 struct {
	dat chan float32
	req chan struct{}
}

// MakeDemandFloat32Chan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFloat32Chan() *DChFloat32 {
	d := new(DChFloat32)
	d.dat = make(chan float32)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFloat32Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFloat32Buff(cap int) *DChFloat32 {
	d := new(DChFloat32)
	d.dat = make(chan float32, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFloat32 is the send function - aka "MyKind <- some Float32"
func (c *DChFloat32) ProvideFloat32(dat float32) {
	<-c.req
	c.dat <- dat
}

// RequestFloat32 is the receive function - aka "some Float32 <- MyKind"
func (c *DChFloat32) RequestFloat32() (dat float32) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFloat32 is the comma-ok multi-valued form of RequestFloat32 and
// reports whether a received value was sent before the Float32 channel was closed.
func (c *DChFloat32) TryFloat32() (dat float32, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
