# Naming conventions

Intentionally, the generated source code provides packages with plenty of visible objects  (exported functions mainly).

Thus, they provide (or extend) a very rich API - and this might be confusing at first sight.

In order to ease navigation in this variety, it may help to understand (and memorize?) a few conventions frequently used in the compositions of their naming.

After some *Meta* about our notations, You'll find typical names chosen used as
- *Prefix*
- *Postfix*
- *Infix*

in composing names such as go identifiers.

Below You'll' also find samples of signatures. Note: For ease of understanding, they are simplified. Hint: You may like to look at the related code sniplets - the `*.tmpl` files in [flavours](flavours.md) [`ssss`](./ssss/) and [`sss`](./sss/) are a good point to start. 

## Meta
- Let *Pack* be the ID of a package such as `strings`, `ioutil` or `os`
- Let *pack* be the import-path for _Pack_ such as `"strings"`, `"io/ioutil"` or `"os"`

- Let *Type* be the Name of a type implemented/supported by _Pack_, such as `String`, `Writer`, `Header`, `File` ... (or none; if it's id is obvious for the given _Pack_)
- Let *type* be the actual type underlying _Type_ such as `string` or `*io.Writer`

- Let *Func* be the ID of an exported function to be wrapped, such as `strings.ToUpper`, `Open`, ...

## Prefixes

- `Make` - a channel creator (convenience)
- `Chan` - a producer, a source
- `Pipe` - a tube, busy working
- `Sink` - intentionally *not* used!
- `Done` - a signalling drainer

## Patterns

### Prefix-patterns
Let `Foo` be a *Type* of *type* `foo`.
Let `Bar` be a *Type* of *type* `bar`.

Note: Any _pipeable-unit_ has (zero or more) input(s) and/or output(s) and may have additional trailing arguments and/or results `[, ...]` - for brevity this is **not** explicitly mentioned below.

---
`Make` - a channel creator (convenience)

	MakeFoo() <-chan foo
	
* Launch: nothing

Note: This is a convenience.

Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
(or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)

---
`Chan` - a producer, a source

	ChanFoo ([args]) <-chan foo

* Launch: a producer

Note: [args] is anything but a read-only chan to be consumed.

---
`Pipe` - a tube, busy working

	PipeFoo (inp <-chan foo) <-chan bar
* Launch: a producer for _Bar_ which consumes _Foo_ and closes `inp`

---
`Sink` - intentionally *not* used!
	SinkFoo (inp <-chan foo)

* Launch: a silent drainer - intentionally *not* provided!

Note: Could also be named *NullXyz*.

---
`Done` - a signalling drainer

	DoneFoo (inp <-chan foo ) <- chan struct{}

* Launch: a drainer, who gives **one** signal on the returned channel when done

---
## Postfixes

- *...`_`* - ignore secondary result(s) such as ok's/errors.
- *...`s`* - Plural via variadic arguments
- *...`S`* - Plural via Slice

### Postfix-patterns

*...`_`* - ignore secondary results such as ok's/errors.

Many functions return multiple values - especially the _"comma-error"_ idiom is very popular.

In order to deal conveniently therewith, we use `_` as postfix for 'silent' implementations which ignore errors (or other secondary data such as bytes written...) and return a chan with the 'main' data only.

Any such quiet/skipping/ignoring implementation shall be complimented with it's fully returning companion. Thus, clients are free to choose.

---
*...`s`* - Plural via variadic arguments

Note: In templated code, only used by generators (Prefix `Chan`).

---
*...`S`* - Plural via Slice

Note: In templated code, only used by generators (Prefix `Chan`).

---
## Infixes
- *...`_`...* ( *`Foo_Bar`* ) - omit optional arguments

### Infix-patterns

...`_`... (`Foo_Bar`) - omit optional arguments
`_` is used as Infix when secondary arguments are ignored/not required,
and reasonable and documented(!) defaults are subsituted.

Note: A twin without `_` should behave exactly similar, if called with nil argument(s).

Note: Yes, this is syntactic sugar.
Yes, this is useful and practical (and thus almost mandatory)
as long as go does not allow optional arguments.
