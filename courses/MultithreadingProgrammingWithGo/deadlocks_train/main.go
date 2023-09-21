package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	trains        [4]*Train
	intersections [4]*Intersection
)

const trainLenght = 70

func main() {
	for i := 0; i < 4; i++ {
		trains[i] = NewTrain(i)
		intersections[i] = NewIntersection(i)
	}

	s := NewSimulation()
	ebiten.SetWindowSize(s.Layout(0, 0))
	ebiten.SetWindowTitle("Trains simulation")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatal(err)
	}
}
