package main

import "fmt"

type User struct {
	Name string
}

type Person struct {
	Name string
}

type value int

func (v value) String() string {
	return fmt.Sprintf("Value: %d", v)
}

func (u User) UserGreet() {
	fmt.Println(u.Name + " says, Hello!")
}

func (p1 Person) PersonGreet() {
	fmt.Println(p1.Name + " says, Hi!")
}

type Product struct {
	Name  string
	Price float64
}

type Vehicle struct {
	Brand string
	Model string
	Price float64
	Speed int
}

func (v *Vehicle) Drive(speed int) {
	v.Speed = speed
	fmt.Printf("Vehicle is running at %d km/h\n", speed)
}

func (v *Vehicle) IncreaseSpeed(increaseBy int) {
	v.Speed = v.Speed + increaseBy
}

func (p *Product) UpdatePrice(newPrice float64) {
	p.Price = newPrice
}

func main() {
	user1 := User{Name: "Alice"}
	user1.UserGreet()

	person1 := Person{Name: "Bob"}
	person1.PersonGreet()

	name := "Charlie"
	// Print memory address of the variable
	fmt.Printf("Memory address of name: %p\n", &name)
	SayHi(name)

	product1 := &Product{Name: "Laptop", Price: 999.99}
	product1.UpdatePrice(899.99)
	fmt.Printf("Updated price of %s: %.2f\n", product1.Name, product1.Price)
}

func SayHi(name string) {
	// Print memory address of the parameter
	fmt.Printf("Memory address of parameter name: %p\n", &name)
	fmt.Println(name + " says, Hi!")
}
