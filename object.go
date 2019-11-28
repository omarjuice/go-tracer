package main

import "math"

//Object is any object in a scene
type Object interface {
	Intersect(*Ray) []*Intersection
}

//Sphere object...
type Sphere struct {
	center *Tuple
}

//NewSphere creates a new sphere
func NewSphere() *Sphere {
	return &Sphere{Point(0, 0, 0)}
}

//Intersect computes the intersection between a sphere and a ray
func (sphere *Sphere) Intersect(ray *Ray) []*Intersection {
	sphereToRay := ray.origin.Sub(sphere.center)
	a := ray.direction.Dot(ray.direction)
	b := 2 * ray.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b * b) - 4*a*c

	if discriminant < 0 {
		return []*Intersection{}
	}
	sqrtDisc := math.Sqrt(discriminant)
	t1 := (-b - sqrtDisc) / (2 * a)
	t2 := (-b + sqrtDisc) / (2 * a)

	return []*Intersection{&Intersection{t1, sphere}, &Intersection{t2, sphere}}

}
