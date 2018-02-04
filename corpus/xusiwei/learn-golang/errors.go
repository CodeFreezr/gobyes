package main

import "fmt"
import "errors"

// In Go it’s idiomatic to communicate errors via an explicit, separate return value.
// This contrasts with the exceptions used in languages like Java and Ruby 
// and the overloaded single result / error value sometimes used in C.
//
// Go’s approach makes it easy to see which functions return errors
// and to handle them using the same language constructs employed for any other, non-error tasks.


// By convention, errors are the last return value and have type error, a built-in interface.
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d-%s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int {17, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int {17, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	// type assertion, https://golang.org/ref/spec#Type_assertions
	// form: x.(T), asserts that x is not nil and that the value stored in x is of type T.
	// If the type assertion holds, the value of the expression is the value stored in x and its type is T.
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

// read more: http://blog.golang.org/error-handling-and-go
