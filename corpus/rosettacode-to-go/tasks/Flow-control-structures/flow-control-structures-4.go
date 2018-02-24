func answer(phone1, phone2 chan int) {
    select {
    case <-phone1:
        // talk on phone one
    case <-phone2:
        // talk on phone two
    }
}

//\Flow-control-structures\flow-control-structures-4.go
