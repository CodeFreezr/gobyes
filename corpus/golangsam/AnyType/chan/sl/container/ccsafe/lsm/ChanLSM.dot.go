// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsm

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/GoLangsam/container/ccsafe/lsm"
)

// LSMChan represents a
// bidirectional
// channel
type LSMChan interface {
	LSMROnlyChan // aka "<-chan" - receive only
	LSMSOnlyChan // aka "chan<-" - send only
}

// LSMROnlyChan represents a
// receive-only
// channel
type LSMROnlyChan interface {
	RequestLSM() (dat *lsm.LazyStringerMap)        // the receive function - aka "MyLSM := <-MyLSMROnlyChan"
	TryLSM() (dat *lsm.LazyStringerMap, open bool) // the multi-valued comma-ok receive function - aka "MyLSM, ok := <-MyLSMROnlyChan"
}

// LSMSOnlyChan represents a
// send-only
// channel
type LSMSOnlyChan interface {
	ProvideLSM(dat *lsm.LazyStringerMap) // the send function - aka "MyKind <- some LSM"
}

// DChLSM is a demand channel
type DChLSM struct {
	dat chan *lsm.LazyStringerMap
	req chan struct{}
}

// MakeDemandLSMChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandLSMChan() *DChLSM {
	d := new(DChLSM)
	d.dat = make(chan *lsm.LazyStringerMap)
	d.req = make(chan struct{})
	return d
}

// MakeDemandLSMBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandLSMBuff(cap int) *DChLSM {
	d := new(DChLSM)
	d.dat = make(chan *lsm.LazyStringerMap, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideLSM is the send function - aka "MyKind <- some LSM"
func (c *DChLSM) ProvideLSM(dat *lsm.LazyStringerMap) {
	<-c.req
	c.dat <- dat
}

// RequestLSM is the receive function - aka "some LSM <- MyKind"
func (c *DChLSM) RequestLSM() (dat *lsm.LazyStringerMap) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryLSM is the comma-ok multi-valued form of RequestLSM and
// reports whether a received value was sent before the LSM channel was closed.
func (c *DChLSM) TryLSM() (dat *lsm.LazyStringerMap, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChLSM is a supply channel
type SChLSM struct {
	dat chan *lsm.LazyStringerMap
	// req chan struct{}
}

// MakeSupplyLSMChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyLSMChan() *SChLSM {
	d := new(SChLSM)
	d.dat = make(chan *lsm.LazyStringerMap)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyLSMBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyLSMBuff(cap int) *SChLSM {
	d := new(SChLSM)
	d.dat = make(chan *lsm.LazyStringerMap, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideLSM is the send function - aka "MyKind <- some LSM"
func (c *SChLSM) ProvideLSM(dat *lsm.LazyStringerMap) {
	// .req
	c.dat <- dat
}

// RequestLSM is the receive function - aka "some LSM <- MyKind"
func (c *SChLSM) RequestLSM() (dat *lsm.LazyStringerMap) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryLSM is the comma-ok multi-valued form of RequestLSM and
// reports whether a received value was sent before the LSM channel was closed.
func (c *SChLSM) TryLSM() (dat *lsm.LazyStringerMap, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
