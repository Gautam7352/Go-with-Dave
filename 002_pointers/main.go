package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func (p *Product) UpdatePrice(Prc float64) {
	p.Price = Prc
}

func main() {
	product1 := &Product{Name: "Laptop", Price: 999.99}
	prod := Product{Name: "Smartphone", Price: 499.99}
	product1.UpdatePrice(399.99)
	prod.UpdatePrice(899.99)

	fmt.Println("Product 1:", product1.Name, "Price:", product1.Price)
	fmt.Println("Product 2:", prod.Name, "Price:", prod.Price)
}

/*
Define a type called car field[Brand, name, engine_type, max speed, running_speed]

Implement a method which gives description of car
Implement a method to run the car
Implement a method to speed up
Implement a method to slow down
Implement a method to stop the car

Don't use function
*/
