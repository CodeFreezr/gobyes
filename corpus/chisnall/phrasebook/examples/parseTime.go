package main
import "fmt"
import "time"

func main() {
	var t string
	fmt.Printf("Enter a time\n")
	fmt.Scanf("%s", &t)
	parsed, err := time.Parse("03:04PM", t)
	if err != nil {
		parsed, err = time.Parse("15:04", t)
	}
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Time in seconds since the Epoc: %d\n", parsed.Unix())
	}
	fmt.Printf("%s\n", parsed.Format(time.RFC3339))
}
