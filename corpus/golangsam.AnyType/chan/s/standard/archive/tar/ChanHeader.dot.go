// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	tar "archive/tar"
)

// HeaderChan represents a
// bidirectional
// channel
type HeaderChan interface {
	HeaderROnlyChan // aka "<-chan" - receive only
	HeaderSOnlyChan // aka "chan<-" - send only
}

// HeaderROnlyChan represents a
// receive-only
// channel
type HeaderROnlyChan interface {
	RequestHeader() (dat *tar.Header)        // the receive function - aka "MyHeader := <-MyHeaderROnlyChan"
	TryHeader() (dat *tar.Header, open bool) // the multi-valued comma-ok receive function - aka "MyHeader, ok := <-MyHeaderROnlyChan"
}

// HeaderSOnlyChan represents a
// send-only
// channel
type HeaderSOnlyChan interface {
	ProvideHeader(dat *tar.Header) // the send function - aka "MyKind <- some Header"
}

// SChHeader is a supply channel
type SChHeader struct {
	dat chan *tar.Header
	// req chan struct{}
}

// MakeSupplyHeaderChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyHeaderChan() *SChHeader {
	d := new(SChHeader)
	d.dat = make(chan *tar.Header)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyHeaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyHeaderBuff(cap int) *SChHeader {
	d := new(SChHeader)
	d.dat = make(chan *tar.Header, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideHeader is the send function - aka "MyKind <- some Header"
func (c *SChHeader) ProvideHeader(dat *tar.Header) {
	// .req
	c.dat <- dat
}

// RequestHeader is the receive function - aka "some Header <- MyKind"
func (c *SChHeader) RequestHeader() (dat *tar.Header) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryHeader is the comma-ok multi-valued form of RequestHeader and
// reports whether a received value was sent before the Header channel was closed.
func (c *SChHeader) TryHeader() (dat *tar.Header, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
