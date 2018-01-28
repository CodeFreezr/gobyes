package main
import "fmt"
import "os"

func main() {
	fmt.Printf("Enter a file name\n")
	var s string
	fmt.Scanf("%s", &s)
	fi, err := os.Stat(s)
	if err != err {
		fmt.Printf("%s does not exist!\n", s)
		return
	}
	if fi.IsDir() {
		fmt.Printf("%s is a directory\n", s)
	}
	mode := fi.Mode()
	if mode & os.ModeSymlink == os.ModeSymlink {
		fmt.Printf("%s is a symbolic link\n", s)
	}
}
