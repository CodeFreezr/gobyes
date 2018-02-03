// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"strings"
)

// ReplacerChan represents a
// bidirectional
// channel
type ReplacerChan interface {
	ReplacerROnlyChan // aka "<-chan" - receive only
	ReplacerSOnlyChan // aka "chan<-" - send only
}

// ReplacerROnlyChan represents a
// receive-only
// channel
type ReplacerROnlyChan interface {
	RequestReplacer() (dat *strings.Replacer)        // the receive function - aka "MyReplacer := <-MyReplacerROnlyChan"
	TryReplacer() (dat *strings.Replacer, open bool) // the multi-valued comma-ok receive function - aka "MyReplacer, ok := <-MyReplacerROnlyChan"
}

// ReplacerSOnlyChan represents a
// send-only
// channel
type ReplacerSOnlyChan interface {
	ProvideReplacer(dat *strings.Replacer) // the send function - aka "MyKind <- some Replacer"
}

// DChReplacer is a demand channel
type DChReplacer struct {
	dat chan *strings.Replacer
	req chan struct{}
}

// MakeDemandReplacerChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandReplacerChan() *DChReplacer {
	d := new(DChReplacer)
	d.dat = make(chan *strings.Replacer)
	d.req = make(chan struct{})
	return d
}

// MakeDemandReplacerBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandReplacerBuff(cap int) *DChReplacer {
	d := new(DChReplacer)
	d.dat = make(chan *strings.Replacer, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideReplacer is the send function - aka "MyKind <- some Replacer"
func (c *DChReplacer) ProvideReplacer(dat *strings.Replacer) {
	<-c.req
	c.dat <- dat
}

// RequestReplacer is the receive function - aka "some Replacer <- MyKind"
func (c *DChReplacer) RequestReplacer() (dat *strings.Replacer) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryReplacer is the comma-ok multi-valued form of RequestReplacer and
// reports whether a received value was sent before the Replacer channel was closed.
func (c *DChReplacer) TryReplacer() (dat *strings.Replacer, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChReplacer is a supply channel
type SChReplacer struct {
	dat chan *strings.Replacer
	// req chan struct{}
}

// MakeSupplyReplacerChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyReplacerChan() *SChReplacer {
	d := new(SChReplacer)
	d.dat = make(chan *strings.Replacer)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyReplacerBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyReplacerBuff(cap int) *SChReplacer {
	d := new(SChReplacer)
	d.dat = make(chan *strings.Replacer, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideReplacer is the send function - aka "MyKind <- some Replacer"
func (c *SChReplacer) ProvideReplacer(dat *strings.Replacer) {
	// .req
	c.dat <- dat
}

// RequestReplacer is the receive function - aka "some Replacer <- MyKind"
func (c *SChReplacer) RequestReplacer() (dat *strings.Replacer) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryReplacer is the comma-ok multi-valued form of RequestReplacer and
// reports whether a received value was sent before the Replacer channel was closed.
func (c *SChReplacer) TryReplacer() (dat *strings.Replacer, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
