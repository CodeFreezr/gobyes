package main
import "fmt"
import "regexp"

var static = regexp.MustCompile(", *")

func main() {
	r, _ := regexp.Compile("abcd*")
	str := "abcddd fish, wibble abcd, abc, foo"
	fmt.Printf("Replaced: %v\n", r.ReplaceAllString(str, "bar"))
	fmt.Printf("Replaced: %v\n", static.ReplaceAllString(str, ".  "))
}
