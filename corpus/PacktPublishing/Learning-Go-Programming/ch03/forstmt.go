package main

import (
	"fmt"
)

type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}

var currencies = []Curr{
	Curr{"NOK", "Norwegian Krone", "Norwary", 578},
	Curr{"KES", "Kenyan Shilling", "Kenya", 404},
	Curr{"DZD", "Algerian Dinar", "Algeria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
	Curr{"MXN", "Mexican Peso", "Mexico", 484},
	Curr{"EUR", "Euro", "Greece", 978},
	Curr{"KHR", "Riel", "Cambodia", 116},
	Curr{"SZL", "Lilangeni", "Swaziland", 748},
	Curr{"GBP", "Pound Sterling", "Isle of Man", 826},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"BWP", "Pula", "Botswana", 72},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
	Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"TRY", "Turkish Lira", "Turkey", 949},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"JMD", "Jamaican Dollar", "Jamaica", 388},
	Curr{"ALL", "Lek", "Albania", 8},
	Curr{"GEL", "Lari", "Georgia", 981},
	Curr{"KFM", "Coromo Franc", "Comoros", 174},
	Curr{"NZD", "New Zeland Dollar", "Tokelau", 554},
}

var sortedCurrs []Curr

func main() {
	fmt.Println("Currencies")
	fmt.Println("----------")
	listCurrs()

	fmt.Println("----------------")
	fmt.Println("Sorted by Number")
	fmt.Println("----------------")
	sortByNumber()
	listCurrs()

	fmt.Println("------------------")
	fmt.Println("Sorted by Currency")
	fmt.Println("------------------")
	sortByCurrency()
	listCurrs()
}

func listCurrs() {
	i := 0
	for i < len(currencies) {
		fmt.Println(currencies[i])
		i++
	}
}

func sortByNumber() {
	N := len(currencies)
	for i := 0; i < N-1; i++ {
		currMin := i
		for k := i + 1; k < N; k++ {
			if currencies[k].Number < currencies[currMin].Number {
				currMin = k
			}
		}
		// swap
		if currMin != i {
			temp := currencies[i]
			currencies[i] = currencies[currMin]
			currencies[currMin] = temp
		}
	}
}

func sortByCurrency() {
	N := len(currencies)
	for i := 0; i < N-1; i++ {
		currMin := i
		for k := i + 1; k < N; k++ {
			if currencies[k].Currency < currencies[currMin].Currency {
				currMin = k
			}
		}
		// swap
		if currMin != i {
			temp := currencies[i]
			currencies[i] = currencies[currMin]
			currencies[currMin] = temp
		}
	}
}
