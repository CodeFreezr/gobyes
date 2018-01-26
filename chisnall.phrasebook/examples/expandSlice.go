package main
import "fmt"
func main() {
	s0 := make([]int, 2, 10)
	s1 := append(s0, 2)
	s2 := append(s0, 3)
	fmt.Printf("Element: %d %d\n", s1[2], s2[2])
	s0 = []int{0, 1}
	s1 = append(s0, 2)
	s2 = append(s0, 3)
	fmt.Printf("Element: %d %d\n", s1[2], s2[2])
}
