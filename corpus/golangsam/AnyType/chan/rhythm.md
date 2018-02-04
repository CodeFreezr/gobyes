# Rhythm
The rhythm of a pipe can vary:
one input does not need to provide one output (immediately).

Thus, some strategy has to be applied, such as

## counting
- diverse: one input results in many outputs
- compact: many inputs result in one output
	- What, if not enough inputs arrive before close?

## dragging
- non-blocking consume input, with or without retry(timeout, maxRetries)

## draining
or *flushing*, or *buffering*, or *grouping*
- the pipe buffers until it becomes triggered due to
	- a special condition which can be determined internally
	- a special input arrives such as a separator or a starter, or
	- a special signal arrives as *HeartBeat* (or *MarchBeat* *GoBeat* *Pop* *Pat* ...)
		- typically from the input process
		- thus, they are two-two, with `SomeType` & `struct{}{}` (= Whip)
		- which could be named as a `SomeTypeWhip` or `SomeTypeTrigger`

## grouping
A grouping tube may have and use(!) functions
- `Prolog` - Gruppenvorlauf
- `Onelog` - element in Gruppe registrieren - für Sub(Total) u.ä. 
- `Epilog` - Gruppennachlauf
and objects Header & Footer to be rendered

Further, a multi-input grouping tube such as a *zipper* has either to:
- drop (and drain remaining inputs)
- fill missing inputs somehow

and has to answer the question
- What, if not enough inputs arrive before close of some?
	- panic?
	- quiet?
	- shout? to whom?
	- log? to where?

It ain't easy - it's an art, as are music and dance.