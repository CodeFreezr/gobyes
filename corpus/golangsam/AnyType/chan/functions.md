# Generic functions

Just a short overview of the generic functions we generate and provide across all [flavours](flavours.md).

Hint: You may like to look at the related code sniplets - the `*.tmpl` files in [flavours](flavours.md) [`ssss`](./ssss/) and [`sss`](./sss/) are a good point to start. 

## Creators
* Make _Type_ ()  <-chan _type_

## Generators
Hint: A [Generator](https://en.wikipedia.org/wiki/Generator_(computer_programming)) yields a sequence of values ... one at a time.
* Chan _Type_ s (inp... _type_ ) <-chan _type_
* Chan _Type_ S (inp [] _type_ ) <-chan _type_

For _Func_ in _pack_ with signature func ( [args] ) ( _type_ [, error]/[, more] )

* Chan _Type_ _Func_ 
* Chan _Type_ _Func_ _ (if [, error] )

## Pipetubes
* Pipe _Type_ Func  (inp [] _type_ , act func( _type_ ) _type_ ) <-chan _type_
	Note: it 'should' be Pipe _Type_ Map for functional people,
	alas: map has a different meaning in go ...

* Pipe _Type_ Filter  (inp [] _type_ , pass func( _type_ ) bool ) <-chan _type_
* Pipe _Type_ Skiper  (inp [] _type_ , skip func( _type_ ) bool ) <-chan _type_

* Done _Type_ (inp <-chan _type_ [, args]) <-chan struct{}
	an "informative sinkhole" with a simple event channel aka "<-done"-idiom.

For _Func_ in _pack_ with signature func ( _type_ ) ( _type_ [, error]/[, more] )

* Pipe _Type_ _Func_ 
* Pipe _Type_ _Func_ _ (if [, error] )

For _Func_ in _pack_ with signature func ( _type_ ) ( _type-out_ [, error]/[, more] )

* Pipe _Type-out_ _Func_ _Type_ 
* Pipe _Type-out_ _Func_ _Type_ _ (if [, error] )

## further functions

Note: This is work in progress.


### Fanning

* FanIns - variadic Fan-In
* FanInS - Fan-In given slices

* FanOut
* FanOutUpTo (with a workLimit semaphore)

* Merge - sorted Fan-In of sorted and duplicate-free inputs

