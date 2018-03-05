package main

import (
	"fmt"
	gc "github.com/go-goodies/go_currency"
)

func main() {
	input := make(chan Order)
	output := make(chan Order)

	go func() {
		for order := range input {
			output <- Pipeline(order)
		}
	}()

	orders := GetOrders()
	for _, order := range orders {
		fmt.Printf("Processed order: %v\n", Pipeline(*order))
	}
	close(input)
}

// close the input channel so start() will exit and can clean up after
// itself if it so wishes.

func Pipeline(o Order) Order {
	o = Authenticate(o)
	o = Decrypt(o)
	o = Charge(o)
	return o
}


func Authenticate(o Order) Order  {
	fmt.Printf("Order %d is Authenticated\n", o.OrderNumber)
	return o
}

func Decrypt(o Order) Order {
	fmt.Printf("Order %d is Decrypted\n", o.OrderNumber)
	return o
}

func Charge(o Order) Order {
	fmt.Printf("Order %d is Charged\n", o.OrderNumber)
	return o
}



type Order struct {
	OrderNumber int
	IsValid bool
	Credentials string
	CCardNumber string
	CCardExpDate string
	LineItems []LineItem
}

type LineItem struct {
	Description string
	Count       int
	PriceUSD    gc.USD
}
func GetOrders() []*Order {

	order1 := &Order{
		10001,
		true,
		"alice,secret",
		"7b/HWvtIB9a16AYk+Yv6WWwer3GFbxpjoR+GO9iHIYY=",
		"0922",
		[]LineItem{
			LineItem{"Apples", 1, gc.USD{4, 50}},
			LineItem{"Oranges", 4, gc.USD{12, 00}},
		},
	}

	order2 := &Order{
		10002,
		true,
		"bob,secret",
		"EOc3kF/OmxY+dRCaYRrey8h24QoGzVU0/T2QKVCHb1Q=",
		"0123",
		[]LineItem{
			LineItem{"Milk", 2, gc.USD{8, 00}},
			LineItem{"Sugar", 1, gc.USD{2, 25}},
			LineItem{"Salt", 3, gc.USD{3, 75}},
		},
	}
	orders := []*Order{order1, order2}
	return orders
}
