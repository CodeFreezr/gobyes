This project [`pile`](https://github.com/GoLangsam/AnyType/tree/master/pile/) provides -well- a Pile.

A concurrency-safe growing-only first-in-first-out (FIFO) cached collection of things - with a simple (single-thread) iterator

Just a simple, lazy and (hopefully) useful container: A **type-safe** Pile of Your favourite type-of-things :-)

## Pile
A pile -as used here- means just:
- a collection
  - of things (of Your favourite type)

Note: It could be [*anything*](https://github.com/GoLangsam/AnyType/blob/master/pile/basic/interface/IsAny/Pile.dot.go) (= `interface{}`) - but [mind You](http://go-proverbs.github.io/): `interface{}` says nothing

with characteristics such as:
- one-way
  - `Next()` returns items in the order they were `Pile()`ed
- FIFO (first-in first-out)
  - `Next()` returns items in the order they were `Pile()`ed
- grow-only
  - it can grow (until closed), but neither shrink nor change. (Thus, it's neither a *stack* nor a *queue*.)
- concurrency safe growth
  - this comes *for free*, as a channel is used under the hood
- lazy
  - it grows as fast (or slow) as You traverse it (well, You may `buff` a growth ahead)
- mindful
  - `Pile()` it, and it shall be remembered (cached) until You discard the entire Pile
- reusable
  - `Reset()` to begin a fresh traversal

## History
I was thinking about something like this for quite a while.

Discovery of things may be slow and/or expensive. (Think of e.g. interesting names of files or directories, or stuff You grab from elsewhere on the net).

I came up with the idea of a `Pile` in order to separate such discovery of things from their use.

## Credit
For the implementation I got inspired by parts of `"cmd/doc/dirs.go"`, one of these little-known beauties - well hidden inside the standard package.

---
Hint:
- If You like it: ***smile***.
- If You don't: ***smile anyway*** - You'll be (a little more) happy - may be.

Mind You: the work done here is for You!. `ʕ◔ϖ◔ʔ`  
For You - to be a ***Happy*** Gopher! `ʕ◔ϖ◔ʔ`

So: be a ***Happy*** Gopher! `ʕ◔ϖ◔ʔ`
