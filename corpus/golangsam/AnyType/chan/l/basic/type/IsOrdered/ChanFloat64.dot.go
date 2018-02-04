// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Float64Chan represents a
// bidirectional
// channel
type Float64Chan interface {
	Float64ROnlyChan // aka "<-chan" - receive only
	Float64SOnlyChan // aka "chan<-" - send only
}

// Float64ROnlyChan represents a
// receive-only
// channel
type Float64ROnlyChan interface {
	RequestFloat64() (dat float64)        // the receive function - aka "MyFloat64 := <-MyFloat64ROnlyChan"
	TryFloat64() (dat float64, open bool) // the multi-valued comma-ok receive function - aka "MyFloat64, ok := <-MyFloat64ROnlyChan"
}

// Float64SOnlyChan represents a
// send-only
// channel
type Float64SOnlyChan interface {
	ProvideFloat64(dat float64) // the send function - aka "MyKind <- some Float64"
}

// DChFloat64 is a demand channel
type DChFloat64 struct {
	dat chan float64
	req chan struct{}
}

// MakeDemandFloat64Chan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFloat64Chan() *DChFloat64 {
	d := new(DChFloat64)
	d.dat = make(chan float64)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFloat64Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFloat64Buff(cap int) *DChFloat64 {
	d := new(DChFloat64)
	d.dat = make(chan float64, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFloat64 is the send function - aka "MyKind <- some Float64"
func (c *DChFloat64) ProvideFloat64(dat float64) {
	<-c.req
	c.dat <- dat
}

// RequestFloat64 is the receive function - aka "some Float64 <- MyKind"
func (c *DChFloat64) RequestFloat64() (dat float64) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFloat64 is the comma-ok multi-valued form of RequestFloat64 and
// reports whether a received value was sent before the Float64 channel was closed.
func (c *DChFloat64) TryFloat64() (dat float64, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
