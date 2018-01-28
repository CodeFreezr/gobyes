package main
import "fmt"

func main() {
	var s string
	var flag bool
	go func() {
		fmt.Scanf("%s", &s)
		flag = true
	}()
	for !flag {}
	fmt.Printf("%s\n", s)
}
