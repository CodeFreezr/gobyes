package main
import "fmt"

func main() {
	str := "Étoilé"
	bytes := str[0:7]
	str2 := string(bytes)
	for i, c := range str2 {
		if (0xFFFD == c) {
			str2 = str2[i:]
			break
		} else {
		fmt.Printf("%c", c)
		}
	}
	fmt.Printf("\n")
}
