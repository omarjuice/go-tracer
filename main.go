package main

import (
	"fmt"
	"sync"
	"time"
)

func clock() {
	colors := [3]*Color{NewColor(1, 0, 0), NewColor(0, 1, 0), NewColor(0, 0, 1)}

	i := 0

	for {
		color := colors[i]
		i++
		i %= 3
		canvas := NewCanvas(400, 400)

		canvas.SetOrigin(canvas.width/2, canvas.height/2)

		origin := Point(0, 0, 0)

		canvas.Write(origin, NewColor(1, 1, 1))

		current := origin.Transform(Translation(0, 150, 0))

		for i := 0; i < 12; i++ {
			current = current.Transform(RotationZ(π / 6))
			canvas.Write(current, color)
		}

		canvas.ToPPM("clock")
		time.Sleep(time.Second)
	}
}

func main() {

}

func sphereCast() {
	start := time.Now()
	canvas := NewCanvas(200, 200)

	rayOrigin := Point(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	pixelSize := wallSize / float64(canvas.width)

	half := wallSize / 2
	color := NewColor(1, 0, 0)
	shape := NewSphere()
	var wg sync.WaitGroup
	for y := 0; y < (canvas.height); y++ {
		wg.Add(1)
		go func(y int) {
			worldY := half - pixelSize*float64(y)
			for x := 0; x < (canvas.width); x++ {
				worldX := -half + pixelSize*float64(x)

				position := Point(worldX, worldY, wallZ)
				r := NewRay(rayOrigin, position.Sub(rayOrigin).Normalize())

				xs := r.Intersect(shape)

				if xs.Hit() != nil {
					canvas.WritePixel(x, y, color)
				}

			}
			wg.Done()
		}(y)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))

	canvas.ToPPM("circle")
}
