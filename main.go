package main

import "time"

func main() {

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
