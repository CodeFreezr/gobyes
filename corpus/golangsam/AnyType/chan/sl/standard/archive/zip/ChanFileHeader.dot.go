// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	zip "archive/zip"
)

// FileHeaderChan represents a
// bidirectional
// channel
type FileHeaderChan interface {
	FileHeaderROnlyChan // aka "<-chan" - receive only
	FileHeaderSOnlyChan // aka "chan<-" - send only
}

// FileHeaderROnlyChan represents a
// receive-only
// channel
type FileHeaderROnlyChan interface {
	RequestFileHeader() (dat *zip.FileHeader)        // the receive function - aka "MyFileHeader := <-MyFileHeaderROnlyChan"
	TryFileHeader() (dat *zip.FileHeader, open bool) // the multi-valued comma-ok receive function - aka "MyFileHeader, ok := <-MyFileHeaderROnlyChan"
}

// FileHeaderSOnlyChan represents a
// send-only
// channel
type FileHeaderSOnlyChan interface {
	ProvideFileHeader(dat *zip.FileHeader) // the send function - aka "MyKind <- some FileHeader"
}

// DChFileHeader is a demand channel
type DChFileHeader struct {
	dat chan *zip.FileHeader
	req chan struct{}
}

// MakeDemandFileHeaderChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFileHeaderChan() *DChFileHeader {
	d := new(DChFileHeader)
	d.dat = make(chan *zip.FileHeader)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFileHeaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFileHeaderBuff(cap int) *DChFileHeader {
	d := new(DChFileHeader)
	d.dat = make(chan *zip.FileHeader, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFileHeader is the send function - aka "MyKind <- some FileHeader"
func (c *DChFileHeader) ProvideFileHeader(dat *zip.FileHeader) {
	<-c.req
	c.dat <- dat
}

// RequestFileHeader is the receive function - aka "some FileHeader <- MyKind"
func (c *DChFileHeader) RequestFileHeader() (dat *zip.FileHeader) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFileHeader is the comma-ok multi-valued form of RequestFileHeader and
// reports whether a received value was sent before the FileHeader channel was closed.
func (c *DChFileHeader) TryFileHeader() (dat *zip.FileHeader, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChFileHeader is a supply channel
type SChFileHeader struct {
	dat chan *zip.FileHeader
	// req chan struct{}
}

// MakeSupplyFileHeaderChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyFileHeaderChan() *SChFileHeader {
	d := new(SChFileHeader)
	d.dat = make(chan *zip.FileHeader)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyFileHeaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyFileHeaderBuff(cap int) *SChFileHeader {
	d := new(SChFileHeader)
	d.dat = make(chan *zip.FileHeader, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideFileHeader is the send function - aka "MyKind <- some FileHeader"
func (c *SChFileHeader) ProvideFileHeader(dat *zip.FileHeader) {
	// .req
	c.dat <- dat
}

// RequestFileHeader is the receive function - aka "some FileHeader <- MyKind"
func (c *SChFileHeader) RequestFileHeader() (dat *zip.FileHeader) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryFileHeader is the comma-ok multi-valued form of RequestFileHeader and
// reports whether a received value was sent before the FileHeader channel was closed.
func (c *SChFileHeader) TryFileHeader() (dat *zip.FileHeader, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
