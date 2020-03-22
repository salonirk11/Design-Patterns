package main

import "fmt"

type CurrentConditionsDisplay struct {
	weatherData Subject
	id int
	temperature float32
	humidity float32
	pressure float32
}

func newItem(weatherData Subject) *CurrentConditionsDisplay {
	return &CurrentConditionsDisplay{
		weatherData: weatherData,
	}
}

func (display *CurrentConditionsDisplay) update(temperature float32, humidity float32, pressure float32) {
	display.temperature = temperature
	display.humidity = humidity
	display.pressure = pressure
	display.display()
}

func (d *CurrentConditionsDisplay) display() {
	fmt.Printf("Temperature: %s\n", d.temperature)
	fmt.Printf("Humidity: %s\n", d.humidity)
	fmt.Printf("Pressure: %s\n", d.pressure)
}

func (display *CurrentConditionsDisplay) getID() int {
	return display.id
}