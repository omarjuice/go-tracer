package main

import (
	"fmt"
	"sync"
	"time"
)

const defaultRecursionDepth = 4

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
	threeSpheres(1000, 500)
}

func circleCast() {
	start := time.Now()
	canvas := NewCanvas(100, 100)

	rayOrigin := Point(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	pixelSize := wallSize / float64(canvas.width)

	half := wallSize / 2
	color := NewColor(0, 1, 0)
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

func shinySphere() {
	start := time.Now()
	canvas := NewCanvas(500, 500)

	rayOrigin := Point(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	pixelSize := wallSize / float64(canvas.width)

	half := wallSize / 2

	shape := NewSphere()
	shape.material = DefaultMaterial()
	shape.material.color = NewColor(0, 1, 0)
	shape.SetTransform(Scaling(1, 1, 1))
	light := NewPointLight(Point(-10, 10, -10), NewColor(1.0, 1.0, 1.0))
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
				hit := xs.Hit()
				if hit != nil {

					point := r.Position(hit.t)
					normal := hit.object.NormalAt(point)
					eye := r.direction.Negate()

					color := Lighting(hit.object.Material(), hit.object, light, point, eye, normal, false)

					canvas.WritePixel(x, y, color)
				}

			}
			wg.Done()
		}(y)
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start))

	canvas.ToPPM("sphere")
}
func threeSpheres(width, height int) {

	start := time.Now()

	floor := NewPlane()
	floor.SetTransform(Scaling(10, 0.01, 10))
	floor.material.pattern = CheckersPattern(White, Black)
	floor.material.pattern.SetTransform((Scaling(.1, .1, .1)))
	floor.material.color = NewColor(1, 1, 1)
	floor.material.specular = 0.3
	floor.material.diffuse = 0.5
	floor.material.shininess = 200
	// floor.material.reflective = 0.5

	leftWall := NewPlane()
	leftWall.SetTransform(Translation(0, 0, 5).MulMatrix(RotationY(-π / 4)).MulMatrix(RotationX(π / 2)).MulMatrix(Scaling(10, 0.01, 10)))
	leftWall.material = floor.material

	rightWall := NewPlane()
	rightWall.SetTransform(Translation(0, 0, 5).MulMatrix(RotationY(π / 4)).MulMatrix(RotationX(π / 2)).MulMatrix(Scaling(10, 0.01, 10)))
	rightWall.material = floor.material

	backWall := NewPlane()
	backWall.SetTransform(Translation(0, 0, -6).MulMatrix(RotationX(π / 2)))
	backWall.material = floor.material

	middle := NewSphere()
	middle.SetTransform(Translation(-0.5, 1, 0.5))
	// middle.material.pattern = GradientPattern(NewColor(1, 0, 0), NewColor(0, 0, 1))
	// middle.material.pattern.SetTransform(Scaling(0.5, 0.5, 0.5))
	middle.material.color = NewColor(1, 1, 1)
	middle.material.diffuse = 0.7
	middle.material.specular = 0.3
	middle.material.reflective = 1
	middle.material.refractiveIndex = 1
	middle.material.transparency = 1.0
	middle.material.shininess = 300

	right := NewSphere()
	right.SetTransform(Translation(1.5, 0.5, -0.5).MulMatrix(Scaling(0.5, 0.5, 0.5)))
	// right.material.pattern = CheckersPattern(NewColor(1, .4, .6), NewColor(0, 1, 1))
	// right.material.pattern.SetTransform(Scaling(0.2, 0.2, 0.2))
	right.material.color = NewColor(1, 1, 1)
	right.material.diffuse = 0.7
	right.material.specular = 0.3
	right.material.reflective = 0.5

	left := NewSphere()
	left.SetTransform(Translation(-1.5, 0.33, -0.75).MulMatrix(Scaling(0.33, 0.33, 0.33)))
	// left.material.pattern = PatternChain(middle.material.pattern, right.material.pattern)
	left.material.color = NewColor(1, 1, 1)
	left.material.diffuse = 0.7
	left.material.specular = 0.3
	left.material.reflective = .5

	lights := []*PointLight{
		NewPointLight(Point(-10, 10, -10), NewColor(1, 1, 1)),
		NewPointLight(Point(0, 10, 0), NewColor(0.5, 0.5, 0.5)),
	}
	world := NewWorld(lights, []Shape{middle, left, right, floor, leftWall, rightWall, backWall})

	camera := NewCamera(width, height, π/3)
	camera.SetTransform(ViewTransform(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0)))

	canvas := camera.Render(world, defaultRecursionDepth)

	fmt.Println(time.Now().Sub(start))

	canvas.ToPPM("world")
}

func planar(width, height int) {
	start := time.Now()
	lights := []*PointLight{
		NewPointLight(Point(-10, 10, -10), NewColor(1, 1, 1)),
		NewPointLight(Point(0, 10, 0), NewColor(0.5, 0.5, 0.5)),
	}
	p1 := NewPlane()
	world := NewWorld(lights, []Shape{p1})

	camera := NewCamera(width, height, π/3)
	camera.SetTransform(ViewTransform(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0)))

	canvas := camera.Render(world, defaultRecursionDepth)

	fmt.Println(time.Now().Sub(start))

	canvas.ToPPM("planar")

}
