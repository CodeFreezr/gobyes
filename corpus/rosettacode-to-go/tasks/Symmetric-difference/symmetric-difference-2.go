func main() {
    for e := range b {
        delete(a, e)
    }
    fmt.Println(a)
}

//\Symmetric-difference\symmetric-difference-2.go
