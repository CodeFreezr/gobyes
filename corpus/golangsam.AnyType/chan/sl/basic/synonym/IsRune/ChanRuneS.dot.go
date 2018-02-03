// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsRune

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// RuneSChan represents a
// bidirectional
// channel
type RuneSChan interface {
	RuneSROnlyChan // aka "<-chan" - receive only
	RuneSSOnlyChan // aka "chan<-" - send only
}

// RuneSROnlyChan represents a
// receive-only
// channel
type RuneSROnlyChan interface {
	RequestRuneS() (dat []rune)        // the receive function - aka "MyRuneS := <-MyRuneSROnlyChan"
	TryRuneS() (dat []rune, open bool) // the multi-valued comma-ok receive function - aka "MyRuneS, ok := <-MyRuneSROnlyChan"
}

// RuneSSOnlyChan represents a
// send-only
// channel
type RuneSSOnlyChan interface {
	ProvideRuneS(dat []rune) // the send function - aka "MyKind <- some RuneS"
}

// DChRuneS is a demand channel
type DChRuneS struct {
	dat chan []rune
	req chan struct{}
}

// MakeDemandRuneSChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandRuneSChan() *DChRuneS {
	d := new(DChRuneS)
	d.dat = make(chan []rune)
	d.req = make(chan struct{})
	return d
}

// MakeDemandRuneSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandRuneSBuff(cap int) *DChRuneS {
	d := new(DChRuneS)
	d.dat = make(chan []rune, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideRuneS is the send function - aka "MyKind <- some RuneS"
func (c *DChRuneS) ProvideRuneS(dat []rune) {
	<-c.req
	c.dat <- dat
}

// RequestRuneS is the receive function - aka "some RuneS <- MyKind"
func (c *DChRuneS) RequestRuneS() (dat []rune) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryRuneS is the comma-ok multi-valued form of RequestRuneS and
// reports whether a received value was sent before the RuneS channel was closed.
func (c *DChRuneS) TryRuneS() (dat []rune, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChRuneS is a supply channel
type SChRuneS struct {
	dat chan []rune
	// req chan struct{}
}

// MakeSupplyRuneSChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyRuneSChan() *SChRuneS {
	d := new(SChRuneS)
	d.dat = make(chan []rune)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyRuneSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyRuneSBuff(cap int) *SChRuneS {
	d := new(SChRuneS)
	d.dat = make(chan []rune, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideRuneS is the send function - aka "MyKind <- some RuneS"
func (c *SChRuneS) ProvideRuneS(dat []rune) {
	// .req
	c.dat <- dat
}

// RequestRuneS is the receive function - aka "some RuneS <- MyKind"
func (c *SChRuneS) RequestRuneS() (dat []rune) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryRuneS is the comma-ok multi-valued form of RequestRuneS and
// reports whether a received value was sent before the RuneS channel was closed.
func (c *SChRuneS) TryRuneS() (dat []rune, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
