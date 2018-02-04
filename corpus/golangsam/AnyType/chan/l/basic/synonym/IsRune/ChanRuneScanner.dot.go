// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsRune

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// RuneScannerChan represents a
// bidirectional
// channel
type RuneScannerChan interface {
	RuneScannerROnlyChan // aka "<-chan" - receive only
	RuneScannerSOnlyChan // aka "chan<-" - send only
}

// RuneScannerROnlyChan represents a
// receive-only
// channel
type RuneScannerROnlyChan interface {
	RequestRuneScanner() (dat io.RuneScanner)        // the receive function - aka "MyRuneScanner := <-MyRuneScannerROnlyChan"
	TryRuneScanner() (dat io.RuneScanner, open bool) // the multi-valued comma-ok receive function - aka "MyRuneScanner, ok := <-MyRuneScannerROnlyChan"
}

// RuneScannerSOnlyChan represents a
// send-only
// channel
type RuneScannerSOnlyChan interface {
	ProvideRuneScanner(dat io.RuneScanner) // the send function - aka "MyKind <- some RuneScanner"
}

// DChRuneScanner is a demand channel
type DChRuneScanner struct {
	dat chan io.RuneScanner
	req chan struct{}
}

// MakeDemandRuneScannerChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandRuneScannerChan() *DChRuneScanner {
	d := new(DChRuneScanner)
	d.dat = make(chan io.RuneScanner)
	d.req = make(chan struct{})
	return d
}

// MakeDemandRuneScannerBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandRuneScannerBuff(cap int) *DChRuneScanner {
	d := new(DChRuneScanner)
	d.dat = make(chan io.RuneScanner, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideRuneScanner is the send function - aka "MyKind <- some RuneScanner"
func (c *DChRuneScanner) ProvideRuneScanner(dat io.RuneScanner) {
	<-c.req
	c.dat <- dat
}

// RequestRuneScanner is the receive function - aka "some RuneScanner <- MyKind"
func (c *DChRuneScanner) RequestRuneScanner() (dat io.RuneScanner) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryRuneScanner is the comma-ok multi-valued form of RequestRuneScanner and
// reports whether a received value was sent before the RuneScanner channel was closed.
func (c *DChRuneScanner) TryRuneScanner() (dat io.RuneScanner, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
