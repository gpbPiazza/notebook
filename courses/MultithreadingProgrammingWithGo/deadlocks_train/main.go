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

	initTrains()

	s := NewSimulation()
	ebiten.SetWindowSize(s.Layout(0, 0))
	ebiten.SetWindowTitle("Trains simulation")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatal(err)
	}
}

func initTrains() {
	go MoveTrain(trains[0], 300,
		[]*Crossing{
			{Position: 125, Intersection: intersections[0]},
			{Position: 175, Intersection: intersections[1]},
		})

	go MoveTrain(trains[1], 300, []*Crossing{
		{Position: 125, Intersection: intersections[1]},
		{Position: 175, Intersection: intersections[2]},
	})

	go MoveTrain(trains[2], 300, []*Crossing{
		{Position: 125, Intersection: intersections[2]},
		{Position: 175, Intersection: intersections[3]},
	})

	go MoveTrain(trains[3], 300, []*Crossing{
		{Position: 125, Intersection: intersections[3]},
		{Position: 175, Intersection: intersections[0]},
	})
}
