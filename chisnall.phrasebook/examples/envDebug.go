package main
import "fmt"
import "os"
import "strconv"

var debugLevel int

func debugLog(level int, msg string, args ...interface{}) {
	if debugLevel > level { fmt.Printf(msg, args...) }
}

func main() {
	debugLevel, _ = strconv.Atoi(os.Getenv("DEBUG"))
	debugLog(3, "Starting\n")
}
