package main

import "fmt"

const carmilage float64 = 15.0

type car struct {
	top_speed_kmh     float64
	gas_volume        float64
	distance_traveled uint16
}

func (c car) can_travel() float64 {
	return c.gas_volume * carmilage
}

func (c *car) update_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	a_car := car{top_speed_kmh: 225.0, gas_volume: 10.5, distance_traveled: 30000}
	fmt.Println("Top Speed = ", a_car.top_speed_kmh)
	fmt.Println("Gas Volume = ", a_car.gas_volume)
	fmt.Println("Distance Traveled = ", a_car.distance_traveled)
	fmt.Println("Can Travel = ", a_car.can_travel())
	a_car.update_top_speed(500.0)
	fmt.Println("Top Speed = ", a_car.top_speed_kmh)
	fmt.Println("Gas Volume = ", a_car.gas_volume)
	fmt.Println("Distance Traveled = ", a_car.distance_traveled)
	fmt.Println("Can Travel = ", a_car.can_travel())
}
