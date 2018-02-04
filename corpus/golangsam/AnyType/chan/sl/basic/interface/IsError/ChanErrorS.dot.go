// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsError

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// ErrorSChan represents a
// bidirectional
// channel
type ErrorSChan interface {
	ErrorSROnlyChan // aka "<-chan" - receive only
	ErrorSSOnlyChan // aka "chan<-" - send only
}

// ErrorSROnlyChan represents a
// receive-only
// channel
type ErrorSROnlyChan interface {
	RequestErrorS() (dat []error)        // the receive function - aka "MyErrorS := <-MyErrorSROnlyChan"
	TryErrorS() (dat []error, open bool) // the multi-valued comma-ok receive function - aka "MyErrorS, ok := <-MyErrorSROnlyChan"
}

// ErrorSSOnlyChan represents a
// send-only
// channel
type ErrorSSOnlyChan interface {
	ProvideErrorS(dat []error) // the send function - aka "MyKind <- some ErrorS"
}

// DChErrorS is a demand channel
type DChErrorS struct {
	dat chan []error
	req chan struct{}
}

// MakeDemandErrorSChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandErrorSChan() *DChErrorS {
	d := new(DChErrorS)
	d.dat = make(chan []error)
	d.req = make(chan struct{})
	return d
}

// MakeDemandErrorSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandErrorSBuff(cap int) *DChErrorS {
	d := new(DChErrorS)
	d.dat = make(chan []error, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideErrorS is the send function - aka "MyKind <- some ErrorS"
func (c *DChErrorS) ProvideErrorS(dat []error) {
	<-c.req
	c.dat <- dat
}

// RequestErrorS is the receive function - aka "some ErrorS <- MyKind"
func (c *DChErrorS) RequestErrorS() (dat []error) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryErrorS is the comma-ok multi-valued form of RequestErrorS and
// reports whether a received value was sent before the ErrorS channel was closed.
func (c *DChErrorS) TryErrorS() (dat []error, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChErrorS is a supply channel
type SChErrorS struct {
	dat chan []error
	// req chan struct{}
}

// MakeSupplyErrorSChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyErrorSChan() *SChErrorS {
	d := new(SChErrorS)
	d.dat = make(chan []error)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyErrorSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyErrorSBuff(cap int) *SChErrorS {
	d := new(SChErrorS)
	d.dat = make(chan []error, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideErrorS is the send function - aka "MyKind <- some ErrorS"
func (c *SChErrorS) ProvideErrorS(dat []error) {
	// .req
	c.dat <- dat
}

// RequestErrorS is the receive function - aka "some ErrorS <- MyKind"
func (c *SChErrorS) RequestErrorS() (dat []error) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryErrorS is the comma-ok multi-valued form of RequestErrorS and
// reports whether a received value was sent before the ErrorS channel was closed.
func (c *SChErrorS) TryErrorS() (dat []error, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
