package main
import "fmt"
import "os"

func main() {
	file, err := os.Open("fileRead.go")
	if err != nil{
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	buffer := make([]byte, 100)
	for n, e := file.Read(buffer) ; e == nil ; n, e = file.Read(buffer) {
		if n > 0 {
			os.Stdout.Write(buffer[0:n])
		}
	}
}
