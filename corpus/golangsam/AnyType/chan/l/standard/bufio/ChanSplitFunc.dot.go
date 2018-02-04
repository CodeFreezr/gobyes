// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	bufio "bufio"
)

// SplitFuncChan represents a
// bidirectional
// channel
type SplitFuncChan interface {
	SplitFuncROnlyChan // aka "<-chan" - receive only
	SplitFuncSOnlyChan // aka "chan<-" - send only
}

// SplitFuncROnlyChan represents a
// receive-only
// channel
type SplitFuncROnlyChan interface {
	RequestSplitFunc() (dat bufio.SplitFunc)        // the receive function - aka "MySplitFunc := <-MySplitFuncROnlyChan"
	TrySplitFunc() (dat bufio.SplitFunc, open bool) // the multi-valued comma-ok receive function - aka "MySplitFunc, ok := <-MySplitFuncROnlyChan"
}

// SplitFuncSOnlyChan represents a
// send-only
// channel
type SplitFuncSOnlyChan interface {
	ProvideSplitFunc(dat bufio.SplitFunc) // the send function - aka "MyKind <- some SplitFunc"
}

// DChSplitFunc is a demand channel
type DChSplitFunc struct {
	dat chan bufio.SplitFunc
	req chan struct{}
}

// MakeDemandSplitFuncChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandSplitFuncChan() *DChSplitFunc {
	d := new(DChSplitFunc)
	d.dat = make(chan bufio.SplitFunc)
	d.req = make(chan struct{})
	return d
}

// MakeDemandSplitFuncBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandSplitFuncBuff(cap int) *DChSplitFunc {
	d := new(DChSplitFunc)
	d.dat = make(chan bufio.SplitFunc, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideSplitFunc is the send function - aka "MyKind <- some SplitFunc"
func (c *DChSplitFunc) ProvideSplitFunc(dat bufio.SplitFunc) {
	<-c.req
	c.dat <- dat
}

// RequestSplitFunc is the receive function - aka "some SplitFunc <- MyKind"
func (c *DChSplitFunc) RequestSplitFunc() (dat bufio.SplitFunc) {
	c.req <- struct{}{}
	return <-c.dat
}

// TrySplitFunc is the comma-ok multi-valued form of RequestSplitFunc and
// reports whether a received value was sent before the SplitFunc channel was closed.
func (c *DChSplitFunc) TrySplitFunc() (dat bufio.SplitFunc, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
