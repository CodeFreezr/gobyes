# Shapes

The shape/form of a _pipeable-unit_ can vary:

### one-sided - unary
- **pumpfeed** (or *pumpfill*):
	- has only output, no input

- **sinkhole**
	* has only input, no output

### two-sided - binary
- **extender**
	- one input & one output
	- both of same type

- **converter**
	- one input & one output
	- of different type

- **adapter** (or *caster*)
	- one input & one output
	- of different type
		- connected by nothing but a type-conversion/extraction/injection

### fan-shaped - unary/n-ary
- **fan-in** (= n-ary/unary)
	- several inputs & one output
	- all of same type

- **fan-out** (= unary/n-ary)
	- one input & several outputs
	- all of same type

Warning: some stupid & distracting question/bad joke/pun is ahead - please feel free to skip!
- A **fan** is also called a *gopher*, is it not? Okey - when it's a fan of go, at least ... , is it not? Yawn ... )

### double-fan - gopher-shaped
- **hourglass**
	- several input & several output
	- all of same type

- **bottleneck**
	- several input & several output
	- each side of different type

* **gopher-task** :-)
	* several input & several output
	* all of different type

## afterthoughts
Note:
* Being short of phantasy, we do not (yet) provide any (atomic) gopher-task. ;-)
* Being lazy, we do not (yet) provide any (atomic) hourglass or bottleneck.
* Hint: Combine **fan-in** & **fan-out** for an hourglass, or stretch with some **converter** for a bottleneck.

* Note: Please feel free to pass enlightening suggestions/remarks/comments on to us.

---
### No sinkholes, but Whips
* Intentionally no 'real' **sinkhole** is supplied, as such would give more pain than gain. ( and should better be named 'wormhole', as it's use would open a can of bugs - I mean: worms. )

* We use the "`done`-idiom" and supply respective `DoneXyz`-functions, which return a single-event-channel (sometimes also called **Whip**) and signal completion hereon.
