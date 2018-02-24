func Y(f FuncFunc) Func {
	return func(x int) int {
		return f(Y(f))(x)
	}
}

//\Y-combinator\y-combinator-2.go
