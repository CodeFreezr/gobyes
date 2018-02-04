package main 

func main() {
	var i int
	var Θ float32
	var explicitly, typed, pointers *complex128
	int_pointer := &i
	another_int_pointer := new(int)
	generic_channel := make(chan interface{})
	// Some random uses to make this file compile
	// This won't actually do anything, don't try running this example!
	int_pointer = another_int_pointer
	another_int_pointer = int_pointer
	i++
	Θ += 12
	go func() { <- generic_channel }()
	generic_channel <- 12
	explicitly = typed
	explicitly = pointers
	pointers = explicitly
	foo := new(chan interface{})
	go func() { <- *foo}()
	*foo <- 12
}
