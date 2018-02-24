    // allocate composed 2d array
    a := make([][]int, row)
    e := make([]int, row * col)
    for i := range a {
        a[i] = e[i*col:(i+1)*col]
    }

//\Create-a-two-dimensional-array-at-runtime\create-a-two-dimensional-array-at-runtime-2.go
