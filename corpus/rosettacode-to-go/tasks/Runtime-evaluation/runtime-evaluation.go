package main

import (
	"bitbucket.org/binet/go-eval/pkg/eval"
	"fmt"
	"go/token"
)

func main() {
	w := eval.NewWorld()
	fset := token.NewFileSet()

	code, err := w.Compile(fset, "1 + 2")
	if err != nil {
		fmt.Println("Compile error")
		return
	}

	val, err := code.Run()
	if err != nil {
		fmt.Println("Run time error")
		return
	}
	fmt.Println("Return value:", val) //prints, well, 3

}

//\Runtime-evaluation\runtime-evaluation.go
