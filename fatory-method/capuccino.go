package fatory_method

type Cappuccino struct {
	CoffeeDrink
}

func newCappuccino() *Cappuccino {
	return &Cappuccino{
		CoffeeDrink: CoffeeDrink{
			name: "cappuccino",
		},
	}
}
