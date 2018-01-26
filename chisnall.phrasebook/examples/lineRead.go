package main
import "fmt"
import "os"
import "bufio"

func main() {
	file, err := os.Open("lineRead.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	lineNumber := 1
	lineReader := bufio.NewReaderSize(file, 20)
	for line, isPrefix, e := lineReader.ReadLine() ; e==nil ;
		line, isPrefix, e = lineReader.ReadLine() {
		fmt.Printf("%.3d: ", lineNumber)
		lineNumber++
		os.Stdout.Write(line)
		if isPrefix {
			for {
				line, isPrefix, _ = lineReader.ReadLine()
				os.Stdout.Write(line)
				if !isPrefix { break }
			}
		}
		fmt.Printf("\n")
	}
}
