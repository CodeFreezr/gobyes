// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ReadCloserChan represents a
// bidirectional
// channel
type ReadCloserChan interface {
	ReadCloserROnlyChan // aka "<-chan" - receive only
	ReadCloserSOnlyChan // aka "chan<-" - send only
}

// ReadCloserROnlyChan represents a
// receive-only
// channel
type ReadCloserROnlyChan interface {
	RequestReadCloser() (dat io.ReadCloser)        // the receive function - aka "MyReadCloser := <-MyReadCloserROnlyChan"
	TryReadCloser() (dat io.ReadCloser, open bool) // the multi-valued comma-ok receive function - aka "MyReadCloser, ok := <-MyReadCloserROnlyChan"
}

// ReadCloserSOnlyChan represents a
// send-only
// channel
type ReadCloserSOnlyChan interface {
	ProvideReadCloser(dat io.ReadCloser) // the send function - aka "MyKind <- some ReadCloser"
}

// DChReadCloser is a demand channel
type DChReadCloser struct {
	dat chan io.ReadCloser
	req chan struct{}
}

// MakeDemandReadCloserChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandReadCloserChan() *DChReadCloser {
	d := new(DChReadCloser)
	d.dat = make(chan io.ReadCloser)
	d.req = make(chan struct{})
	return d
}

// MakeDemandReadCloserBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandReadCloserBuff(cap int) *DChReadCloser {
	d := new(DChReadCloser)
	d.dat = make(chan io.ReadCloser, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideReadCloser is the send function - aka "MyKind <- some ReadCloser"
func (c *DChReadCloser) ProvideReadCloser(dat io.ReadCloser) {
	<-c.req
	c.dat <- dat
}

// RequestReadCloser is the receive function - aka "some ReadCloser <- MyKind"
func (c *DChReadCloser) RequestReadCloser() (dat io.ReadCloser) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryReadCloser is the comma-ok multi-valued form of RequestReadCloser and
// reports whether a received value was sent before the ReadCloser channel was closed.
func (c *DChReadCloser) TryReadCloser() (dat io.ReadCloser, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
