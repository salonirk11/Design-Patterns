package main

type Subject interface {
	registerObserver(Observer observer)
	removeObserver(Observer observer)
	notifyObservers()
}

