// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/GoLangsam/container/ccsafe/dot"
)

// DotChan represents a
// bidirectional
// channel
type DotChan interface {
	DotROnlyChan // aka "<-chan" - receive only
	DotSOnlyChan // aka "chan<-" - send only
}

// DotROnlyChan represents a
// receive-only
// channel
type DotROnlyChan interface {
	RequestDot() (dat *dot.Dot)        // the receive function - aka "MyDot := <-MyDotROnlyChan"
	TryDot() (dat *dot.Dot, open bool) // the multi-valued comma-ok receive function - aka "MyDot, ok := <-MyDotROnlyChan"
}

// DotSOnlyChan represents a
// send-only
// channel
type DotSOnlyChan interface {
	ProvideDot(dat *dot.Dot) // the send function - aka "MyKind <- some Dot"
}

// SChDot is a supply channel
type SChDot struct {
	dat chan *dot.Dot
	// req chan struct{}
}

// MakeSupplyDotChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyDotChan() *SChDot {
	d := new(SChDot)
	d.dat = make(chan *dot.Dot)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyDotBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyDotBuff(cap int) *SChDot {
	d := new(SChDot)
	d.dat = make(chan *dot.Dot, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideDot is the send function - aka "MyKind <- some Dot"
func (c *SChDot) ProvideDot(dat *dot.Dot) {
	// .req
	c.dat <- dat
}

// RequestDot is the receive function - aka "some Dot <- MyKind"
func (c *SChDot) RequestDot() (dat *dot.Dot) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryDot is the comma-ok multi-valued form of RequestDot and
// reports whether a received value was sent before the Dot channel was closed.
func (c *SChDot) TryDot() (dat *dot.Dot, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
