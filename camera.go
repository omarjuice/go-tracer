package main

import (
	"math"
	"sync"
)

//Camera ...
type Camera struct {
	hsize, vsize                                  int
	fieldOfView, halfWidth, halfHeight, pixelSize float64
	tranform                                      Matrix
}

//NewCamera ...
func NewCamera(hsize, vsize int, fieldOfView float64) *Camera {
	c := &Camera{hsize, vsize, fieldOfView, 0, 0, 0, NewIdentityMatrix()}
	c.SetPixelSize()
	return c
}

//SetPixelSize updates the pixel size of the camera
func (cam *Camera) SetPixelSize() {
	halfView := math.Tan(cam.fieldOfView / 2)
	aspect := float64(cam.hsize) / float64(cam.vsize)
	if aspect >= 1 {
		cam.halfWidth = halfView
		cam.halfHeight = halfView / aspect
	} else {
		cam.halfWidth = halfView * aspect
		cam.halfHeight = halfView
	}
	cam.pixelSize = (cam.halfWidth * 2) / float64(cam.hsize)
}

//SetTransform ...
func (cam *Camera) SetTransform(transform Matrix) {
	cam.tranform = transform.Inverse()
}

//RayForPixel ...
func (cam *Camera) RayForPixel(x, y int) *Ray {
	px := float64(x)
	py := float64(y)
	xoffset := (px + 0.5) * cam.pixelSize
	yoffset := (py + 0.5) * cam.pixelSize

	worldx := cam.halfWidth - xoffset
	worldy := cam.halfHeight - yoffset

	pixel := cam.tranform.MulTuple(Point(worldx, worldy, -1))

	origin := cam.tranform.MulTuple(Point(0, 0, 0))

	direction := pixel.Sub(origin).Normalize()

	return NewRay(origin, direction)
}

//Render renders the image of a given world on a canavs from the view of the camera
func (cam *Camera) Render(world *World) *Canvas {
	image := NewCanvas(cam.hsize, cam.vsize)

	var wg sync.WaitGroup

	for y := 0; y < cam.vsize; y++ {
		wg.Add(1)
		go func(y int) {
			for x := 0; x < cam.hsize; x++ {
				ray := cam.RayForPixel(x, y)
				color := world.ColorAt(ray)
				image.WritePixel(x, y, color)
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
	return image
}
