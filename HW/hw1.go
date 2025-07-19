/*
Define a type called car field[Brand, name, engine_type, running_status, running_speed]

Implement a method which gives description of car
Implement a method to run the car
Implement a method to speed up
Implement a method to slow down
Implement a method to stop the car

Don't use function
*/

package main

import "fmt"

type car struct {
	Brand          string
	Name           string
	engine_type    string
	running_status bool
	running_speed  int
}

//Implement a method which gives description of car
func (c *car) description() {
	fmt.Println("Car Brand:", c.Brand)
	fmt.Println("Car Name:", c.Name)
	fmt.Println("Car engine type:", c.engine_type)
	fmt.Println("Car running status:", c.running_status)
	fmt.Println("Car running speed:", c.running_speed)

}

//Implement a method to run the car
func (c *car) start() {
	c.accelerate(0)
}

//Implement a method to speed up
func (c *car) accelerate(spd int) {
	c.running_speed = spd + 10
}

//Implement a method to slow down
func (c *car) decelerate(spd int) {
	c.running_speed = spd - 10
}

//Implement a method to stop the car
func (c *car) stop() {
	c.running_speed = 0
}

func main() {
	Car1 := car{Brand: "Mercedes", Name: "Benz", engine_type: "Petrol", running_status: false, running_speed: 0}

	Car1.description()

	fmt.Println("Car speed before start:", Car1.running_speed)
	Car1.start()
	fmt.Println("Car speed after start:", Car1.running_speed)
	Car1.accelerate(Car1.running_speed)
	fmt.Println("Car speed after acceleration:", Car1.running_speed)
	Car1.accelerate(Car1.running_speed)
	fmt.Println("Car speed after acceleration:", Car1.running_speed)
	Car1.decelerate(Car1.running_speed)
	fmt.Println("Car speed after deceleration:", Car1.running_speed)
	Car1.stop()
	fmt.Println("Car speed after stop:", Car1.running_speed)
}
