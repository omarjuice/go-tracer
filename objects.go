package main

import (
	"math"
)

//Object is any object in a scene
type Object interface {
	Material() *Material
	Intersect(*Ray) []*Intersection
	NormalAt(*Tuple) *Tuple
}

//Sphere object...
type Sphere struct {
	origin    *Tuple
	transform Matrix
	material  *Material
}

//NewSphere creates a new sphere
func NewSphere() *Sphere {
	return &Sphere{Point(0, 0, 0), IdentityMatrix, DefaultMaterial()}
}

//SetTransform sets the spheres transformation
func (sphere *Sphere) SetTransform(transformation Matrix) {
	sphere.transform = transformation.Inverse()
}

//Intersect computes the intersection between a sphere and a ray
func (sphere *Sphere) Intersect(ray *Ray) []*Intersection {
	ray = ray.Transform(sphere.transform)
	sphereToRay := ray.origin.Sub(sphere.origin)
	a := ray.direction.Dot(ray.direction)
	b := 2 * ray.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b * b) - 4*a*c

	if discriminant < 0 {
		return []*Intersection{}
	}
	sqrtDisc := math.Sqrt(discriminant)
	div := (2 * a)
	t1 := (-b - sqrtDisc) / div
	t2 := (-b + sqrtDisc) / div
	return []*Intersection{&Intersection{t1, sphere, -1}, &Intersection{t2, sphere, -1}}

}

//NormalAt calculates the normal(vector perpendicular to the surface) at a given point
func (sphere *Sphere) NormalAt(point *Tuple) *Tuple {
	objectPoint := sphere.transform.MulTuple(point)
	objectNormal := objectPoint.Sub(sphere.origin)
	worldNormal := sphere.transform.Transpose().MulTuple(objectNormal)

	worldNormal.w = 0.0
	return worldNormal.Normalize()
}

//Material returns the material of a Sphere
func (sphere *Sphere) Material() *Material {
	return sphere.material
}
