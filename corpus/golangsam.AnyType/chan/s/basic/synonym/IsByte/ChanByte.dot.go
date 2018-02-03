// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// ByteChan represents a
// bidirectional
// channel
type ByteChan interface {
	ByteROnlyChan // aka "<-chan" - receive only
	ByteSOnlyChan // aka "chan<-" - send only
}

// ByteROnlyChan represents a
// receive-only
// channel
type ByteROnlyChan interface {
	RequestByte() (dat byte)        // the receive function - aka "MyByte := <-MyByteROnlyChan"
	TryByte() (dat byte, open bool) // the multi-valued comma-ok receive function - aka "MyByte, ok := <-MyByteROnlyChan"
}

// ByteSOnlyChan represents a
// send-only
// channel
type ByteSOnlyChan interface {
	ProvideByte(dat byte) // the send function - aka "MyKind <- some Byte"
}

// SChByte is a supply channel
type SChByte struct {
	dat chan byte
	// req chan struct{}
}

// MakeSupplyByteChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyByteChan() *SChByte {
	d := new(SChByte)
	d.dat = make(chan byte)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyByteBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyByteBuff(cap int) *SChByte {
	d := new(SChByte)
	d.dat = make(chan byte, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideByte is the send function - aka "MyKind <- some Byte"
func (c *SChByte) ProvideByte(dat byte) {
	// .req
	c.dat <- dat
}

// RequestByte is the receive function - aka "some Byte <- MyKind"
func (c *SChByte) RequestByte() (dat byte) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryByte is the comma-ok multi-valued form of RequestByte and
// reports whether a received value was sent before the Byte channel was closed.
func (c *SChByte) TryByte() (dat byte, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
