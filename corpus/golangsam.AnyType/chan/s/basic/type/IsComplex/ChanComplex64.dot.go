// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsComplex

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Complex64Chan represents a
// bidirectional
// channel
type Complex64Chan interface {
	Complex64ROnlyChan // aka "<-chan" - receive only
	Complex64SOnlyChan // aka "chan<-" - send only
}

// Complex64ROnlyChan represents a
// receive-only
// channel
type Complex64ROnlyChan interface {
	RequestComplex64() (dat complex64)        // the receive function - aka "MyComplex64 := <-MyComplex64ROnlyChan"
	TryComplex64() (dat complex64, open bool) // the multi-valued comma-ok receive function - aka "MyComplex64, ok := <-MyComplex64ROnlyChan"
}

// Complex64SOnlyChan represents a
// send-only
// channel
type Complex64SOnlyChan interface {
	ProvideComplex64(dat complex64) // the send function - aka "MyKind <- some Complex64"
}

// SChComplex64 is a supply channel
type SChComplex64 struct {
	dat chan complex64
	// req chan struct{}
}

// MakeSupplyComplex64Chan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyComplex64Chan() *SChComplex64 {
	d := new(SChComplex64)
	d.dat = make(chan complex64)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyComplex64Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyComplex64Buff(cap int) *SChComplex64 {
	d := new(SChComplex64)
	d.dat = make(chan complex64, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideComplex64 is the send function - aka "MyKind <- some Complex64"
func (c *SChComplex64) ProvideComplex64(dat complex64) {
	// .req
	c.dat <- dat
}

// RequestComplex64 is the receive function - aka "some Complex64 <- MyKind"
func (c *SChComplex64) RequestComplex64() (dat complex64) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryComplex64 is the comma-ok multi-valued form of RequestComplex64 and
// reports whether a received value was sent before the Complex64 channel was closed.
func (c *SChComplex64) TryComplex64() (dat complex64, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
