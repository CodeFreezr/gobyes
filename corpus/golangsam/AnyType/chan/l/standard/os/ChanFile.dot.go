// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

// FileChan represents a
// bidirectional
// channel
type FileChan interface {
	FileROnlyChan // aka "<-chan" - receive only
	FileSOnlyChan // aka "chan<-" - send only
}

// FileROnlyChan represents a
// receive-only
// channel
type FileROnlyChan interface {
	RequestFile() (dat *os.File)        // the receive function - aka "MyFile := <-MyFileROnlyChan"
	TryFile() (dat *os.File, open bool) // the multi-valued comma-ok receive function - aka "MyFile, ok := <-MyFileROnlyChan"
}

// FileSOnlyChan represents a
// send-only
// channel
type FileSOnlyChan interface {
	ProvideFile(dat *os.File) // the send function - aka "MyKind <- some File"
}

// DChFile is a demand channel
type DChFile struct {
	dat chan *os.File
	req chan struct{}
}

// MakeDemandFileChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFileChan() *DChFile {
	d := new(DChFile)
	d.dat = make(chan *os.File)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFileBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFileBuff(cap int) *DChFile {
	d := new(DChFile)
	d.dat = make(chan *os.File, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFile is the send function - aka "MyKind <- some File"
func (c *DChFile) ProvideFile(dat *os.File) {
	<-c.req
	c.dat <- dat
}

// RequestFile is the receive function - aka "some File <- MyKind"
func (c *DChFile) RequestFile() (dat *os.File) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFile is the comma-ok multi-valued form of RequestFile and
// reports whether a received value was sent before the File channel was closed.
func (c *DChFile) TryFile() (dat *os.File, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
