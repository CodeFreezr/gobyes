import (
    "fmt"
    "os"
)

if err := os.Truncate("filename", newSize); err != nil {
    fmt.Println(err)
}

//\Truncate-a-file\truncate-a-file.go
