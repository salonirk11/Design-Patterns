package main

type WeatherData struct {
	observerList []Observer
	temperature float32
	humidity float32
	pressure float32
}

func (data *WeatherData) registerObserver(o Observer) {
	data.observerList = append(data.observerList, o)
}

func (data *WeatherData) removeObserver(o Observer) {
	data.observerList = removeFromslice(data.observerList, o)
}

func (data *WeatherData) notifyObservers() {
	for _, observer := range data.observerList {
		observer.update(data.temperature, data.humidity, data.pressure)
	}
}

func (data *WeatherData) measurementsChanged() {
	data.notifyObservers()
}

func (data *WeatherData) setMeasurements(temperature float32, humidity float32, pressure float32) {
	data.temperature = temperature
	data.humidity = humidity
	data.pressure = pressure
	data.measurementsChanged()
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}


