// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	list "container/list"
)

// ElementChan represents a
// bidirectional
// channel
type ElementChan interface {
	ElementROnlyChan // aka "<-chan" - receive only
	ElementSOnlyChan // aka "chan<-" - send only
}

// ElementROnlyChan represents a
// receive-only
// channel
type ElementROnlyChan interface {
	RequestElement() (dat *list.Element)        // the receive function - aka "MyElement := <-MyElementROnlyChan"
	TryElement() (dat *list.Element, open bool) // the multi-valued comma-ok receive function - aka "MyElement, ok := <-MyElementROnlyChan"
}

// ElementSOnlyChan represents a
// send-only
// channel
type ElementSOnlyChan interface {
	ProvideElement(dat *list.Element) // the send function - aka "MyKind <- some Element"
}

// DChElement is a demand channel
type DChElement struct {
	dat chan *list.Element
	req chan struct{}
}

// MakeDemandElementChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandElementChan() *DChElement {
	d := new(DChElement)
	d.dat = make(chan *list.Element)
	d.req = make(chan struct{})
	return d
}

// MakeDemandElementBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandElementBuff(cap int) *DChElement {
	d := new(DChElement)
	d.dat = make(chan *list.Element, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideElement is the send function - aka "MyKind <- some Element"
func (c *DChElement) ProvideElement(dat *list.Element) {
	<-c.req
	c.dat <- dat
}

// RequestElement is the receive function - aka "some Element <- MyKind"
func (c *DChElement) RequestElement() (dat *list.Element) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryElement is the comma-ok multi-valued form of RequestElement and
// reports whether a received value was sent before the Element channel was closed.
func (c *DChElement) TryElement() (dat *list.Element, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChElement is a supply channel
type SChElement struct {
	dat chan *list.Element
	// req chan struct{}
}

// MakeSupplyElementChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyElementChan() *SChElement {
	d := new(SChElement)
	d.dat = make(chan *list.Element)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyElementBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyElementBuff(cap int) *SChElement {
	d := new(SChElement)
	d.dat = make(chan *list.Element, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideElement is the send function - aka "MyKind <- some Element"
func (c *SChElement) ProvideElement(dat *list.Element) {
	// .req
	c.dat <- dat
}

// RequestElement is the receive function - aka "some Element <- MyKind"
func (c *SChElement) RequestElement() (dat *list.Element) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryElement is the comma-ok multi-valued form of RequestElement and
// reports whether a received value was sent before the Element channel was closed.
func (c *SChElement) TryElement() (dat *list.Element, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
