// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// SomeTypeSChan represents a
// bidirectional
// channel
type SomeTypeSChan interface {
	SomeTypeSROnlyChan // aka "<-chan" - receive only
	SomeTypeSSOnlyChan // aka "chan<-" - send only
}

// SomeTypeSROnlyChan represents a
// receive-only
// channel
type SomeTypeSROnlyChan interface {
	RequestSomeTypeS() (dat []SomeType)        // the receive function - aka "MySomeTypeS := <-MySomeTypeSROnlyChan"
	TrySomeTypeS() (dat []SomeType, open bool) // the multi-valued comma-ok receive function - aka "MySomeTypeS, ok := <-MySomeTypeSROnlyChan"
}

// SomeTypeSSOnlyChan represents a
// send-only
// channel
type SomeTypeSSOnlyChan interface {
	ProvideSomeTypeS(dat []SomeType) // the send function - aka "MyKind <- some SomeTypeS"
}

// DChSomeTypeS is a demand channel
type DChSomeTypeS struct {
	dat chan []SomeType
	req chan struct{}
}

// MakeDemandSomeTypeSChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandSomeTypeSChan() *DChSomeTypeS {
	d := new(DChSomeTypeS)
	d.dat = make(chan []SomeType)
	d.req = make(chan struct{})
	return d
}

// MakeDemandSomeTypeSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandSomeTypeSBuff(cap int) *DChSomeTypeS {
	d := new(DChSomeTypeS)
	d.dat = make(chan []SomeType, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideSomeTypeS is the send function - aka "MyKind <- some SomeTypeS"
func (c *DChSomeTypeS) ProvideSomeTypeS(dat []SomeType) {
	<-c.req
	c.dat <- dat
}

// RequestSomeTypeS is the receive function - aka "some SomeTypeS <- MyKind"
func (c *DChSomeTypeS) RequestSomeTypeS() (dat []SomeType) {
	c.req <- struct{}{}
	return <-c.dat
}

// TrySomeTypeS is the comma-ok multi-valued form of RequestSomeTypeS and
// reports whether a received value was sent before the SomeTypeS channel was closed.
func (c *DChSomeTypeS) TrySomeTypeS() (dat []SomeType, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChSomeTypeS is a supply channel
type SChSomeTypeS struct {
	dat chan []SomeType
	// req chan struct{}
}

// MakeSupplySomeTypeSChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplySomeTypeSChan() *SChSomeTypeS {
	d := new(SChSomeTypeS)
	d.dat = make(chan []SomeType)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplySomeTypeSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplySomeTypeSBuff(cap int) *SChSomeTypeS {
	d := new(SChSomeTypeS)
	d.dat = make(chan []SomeType, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideSomeTypeS is the send function - aka "MyKind <- some SomeTypeS"
func (c *SChSomeTypeS) ProvideSomeTypeS(dat []SomeType) {
	// .req
	c.dat <- dat
}

// RequestSomeTypeS is the receive function - aka "some SomeTypeS <- MyKind"
func (c *SChSomeTypeS) RequestSomeTypeS() (dat []SomeType) {
	// eq <- struct{}{}
	return <-c.dat
}

// TrySomeTypeS is the comma-ok multi-valued form of RequestSomeTypeS and
// reports whether a received value was sent before the SomeTypeS channel was closed.
func (c *SChSomeTypeS) TrySomeTypeS() (dat []SomeType, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
