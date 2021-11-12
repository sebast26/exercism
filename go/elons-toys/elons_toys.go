package elon

import "fmt"

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

// Track implements a race track.
type Track struct {
	distance int
}

// CreateCar creates a new car with given specifications.
func CreateCar(speed, batteryDrain int) *Car {
	return &Car{speed: speed, batteryDrain: batteryDrain, battery: 100}
}

// CreateTrack creates a new track with given distance.
func CreateTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time.
func (car *Car) Drive() {
	if car.battery-car.batteryDrain <= 0 {
		return
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
}

// CanFinish checks if a car is able to finish a certain track.
func (car *Car) CanFinish(track Track) bool {
	currentBatteryLevel := car.battery
	currentDistance := car.distance
	for currentBatteryLevel >= 0 {
		currentBatteryLevel -= car.batteryDrain
		currentDistance += car.speed
	}
	return currentDistance >= track.distance
}

// DisplayDistance displays the distance the car is driven.
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery displays the battery level.
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}
