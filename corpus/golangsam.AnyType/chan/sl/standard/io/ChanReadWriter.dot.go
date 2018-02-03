// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ReadWriterChan represents a
// bidirectional
// channel
type ReadWriterChan interface {
	ReadWriterROnlyChan // aka "<-chan" - receive only
	ReadWriterSOnlyChan // aka "chan<-" - send only
}

// ReadWriterROnlyChan represents a
// receive-only
// channel
type ReadWriterROnlyChan interface {
	RequestReadWriter() (dat io.ReadWriter)        // the receive function - aka "MyReadWriter := <-MyReadWriterROnlyChan"
	TryReadWriter() (dat io.ReadWriter, open bool) // the multi-valued comma-ok receive function - aka "MyReadWriter, ok := <-MyReadWriterROnlyChan"
}

// ReadWriterSOnlyChan represents a
// send-only
// channel
type ReadWriterSOnlyChan interface {
	ProvideReadWriter(dat io.ReadWriter) // the send function - aka "MyKind <- some ReadWriter"
}

// DChReadWriter is a demand channel
type DChReadWriter struct {
	dat chan io.ReadWriter
	req chan struct{}
}

// MakeDemandReadWriterChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandReadWriterChan() *DChReadWriter {
	d := new(DChReadWriter)
	d.dat = make(chan io.ReadWriter)
	d.req = make(chan struct{})
	return d
}

// MakeDemandReadWriterBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandReadWriterBuff(cap int) *DChReadWriter {
	d := new(DChReadWriter)
	d.dat = make(chan io.ReadWriter, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideReadWriter is the send function - aka "MyKind <- some ReadWriter"
func (c *DChReadWriter) ProvideReadWriter(dat io.ReadWriter) {
	<-c.req
	c.dat <- dat
}

// RequestReadWriter is the receive function - aka "some ReadWriter <- MyKind"
func (c *DChReadWriter) RequestReadWriter() (dat io.ReadWriter) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryReadWriter is the comma-ok multi-valued form of RequestReadWriter and
// reports whether a received value was sent before the ReadWriter channel was closed.
func (c *DChReadWriter) TryReadWriter() (dat io.ReadWriter, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChReadWriter is a supply channel
type SChReadWriter struct {
	dat chan io.ReadWriter
	// req chan struct{}
}

// MakeSupplyReadWriterChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyReadWriterChan() *SChReadWriter {
	d := new(SChReadWriter)
	d.dat = make(chan io.ReadWriter)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyReadWriterBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyReadWriterBuff(cap int) *SChReadWriter {
	d := new(SChReadWriter)
	d.dat = make(chan io.ReadWriter, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideReadWriter is the send function - aka "MyKind <- some ReadWriter"
func (c *SChReadWriter) ProvideReadWriter(dat io.ReadWriter) {
	// .req
	c.dat <- dat
}

// RequestReadWriter is the receive function - aka "some ReadWriter <- MyKind"
func (c *SChReadWriter) RequestReadWriter() (dat io.ReadWriter) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryReadWriter is the comma-ok multi-valued form of RequestReadWriter and
// reports whether a received value was sent before the ReadWriter channel was closed.
func (c *SChReadWriter) TryReadWriter() (dat io.ReadWriter, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
