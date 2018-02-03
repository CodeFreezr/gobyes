// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

// FileInfoChan represents a
// bidirectional
// channel
type FileInfoChan interface {
	FileInfoROnlyChan // aka "<-chan" - receive only
	FileInfoSOnlyChan // aka "chan<-" - send only
}

// FileInfoROnlyChan represents a
// receive-only
// channel
type FileInfoROnlyChan interface {
	RequestFileInfo() (dat os.FileInfo)        // the receive function - aka "MyFileInfo := <-MyFileInfoROnlyChan"
	TryFileInfo() (dat os.FileInfo, open bool) // the multi-valued comma-ok receive function - aka "MyFileInfo, ok := <-MyFileInfoROnlyChan"
}

// FileInfoSOnlyChan represents a
// send-only
// channel
type FileInfoSOnlyChan interface {
	ProvideFileInfo(dat os.FileInfo) // the send function - aka "MyKind <- some FileInfo"
}

// DChFileInfo is a demand channel
type DChFileInfo struct {
	dat chan os.FileInfo
	req chan struct{}
}

// MakeDemandFileInfoChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFileInfoChan() *DChFileInfo {
	d := new(DChFileInfo)
	d.dat = make(chan os.FileInfo)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFileInfoBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFileInfoBuff(cap int) *DChFileInfo {
	d := new(DChFileInfo)
	d.dat = make(chan os.FileInfo, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFileInfo is the send function - aka "MyKind <- some FileInfo"
func (c *DChFileInfo) ProvideFileInfo(dat os.FileInfo) {
	<-c.req
	c.dat <- dat
}

// RequestFileInfo is the receive function - aka "some FileInfo <- MyKind"
func (c *DChFileInfo) RequestFileInfo() (dat os.FileInfo) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFileInfo is the comma-ok multi-valued form of RequestFileInfo and
// reports whether a received value was sent before the FileInfo channel was closed.
func (c *DChFileInfo) TryFileInfo() (dat os.FileInfo, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
