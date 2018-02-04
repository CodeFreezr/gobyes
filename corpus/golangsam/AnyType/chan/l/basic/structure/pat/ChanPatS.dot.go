// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pat

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// PatSChan represents a
// bidirectional
// channel
type PatSChan interface {
	PatSROnlyChan // aka "<-chan" - receive only
	PatSSOnlyChan // aka "chan<-" - send only
}

// PatSROnlyChan represents a
// receive-only
// channel
type PatSROnlyChan interface {
	RequestPatS() (dat []struct{})        // the receive function - aka "MyPatS := <-MyPatSROnlyChan"
	TryPatS() (dat []struct{}, open bool) // the multi-valued comma-ok receive function - aka "MyPatS, ok := <-MyPatSROnlyChan"
}

// PatSSOnlyChan represents a
// send-only
// channel
type PatSSOnlyChan interface {
	ProvidePatS(dat []struct{}) // the send function - aka "MyKind <- some PatS"
}

// DChPatS is a demand channel
type DChPatS struct {
	dat chan []struct{}
	req chan struct{}
}

// MakeDemandPatSChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandPatSChan() *DChPatS {
	d := new(DChPatS)
	d.dat = make(chan []struct{})
	d.req = make(chan struct{})
	return d
}

// MakeDemandPatSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandPatSBuff(cap int) *DChPatS {
	d := new(DChPatS)
	d.dat = make(chan []struct{}, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvidePatS is the send function - aka "MyKind <- some PatS"
func (c *DChPatS) ProvidePatS(dat []struct{}) {
	<-c.req
	c.dat <- dat
}

// RequestPatS is the receive function - aka "some PatS <- MyKind"
func (c *DChPatS) RequestPatS() (dat []struct{}) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryPatS is the comma-ok multi-valued form of RequestPatS and
// reports whether a received value was sent before the PatS channel was closed.
func (c *DChPatS) TryPatS() (dat []struct{}, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
