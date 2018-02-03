// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsFloat

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

// SChFloat32 is a supply channel
type SChFloat32 struct {
	dat chan float32
	// req chan struct{}
}

// MakeSupplyFloat32Chan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyFloat32Chan() *SChFloat32 {
	d := new(SChFloat32)
	d.dat = make(chan float32)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyFloat32Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyFloat32Buff(cap int) *SChFloat32 {
	d := new(SChFloat32)
	d.dat = make(chan float32, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideFloat32 is the send function - aka "MyKind <- some Float32"
func (c *SChFloat32) ProvideFloat32(dat float32) {
	// .req
	c.dat <- dat
}

// RequestFloat32 is the receive function - aka "some Float32 <- MyKind"
func (c *SChFloat32) RequestFloat32() (dat float32) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryFloat32 is the comma-ok multi-valued form of RequestFloat32 and
// reports whether a received value was sent before the Float32 channel was closed.
func (c *SChFloat32) TryFloat32() (dat float32, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
