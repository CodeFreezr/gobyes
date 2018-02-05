// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dotpath

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/GoLangsam/container/ccsafe/dotpath"
)

// DotPathChan represents a
// bidirectional
// channel
type DotPathChan interface {
	DotPathROnlyChan // aka "<-chan" - receive only
	DotPathSOnlyChan // aka "chan<-" - send only
}

// DotPathROnlyChan represents a
// receive-only
// channel
type DotPathROnlyChan interface {
	RequestDotPath() (dat *dotpath.DotPath)        // the receive function - aka "MyDotPath := <-MyDotPathROnlyChan"
	TryDotPath() (dat *dotpath.DotPath, open bool) // the multi-valued comma-ok receive function - aka "MyDotPath, ok := <-MyDotPathROnlyChan"
}

// DotPathSOnlyChan represents a
// send-only
// channel
type DotPathSOnlyChan interface {
	ProvideDotPath(dat *dotpath.DotPath) // the send function - aka "MyKind <- some DotPath"
}

// DChDotPath is a demand channel
type DChDotPath struct {
	dat chan *dotpath.DotPath
	req chan struct{}
}

// MakeDemandDotPathChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandDotPathChan() *DChDotPath {
	d := new(DChDotPath)
	d.dat = make(chan *dotpath.DotPath)
	d.req = make(chan struct{})
	return d
}

// MakeDemandDotPathBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandDotPathBuff(cap int) *DChDotPath {
	d := new(DChDotPath)
	d.dat = make(chan *dotpath.DotPath, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideDotPath is the send function - aka "MyKind <- some DotPath"
func (c *DChDotPath) ProvideDotPath(dat *dotpath.DotPath) {
	<-c.req
	c.dat <- dat
}

// RequestDotPath is the receive function - aka "some DotPath <- MyKind"
func (c *DChDotPath) RequestDotPath() (dat *dotpath.DotPath) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryDotPath is the comma-ok multi-valued form of RequestDotPath and
// reports whether a received value was sent before the DotPath channel was closed.
func (c *DChDotPath) TryDotPath() (dat *dotpath.DotPath, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChDotPath is a supply channel
type SChDotPath struct {
	dat chan *dotpath.DotPath
	// req chan struct{}
}

// MakeSupplyDotPathChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyDotPathChan() *SChDotPath {
	d := new(SChDotPath)
	d.dat = make(chan *dotpath.DotPath)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyDotPathBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyDotPathBuff(cap int) *SChDotPath {
	d := new(SChDotPath)
	d.dat = make(chan *dotpath.DotPath, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideDotPath is the send function - aka "MyKind <- some DotPath"
func (c *SChDotPath) ProvideDotPath(dat *dotpath.DotPath) {
	// .req
	c.dat <- dat
}

// RequestDotPath is the receive function - aka "some DotPath <- MyKind"
func (c *SChDotPath) RequestDotPath() (dat *dotpath.DotPath) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryDotPath is the comma-ok multi-valued form of RequestDotPath and
// reports whether a received value was sent before the DotPath channel was closed.
func (c *SChDotPath) TryDotPath() (dat *dotpath.DotPath, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
