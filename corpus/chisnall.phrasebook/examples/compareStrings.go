package main
import "fmt"

func main() {
	str1 := "A string"
	str2 := "A " + "string"
	if (str1 == str2) {
		fmt.Printf("'%s' and '%s' are equal\n", str1, str2)
	}
	if (&str1 == &str2) {
		fmt.Printf("'%s' and '%s' are identical\n", str1, str2)
	}
	str2 += " with a suffix"
	if (str1 < str2) {
		fmt.Printf("'%s' comes before '%s'\n", str1, str2)
	}

}
