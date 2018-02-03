# Gain concurrency - use go-channels for piping (and more)

*[Piping](pipe.md)* is a popular concept, method, facility, mechanism ... whatever You call it.

- **And**: piping is a convenient way to design a process - a process with **concurrent** parts.
  Batteries included :-)

- **And**: Mind You: *[Concurrency enables parallelism!](http://go-proverbs.github.io/)*

[This project `chan`]([here](https://github.com/GoLangsam/AnyType/tree/master/chan/) provides facilites and mechanisms to build, invoke and manage such process.

Each such concurrent part (and any composite hereof(!)) may be seen as a _pipeable-unit_.

Note: Some name it 'transitor' (not transi`s`tor), as it facilitates stuff to transit thru it. Or 'pipe-station', or 'pipe-tube'). [Some](https://blog.golang.org/pipelines) call it 'stage'.

[They](functions.md) come in a variety of [flavours](flavours.md), have different [sizes](sizes.md), and obey to strict and consistent [namings](namings.md).

---
Connect several such pipeable-unit into a **networked** ensemble, and You desing, build, create and operate a process.

**And**: Such process is **concurrent** by design!

Thus: it's components may execute as parallel(!)
(as much as Your environemt permits/supports).

**Really** parallel (on multi-cores), that is.
Not only quasi-parallel (via some pre-emptive multitasking).

Note: Mind You: Conceptually, *Piping* is a mechanism as much as *channeling*.  
`chan` is not another type, but an atom of such mechanism (think: CSP)  
a means to some end.

---
[Available literature](resources.md) about piping and concurrency is inconsistent and/or incomplete in terms of nomenclatura (so far we've read *both* good books available to us).

Thus, we take liberty to introduce some freshly invented [vocabulary](Vocabulary.md), complementing the [namings](namings.md).

---
Hint:
- If You like it: ***smile***.
- If You don't: ***smile anyway*** - You'll be (a little more) happy - may be.

Mind You: the work done here is for You!. `ʕ◔ϖ◔ʔ`  
For You - to be a ***Happy*** Gopher! `ʕ◔ϖ◔ʔ`

So: be a ***Happy*** Gopher! `ʕ◔ϖ◔ʔ`
