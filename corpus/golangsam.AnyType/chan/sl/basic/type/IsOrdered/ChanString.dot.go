// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// StringChan represents a
// bidirectional
// channel
type StringChan interface {
	StringROnlyChan // aka "<-chan" - receive only
	StringSOnlyChan // aka "chan<-" - send only
}

// StringROnlyChan represents a
// receive-only
// channel
type StringROnlyChan interface {
	RequestString() (dat string)        // the receive function - aka "MyString := <-MyStringROnlyChan"
	TryString() (dat string, open bool) // the multi-valued comma-ok receive function - aka "MyString, ok := <-MyStringROnlyChan"
}

// StringSOnlyChan represents a
// send-only
// channel
type StringSOnlyChan interface {
	ProvideString(dat string) // the send function - aka "MyKind <- some String"
}

// DChString is a demand channel
type DChString struct {
	dat chan string
	req chan struct{}
}

// MakeDemandStringChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandStringChan() *DChString {
	d := new(DChString)
	d.dat = make(chan string)
	d.req = make(chan struct{})
	return d
}

// MakeDemandStringBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandStringBuff(cap int) *DChString {
	d := new(DChString)
	d.dat = make(chan string, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideString is the send function - aka "MyKind <- some String"
func (c *DChString) ProvideString(dat string) {
	<-c.req
	c.dat <- dat
}

// RequestString is the receive function - aka "some String <- MyKind"
func (c *DChString) RequestString() (dat string) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryString is the comma-ok multi-valued form of RequestString and
// reports whether a received value was sent before the String channel was closed.
func (c *DChString) TryString() (dat string, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChString is a supply channel
type SChString struct {
	dat chan string
	// req chan struct{}
}

// MakeSupplyStringChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyStringChan() *SChString {
	d := new(SChString)
	d.dat = make(chan string)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyStringBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyStringBuff(cap int) *SChString {
	d := new(SChString)
	d.dat = make(chan string, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideString is the send function - aka "MyKind <- some String"
func (c *SChString) ProvideString(dat string) {
	// .req
	c.dat <- dat
}

// RequestString is the receive function - aka "some String <- MyKind"
func (c *SChString) RequestString() (dat string) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryString is the comma-ok multi-valued form of RequestString and
// reports whether a received value was sent before the String channel was closed.
func (c *SChString) TryString() (dat string, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
