package main
import "fmt"
import "path"
import "path/filepath"

func main() {
	components := []string{"a", "path", "..", "with", "relative", "elements"}
	path := path.Join(components...)
	fmt.Printf("Path: %s\n", path)
	decomposed := filepath.SplitList(path)
	for _, dir := range decomposed {
		fmt.Printf("%s%c", dir, filepath.Separator)
	}
	fmt.Printf("\n")
}
