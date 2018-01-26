package main
import "errors"
import "fmt"
import "math"

func sqrt(i int) (result float64, error error) {
	if i < 0 {
		return 0, errors.New("Invalid argument")
	}
	return math.Sqrt(float64(i)), nil
}

func main() {
	// Ignoring error value, because 2 is a valid input
	r, _ := sqrt(2)
	fmt.Printf("sqrt(2) = %f\nEnter another number\n", r)
	var i int
	fmt.Scanf("%d", &i)
	root, err := sqrt(i)
	if err == nil {
		fmt.Printf("sqrt(%d) = %f\n", i, root)
	} else {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
