package main

//Interface in golang
type MotorVehicle interface {
	Mileage() float64
}

type BMW struct {
	Distance    float64
	Fuel        float64
	Aeragespeed string
}

type Audi struct {
	Distance float64
	Fuel     float64
}
