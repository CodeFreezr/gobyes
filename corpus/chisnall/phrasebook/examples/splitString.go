package main
import "strings"
import "unicode"
import "exp/utf8string"
import "fmt"

func main() {
	str := "\tthe important rôles of utf8 text\n"
	str = strings.TrimFunc(str, unicode.IsSpace)
	// The wrong way
	fmt.Printf("%s\n", str[0:len(str)/2])
	// The right way
	u8 := utf8string.NewString(str)
	FirstHalf := u8.Slice(0, u8.RuneCount()/2)
	fmt.Printf("%s\n", FirstHalf)
}
