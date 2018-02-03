// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

// SignalChan represents a
// bidirectional
// channel
type SignalChan interface {
	SignalROnlyChan // aka "<-chan" - receive only
	SignalSOnlyChan // aka "chan<-" - send only
}

// SignalROnlyChan represents a
// receive-only
// channel
type SignalROnlyChan interface {
	RequestSignal() (dat os.Signal)        // the receive function - aka "MySignal := <-MySignalROnlyChan"
	TrySignal() (dat os.Signal, open bool) // the multi-valued comma-ok receive function - aka "MySignal, ok := <-MySignalROnlyChan"
}

// SignalSOnlyChan represents a
// send-only
// channel
type SignalSOnlyChan interface {
	ProvideSignal(dat os.Signal) // the send function - aka "MyKind <- some Signal"
}

// DChSignal is a demand channel
type DChSignal struct {
	dat chan os.Signal
	req chan struct{}
}

// MakeDemandSignalChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandSignalChan() *DChSignal {
	d := new(DChSignal)
	d.dat = make(chan os.Signal)
	d.req = make(chan struct{})
	return d
}

// MakeDemandSignalBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandSignalBuff(cap int) *DChSignal {
	d := new(DChSignal)
	d.dat = make(chan os.Signal, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideSignal is the send function - aka "MyKind <- some Signal"
func (c *DChSignal) ProvideSignal(dat os.Signal) {
	<-c.req
	c.dat <- dat
}

// RequestSignal is the receive function - aka "some Signal <- MyKind"
func (c *DChSignal) RequestSignal() (dat os.Signal) {
	c.req <- struct{}{}
	return <-c.dat
}

// TrySignal is the comma-ok multi-valued form of RequestSignal and
// reports whether a received value was sent before the Signal channel was closed.
func (c *DChSignal) TrySignal() (dat os.Signal, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChSignal is a supply channel
type SChSignal struct {
	dat chan os.Signal
	// req chan struct{}
}

// MakeSupplySignalChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplySignalChan() *SChSignal {
	d := new(SChSignal)
	d.dat = make(chan os.Signal)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplySignalBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplySignalBuff(cap int) *SChSignal {
	d := new(SChSignal)
	d.dat = make(chan os.Signal, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideSignal is the send function - aka "MyKind <- some Signal"
func (c *SChSignal) ProvideSignal(dat os.Signal) {
	// .req
	c.dat <- dat
}

// RequestSignal is the receive function - aka "some Signal <- MyKind"
func (c *SChSignal) RequestSignal() (dat os.Signal) {
	// eq <- struct{}{}
	return <-c.dat
}

// TrySignal is the comma-ok multi-valued form of RequestSignal and
// reports whether a received value was sent before the Signal channel was closed.
func (c *SChSignal) TrySignal() (dat os.Signal, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
