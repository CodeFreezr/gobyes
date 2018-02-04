// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// StringSChan represents a
// bidirectional
// channel
type StringSChan interface {
	StringSROnlyChan // aka "<-chan" - receive only
	StringSSOnlyChan // aka "chan<-" - send only
}

// StringSROnlyChan represents a
// receive-only
// channel
type StringSROnlyChan interface {
	RequestStringS() (dat []string)        // the receive function - aka "MyStringS := <-MyStringSROnlyChan"
	TryStringS() (dat []string, open bool) // the multi-valued comma-ok receive function - aka "MyStringS, ok := <-MyStringSROnlyChan"
}

// StringSSOnlyChan represents a
// send-only
// channel
type StringSSOnlyChan interface {
	ProvideStringS(dat []string) // the send function - aka "MyKind <- some StringS"
}

// DChStringS is a demand channel
type DChStringS struct {
	dat chan []string
	req chan struct{}
}

// MakeDemandStringSChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandStringSChan() *DChStringS {
	d := new(DChStringS)
	d.dat = make(chan []string)
	d.req = make(chan struct{})
	return d
}

// MakeDemandStringSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandStringSBuff(cap int) *DChStringS {
	d := new(DChStringS)
	d.dat = make(chan []string, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideStringS is the send function - aka "MyKind <- some StringS"
func (c *DChStringS) ProvideStringS(dat []string) {
	<-c.req
	c.dat <- dat
}

// RequestStringS is the receive function - aka "some StringS <- MyKind"
func (c *DChStringS) RequestStringS() (dat []string) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryStringS is the comma-ok multi-valued form of RequestStringS and
// reports whether a received value was sent before the StringS channel was closed.
func (c *DChStringS) TryStringS() (dat []string, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChStringS is a supply channel
type SChStringS struct {
	dat chan []string
	// req chan struct{}
}

// MakeSupplyStringSChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyStringSChan() *SChStringS {
	d := new(SChStringS)
	d.dat = make(chan []string)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyStringSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyStringSBuff(cap int) *SChStringS {
	d := new(SChStringS)
	d.dat = make(chan []string, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideStringS is the send function - aka "MyKind <- some StringS"
func (c *SChStringS) ProvideStringS(dat []string) {
	// .req
	c.dat <- dat
}

// RequestStringS is the receive function - aka "some StringS <- MyKind"
func (c *SChStringS) RequestStringS() (dat []string) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryStringS is the comma-ok multi-valued form of RequestStringS and
// reports whether a received value was sent before the StringS channel was closed.
func (c *SChStringS) TryStringS() (dat []string, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
