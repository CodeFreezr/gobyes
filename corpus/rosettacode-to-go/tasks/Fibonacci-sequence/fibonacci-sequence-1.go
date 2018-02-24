func fib(a int) int {
  if a < 2 {
    return a
  }
  return fib(a - 1) + fib(a - 2)
}

//\Fibonacci-sequence\fibonacci-sequence-1.go
