package main

type Observer interface {
	update(float32, float32, float32)
	getID() int
}
