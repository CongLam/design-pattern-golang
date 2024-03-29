package main

import (
	"fmt"

	abstract_factory "qna-manager/abstract-factory"
	"qna-manager/builder"
	fatory_method "qna-manager/fatory-method"
	"qna-manager/prototype"
	counter "qna-manager/singleton"
)

func main() {
	/*
		Example Singleton
	*/
	fmt.Println("*** Example Singleton ***")
	// User does: Print, then increase counter...
	// At this time, we have no Counter instance yet
	// So, we create one. With the value of 0
	counter.GetInstance().Increase()

	// User print again..., also increase counter...
	counter.GetInstance().Increase()

	fmt.Println("Print count: ", counter.GetInstance().GetCount())
	fmt.Print("*** End of Singleton ***\n\n\n")

	/*
		Example Builder
	*/
	fmt.Println("*** Example Builder ***")
	manufacturingVehicle := builder.ManufacturingDirector{}

	bicycleBuilder := &builder.BicycleBuilder{}
	// motorbikeBuilder := &builder.MotorbikeBuilder{}

	manufacturingVehicle.SetBuilder(bicycleBuilder)
	manufacturingVehicle.Construct()

	bicycle := bicycleBuilder.GetVehicle()
	fmt.Printf("Vehicle is %s has %d wheels, %d seats.\n", bicycle.Structure, bicycle.Wheels, bicycle.Seats)
	fmt.Print("*** End of Builder ***\n\n\n")

	/*
		Example Factory Method
	*/
	fmt.Println("*** Example Factory Method ***")
	cappuccino, err := fatory_method.GetCoffeeDrink("Cappuccino")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cappuccino.GetName())
	}

	latte, err := fatory_method.GetCoffeeDrink("Latte")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(latte.GetName())
	}

	_, err = fatory_method.GetCoffeeDrink("Milk")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("*** End of Factory Method ***\n\n\n")

	/*
		Example Abstract Factory
	*/
	fmt.Println("*** Example Abstract Factory ***")
	bicycleFactory, err := abstract_factory.BuildFactory(abstract_factory.BicycleFactoryType)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	sportsBicycle, err := bicycleFactory.NewVehicle(abstract_factory.SportBicycleType)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Sport Bicycle:")
	fmt.Printf("Vehicle has %d wheels, %d seats.\n", sportsBicycle.NumWheels(), sportsBicycle.NumSeats())

	fmt.Print("*** End of Abstract Factory ***\n\n\n")

	/*
	   Example Prototype
	*/
	fmt.Println("*** Example Prototype ***")
	homePage := prototype.NewPage("/home", prototype.MAIN_LAYOUT)
	homePage.SetBody("Home Body")
	fmt.Println(homePage.GetInfo())

	profilePage := prototype.NewPage("/profile", prototype.MAIN_LAYOUT)
	fmt.Println(profilePage.GetInfo())
	profilePage.SetBody("Profile Body")
	fmt.Println(profilePage.GetInfo())

	loginPage := prototype.NewPage("/login", prototype.BLANK_LAYOUT)
	fmt.Println(loginPage.GetInfo())
	fmt.Print("*** End of Prototype ***\n\n\n")
}
