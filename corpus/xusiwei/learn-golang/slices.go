package main

import "fmt"

func main() {

	s := make([]string, 3) // create a slice using `make`
	fmt.Println("init, s:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("after set, s:", s)
	fmt.Println("s[2]: ", s[2])

	fmt.Println("len(s):", len(s))
	fmt.Println("cap(s):", cap(s))

	// slice size not fixed, unlike array
	// it's no specified length.
	s = append(s, "d")
	s = append(s, "e")
	fmt.Println("after append, s:", s)
	fmt.Println("len(s):", len(s))
	fmt.Println("cap(s):", cap(s))


	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c:", c)


	// slice operatoration [:]
	slice := s[2:5]
	fmt.Println("s[2:5]:", slice)
	fmt.Println("s[2:]:", s[2:])
	fmt.Println("s[:5]:", s[:5])
	fmt.Println("s[:]:", s[:])
//	fmt.Println("s[:-1]:", s[:-1]) // don't support


	t := []string {"s", "l", "i", "c", "e"}
	fmt.Println("t:", t)
	fmt.Println("t[:]:", t[:])

	a := [5]string {"a", "r", "r", "a", "y"}
	fmt.Println("a:", a)
	fmt.Println("a[:]:", a[:])

	sa := a[2:4]
	sa[0] = "R"
	sa[1] = "A"
	fmt.Println("sa:", sa)
	fmt.Println("a:", a)


	// slice can be multi-dimensional:
	//  the length of inner slices can vary,
	//  unlike multi-dimensional array.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

// refers: http://blog.golang.org/go-slices-usage-and-internals
