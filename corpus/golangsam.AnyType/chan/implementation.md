# Implementation

## functions

Any _pipeable-unit_ is modeled by a suitable function.

(And -mind You- functions are first class citizens in go land.)

Being lazy, we pack and provide a library (and intentionally(!) *not* a framework).
A package with a library of functions => a library of pipable-unit. Batteries included :-)

Being *very* lazy, we (love to) generate go-code for any type.

## _chan_ - go-channels
We use (mostly) nonblocking go-channels to model the inputs and outputs of any pipeable-unit.

They may have various types.

Note: As You know, the strict type-system of go assures for any joined connection to be a proper match and conveniently signals eventual errors at compile time already... a great help.


### directions
Any use of some go-channel may be restricted to a single direction.

#### send-only channel of `SomeType`
Pass type:

	chan<- SomeType

Note: Allows sends but not receives.

#### receive-only channel of `SomeType`
Pass type:

	<-chan SomeType

Note: Allows receives but not sends.

Hint:
* The position of the `<-` arrow relative to the `chan` keyword is a mnemonic.
* Note: Violations of this discipline are detected at compile time.

Note:
* The supplied functions accept and/or supply/return receive-only channels only.
* The sending into the returned channel is launched as part of the function.

* Batteries included :-)
