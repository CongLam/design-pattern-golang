package fatory_method

import "fmt"

const (
	latteName      = "Latte"
	cappuccinoName = "Cappuccino"
)

func GetCoffeeDrink(name string) (ICoffeeDrink, error) {
	switch name {
	case latteName:
		return newLatte(), nil
	case cappuccinoName:
		return newCappuccino(), nil
	default:
		return nil, fmt.Errorf("coffee %s not found", name)
	}
}
