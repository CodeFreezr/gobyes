// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ByteScannerChan represents a
// bidirectional
// channel
type ByteScannerChan interface {
	ByteScannerROnlyChan // aka "<-chan" - receive only
	ByteScannerSOnlyChan // aka "chan<-" - send only
}

// ByteScannerROnlyChan represents a
// receive-only
// channel
type ByteScannerROnlyChan interface {
	RequestByteScanner() (dat io.ByteScanner)        // the receive function - aka "MyByteScanner := <-MyByteScannerROnlyChan"
	TryByteScanner() (dat io.ByteScanner, open bool) // the multi-valued comma-ok receive function - aka "MyByteScanner, ok := <-MyByteScannerROnlyChan"
}

// ByteScannerSOnlyChan represents a
// send-only
// channel
type ByteScannerSOnlyChan interface {
	ProvideByteScanner(dat io.ByteScanner) // the send function - aka "MyKind <- some ByteScanner"
}

// DChByteScanner is a demand channel
type DChByteScanner struct {
	dat chan io.ByteScanner
	req chan struct{}
}

// MakeDemandByteScannerChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandByteScannerChan() *DChByteScanner {
	d := new(DChByteScanner)
	d.dat = make(chan io.ByteScanner)
	d.req = make(chan struct{})
	return d
}

// MakeDemandByteScannerBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandByteScannerBuff(cap int) *DChByteScanner {
	d := new(DChByteScanner)
	d.dat = make(chan io.ByteScanner, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideByteScanner is the send function - aka "MyKind <- some ByteScanner"
func (c *DChByteScanner) ProvideByteScanner(dat io.ByteScanner) {
	<-c.req
	c.dat <- dat
}

// RequestByteScanner is the receive function - aka "some ByteScanner <- MyKind"
func (c *DChByteScanner) RequestByteScanner() (dat io.ByteScanner) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryByteScanner is the comma-ok multi-valued form of RequestByteScanner and
// reports whether a received value was sent before the ByteScanner channel was closed.
func (c *DChByteScanner) TryByteScanner() (dat io.ByteScanner, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChByteScanner is a supply channel
type SChByteScanner struct {
	dat chan io.ByteScanner
	// req chan struct{}
}

// MakeSupplyByteScannerChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyByteScannerChan() *SChByteScanner {
	d := new(SChByteScanner)
	d.dat = make(chan io.ByteScanner)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyByteScannerBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyByteScannerBuff(cap int) *SChByteScanner {
	d := new(SChByteScanner)
	d.dat = make(chan io.ByteScanner, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideByteScanner is the send function - aka "MyKind <- some ByteScanner"
func (c *SChByteScanner) ProvideByteScanner(dat io.ByteScanner) {
	// .req
	c.dat <- dat
}

// RequestByteScanner is the receive function - aka "some ByteScanner <- MyKind"
func (c *SChByteScanner) RequestByteScanner() (dat io.ByteScanner) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryByteScanner is the comma-ok multi-valued form of RequestByteScanner and
// reports whether a received value was sent before the ByteScanner channel was closed.
func (c *SChByteScanner) TryByteScanner() (dat io.ByteScanner, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
