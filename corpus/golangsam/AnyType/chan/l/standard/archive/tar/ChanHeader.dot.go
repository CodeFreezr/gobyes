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

// DChHeader is a demand channel
type DChHeader struct {
	dat chan *tar.Header
	req chan struct{}
}

// MakeDemandHeaderChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandHeaderChan() *DChHeader {
	d := new(DChHeader)
	d.dat = make(chan *tar.Header)
	d.req = make(chan struct{})
	return d
}

// MakeDemandHeaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandHeaderBuff(cap int) *DChHeader {
	d := new(DChHeader)
	d.dat = make(chan *tar.Header, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideHeader is the send function - aka "MyKind <- some Header"
func (c *DChHeader) ProvideHeader(dat *tar.Header) {
	<-c.req
	c.dat <- dat
}

// RequestHeader is the receive function - aka "some Header <- MyKind"
func (c *DChHeader) RequestHeader() (dat *tar.Header) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryHeader is the comma-ok multi-valued form of RequestHeader and
// reports whether a received value was sent before the Header channel was closed.
func (c *DChHeader) TryHeader() (dat *tar.Header, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
