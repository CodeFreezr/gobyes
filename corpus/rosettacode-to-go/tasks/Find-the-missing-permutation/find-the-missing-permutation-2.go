func main() {
    b := make([]byte, len(given[0]))
    for _, p := range given {
        for i, c := range []byte(p) {
            b[i] ^= c
        }
    }
    fmt.Println(string(b))
}

//\Find-the-missing-permutation\find-the-missing-permutation-2.go
