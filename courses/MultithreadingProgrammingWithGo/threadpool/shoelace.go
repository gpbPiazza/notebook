package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Point2D struct {
	x int
	y int
}

var (
	r = regexp.MustCompile(`\((\d*),(\d*)\)`)
)

func calcArea(pointsStr string) float64 {
	var points []Point2D

	for _, point := range r.FindAllStringSubmatch(pointsStr, -1) {
		x, _ := strconv.Atoi(point[1])
		y, _ := strconv.Atoi(point[2])
		points = append(points, Point2D{x: x, y: y})
	}

	var area float64
	for i := 0; i < len(points); i++ {
		neverWillBetOutOfRange := (i + 1) % len(points)
		a, b := points[i], points[neverWillBetOutOfRange]
		area += float64(a.x*b.y) - float64(a.y*b.x)
	}

	area = math.Abs(area) / 2
	return area
}

func main() {
	line := "(4,10),(12,8),(10,3),(2,2),(7,5)"
	area := calcArea(line)
	fmt.Println("Polygon area:", area)
}
