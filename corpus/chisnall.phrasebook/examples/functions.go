package main
import "fmt"

func printf(str string, args ...interface{}) (int, error) {
	_, err := fmt.Printf(str, args...)
	return len(args), err
}

func main() {
	count := 1
	closure := func(msg string) {
		printf("%d %s\n", count, msg)
		count++
	}
	closure("A Message")
	closure("Another Message")
}
